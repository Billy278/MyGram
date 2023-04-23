package socialmedia

import "time"

type SocialMedia struct {
	Id               uint64
	Name             string
	Social_media_url string
	User_id          uint64
	Created_at       *time.Time
	Updated_at       *time.Time
}
