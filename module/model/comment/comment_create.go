package comment

import "time"

type CommentCreate struct {
	User_id    uint64     `json:"user_id" `
	Photo_id   uint64     `json:"photo_id" validate:"required"`
	Message    string     `json:"message" validate:"required"`
	Created_at *time.Time `json:"created_at"`
}
