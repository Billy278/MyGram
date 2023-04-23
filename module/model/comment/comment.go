package comment

import "time"

type Comment struct {
	Id         uint64
	User_id    uint64
	Photo_id   uint64
	Message    string
	Created_at *time.Time
	Updated_at *time.Time
}
