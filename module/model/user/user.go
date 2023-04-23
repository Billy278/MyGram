package user

import "time"

type User struct {
	Id         uint64
	Username   string
	Email      string
	Password   string
	Age        uint64
	Created_at *time.Time
	Updated_at *time.Time
}
