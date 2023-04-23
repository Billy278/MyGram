package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	modelUser "github.com/Billy278/MyGram/module/model/user"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{
		DB: db,
	}
}

func (u_repo *UserRepoImpl) InsertUser(ctx context.Context, userIn modelUser.User) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - CreateUser", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_user(username,email,password,age,created_at) VALUES($1,$2,$3,$4,$5) RETURNING username"
	row, err := u_repo.DB.QueryContext(ctx, sql, userIn.Username, userIn.Email, userIn.Password, userIn.Age, userIn.Created_at)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&user.Username)
		if err != nil {
			return
		}
		return
	} else {
		return user, errors.New("FAILED TO CREATE USER")
	}

}
func (u_repo *UserRepoImpl) UpdateUser(ctx context.Context, userIn modelUser.User) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - UpdateUser", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "UPDATE data_user SET username=$1,email=$2,password=$3,age=$4,updated_at=$5 WHERE id=$6"
	_, err = u_repo.DB.ExecContext(ctx, sql, userIn.Username, userIn.Email, userIn.Password, userIn.Age, userIn.Updated_at, userIn.Id)
	if err != nil {
		return
	}
	user.Id = userIn.Id
	return
}
func (u_repo *UserRepoImpl) LoginUsername(ctx context.Context, username string) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - FindUserByid", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,username,email,password,age FROM data_user WHERE username=$1"
	row, err := u_repo.DB.QueryContext(ctx, sql, username)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			return
		}
		return
	} else {
		return user, errors.New("USER NOT FOUND")
	}
}
func (u_repo *UserRepoImpl) FindAllUser(ctx context.Context) (users []modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - FindAllUser", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,username,email,password,age FROM data_user"
	rows, err := u_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer rows.Close()
	user := modelUser.User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return

}
func (u_repo *UserRepoImpl) DeleteUser(ctx context.Context, userid uint64) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - DeleteUserByid", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "DELETE FROM data_user WHERE id=$1"
	_, err = u_repo.DB.ExecContext(ctx, sql, userid)
	if err != nil {
		return
	}
	user.Id = userid
	return
}

func (u_repo *UserRepoImpl) FindByUsername(ctx context.Context, usernameIn string) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - FindUserByusername", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT username FROM data_user WHERE username=$1"
	row, err := u_repo.DB.QueryContext(ctx, sql, usernameIn)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		return user, errors.New("USERNAME SUDAH ADA")
	} else {
		return
	}
}
func (u_repo *UserRepoImpl) FindByEmail(ctx context.Context, emailIn string) (user modelUser.User, err error) {
	logCtx := fmt.Sprintf("%T - FindUserByEmail", u_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT email FROM data_user WHERE email=$1"
	row, err := u_repo.DB.QueryContext(ctx, sql, emailIn)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		return user, errors.New("EMAIL SUDAH ADA")
	} else {
		return
	}
}
