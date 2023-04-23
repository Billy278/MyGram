package socialmedia

type SocialMediaRes struct {
	Id               uint64 `json:"id"`
	Name             string `json:"name"`
	Social_media_url string `json:"media_url"`
	User_id          uint64 `json:"user_id"`
}
