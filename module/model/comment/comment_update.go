package comment

import "time"

type CommentUpdate struct {
	Id         uint64     `json:"id"`
	User_id    uint64     `json:"user_id"`
	Photo_id   uint64     `json:"photo_id" validate:"required"`
	Message    string     `json:"message" validate:"required"`
	Updated_at *time.Time `json:"updated_at"`
}
