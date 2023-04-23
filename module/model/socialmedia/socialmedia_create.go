package socialmedia

type SocialMediaCreate struct {
	Name             string `json:"name" validate:"required"`
	Social_media_url string `json:"media_url" validate:"required"`
	User_id          uint64 `json:"user_id"`
}
