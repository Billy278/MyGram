package comment

type CommentRes struct {
	Id       uint64 `json:"id"`
	User_id  uint64 `json:"user_id"`
	Photo_id uint64 `json:"photo_id"`
	Message  string `json:"message"`
}
