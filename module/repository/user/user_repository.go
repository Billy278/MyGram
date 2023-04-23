package repository

import (
	"context"

	modelUser "github.com/Billy278/MyGram/module/model/user"
)

type UserRepository interface {
	InsertUser(ctx context.Context, userIn modelUser.User) (user modelUser.User, err error)
	UpdateUser(ctx context.Context, userIn modelUser.User) (user modelUser.User, err error)
	LoginUsername(ctx context.Context, username string) (user modelUser.User, err error)
	FindAllUser(ctx context.Context) (users []modelUser.User, err error)
	DeleteUser(ctx context.Context, userid uint64) (user modelUser.User, err error)
	FindByUsername(ctx context.Context, usernameIn string) (user modelUser.User, err error)
	FindByEmail(ctx context.Context, emailIn string) (user modelUser.User, err error)
}
