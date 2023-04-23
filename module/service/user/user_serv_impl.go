package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	token "github.com/Billy278/MyGram/module/model/token"
	modelUser "github.com/Billy278/MyGram/module/model/user"
	repoUser "github.com/Billy278/MyGram/module/repository/user"
	cripto "github.com/Billy278/MyGram/pkg/cripto"
)

type USerServImpl struct {
	RepoUser repoUser.UserRepository
}

func NewUSerServImpl(repouser repoUser.UserRepository) USerServ {
	return &USerServImpl{
		RepoUser: repouser,
	}
}

func (u_serv *USerServImpl) SrvInsertUser(ctx context.Context, userIn modelUser.UserCreate) (userRes modelUser.UserRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvCreateUser", u_serv)
	log.Printf("%v invoked logCtx", logCtx)

	hashpass, err := cripto.GenerateHash(userIn.Password)
	if err != nil {
		log.Printf("[ERROR] error HashPassword :%v\n", err)
		return userRes, errors.New("HashPassword")
	}

	tnow := time.Now()
	userreq := modelUser.User{
		Username:   userIn.Username,
		Password:   hashpass,
		Email:      userIn.Email,
		Age:        userIn.Age,
		Created_at: &tnow,
	}
	user, err := u_serv.RepoUser.InsertUser(ctx, userreq)
	if err != nil {
		log.Printf("[ERROR] error create user :%v\n", err)
		return
	}
	userRes.Username = user.Username
	return
}
func (u_serv *USerServImpl) SrvLoginUsername(ctx context.Context, username string, password string) (tokens token.Tokens, err error) {
	logCtx := fmt.Sprintf("%T - SrvLoginUsername", u_serv)
	log.Printf("%v invoked logCtx", logCtx)

	userRes, err := u_serv.RepoUser.LoginUsername(ctx, username)
	if err != nil {
		log.Printf("[ERROR] error SrvLoginUsername :%v\n", err)
		return
	}
	err = cripto.CompareHash(userRes.Password, password)
	if err != nil {
		log.Printf("[ERROR] PASSWORD IS VAILED  :%v\n", err)
		return
	}
	//create token
	jti := "jti"
	usr_id := strconv.Itoa(int(userRes.Id))
	idToken, accessToken, refreshToken, err := u_serv.generateAllTokensConcurency(ctx, usr_id, userRes.Username, jti)
	if err != nil {
		return
	}

	return token.Tokens{
		IDToken:      idToken,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}

func (u_serv *USerServImpl) SrvFindByUsername(ctx context.Context, usernameIn string) (err error) {
	logCtx := fmt.Sprintf("%T - SrvFindByUsername", u_serv)
	log.Printf("%v invoked logCtx", logCtx)
	_, err = u_serv.RepoUser.FindByUsername(ctx, usernameIn)
	if err != nil {
		log.Printf("[INFO] SrvLoginUsername :%v\n", err)
		return
	}
	return
}
func (u_serv *USerServImpl) SrvFindByEmail(ctx context.Context, emailIn string) (err error) {
	logCtx := fmt.Sprintf("%T - SrvFindByEmail", u_serv)
	log.Printf("%v invoked logCtx", logCtx)
	_, err = u_serv.RepoUser.FindByUsername(ctx, emailIn)
	if err != nil {
		log.Printf("[INFO] SrvFindByEmail :%v\n", err)
		return
	}
	return
}

func (u_serv *USerServImpl) generateAllTokensConcurency(ctx context.Context, userid, username, jti string) (idToken, accessToken, refreshToken string, err error) {
	log.Printf("[INFO] Generate token  :%v\n", err)

	timeNow := time.Now()
	defaultClaim := token.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    userid,
		Audience:  "Challenges_12-13",
		JTI:       jti,
		Type:      token.ID_TOKEN,
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		//generate id token
		idTokenClaim := struct {
			token.DefaultClaim
			token.IDClaim
		}{
			defaultClaim_,
			token.IDClaim{
				Username: username,
			},
		}
		idToken, err = cripto.SignJWT(idTokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating id token  :%v\n", err)
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		//generate access token
		defaultClaim_.Expired = int(timeNow.Add(2 * time.Hour).Unix())
		defaultClaim_.Type = token.ACCESS_TOKEN
		accessTokenClaim := struct {
			token.DefaultClaim
			token.AccessClaim
		}{
			defaultClaim_,
			token.AccessClaim{
				UserID: userid,
			},
		}
		accessToken, err = cripto.SignJWT(accessTokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating access token  :%v\n", err)
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		//generate refresh token
		defaultClaim_.Expired = int(timeNow.Add(time.Hour).Unix())
		defaultClaim_.Type = token.REFRESH_TOKEN

		refreshTokenClaim := struct {
			token.DefaultClaim
		}{
			defaultClaim_,
		}
		refreshToken, err = cripto.SignJWT(refreshTokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating refresh token  :%v\n", err)
			return
		}
	}(defaultClaim)

	wg.Wait()
	return
}
