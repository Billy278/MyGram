package photo

type PhotoRes struct {
	Id        uint64 `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   int64  `json:"user_id"`
}
