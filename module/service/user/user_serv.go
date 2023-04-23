package user

import (
	"context"

	token "github.com/Billy278/MyGram/module/model/token"
	modelUser "github.com/Billy278/MyGram/module/model/user"
)

type USerServ interface {
	SrvInsertUser(ctx context.Context, userIn modelUser.UserCreate) (userRes modelUser.UserRes, err error)
	SrvLoginUsername(ctx context.Context, username string, password string) (tokens token.Tokens, err error)
	SrvFindByUsername(ctx context.Context, usernameIn string) (err error)
	SrvFindByEmail(ctx context.Context, emailIn string) (err error)
	//SrvUpdateUser(ctx context.Context, userIn modelUser.User) (user modelUser.User, err error)
	//SrvFindAllUser(ctx context.Context) (users []modelUser.User, err error)
	//SrvDeleteUser(ctx context.Context, userid uint64) (user modelUser.User, err error)
}
