package photo

import "time"

type Photo struct {
	Id         uint64
	Title      string
	Caption    string
	Photo_url  string
	User_id    int64
	Created_at *time.Time
	Updated_at *time.Time
}
