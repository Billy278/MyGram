package photo

type PhotoCreate struct {
	Title     string `json:"title" validate:"required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" validate:"required"`
	User_id   int64  `json:"user_id"`
}
