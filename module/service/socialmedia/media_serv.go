package socialmedia

import (
	"context"

	modelMedia "github.com/Billy278/MyGram/module/model/socialmedia"
)

type SocialMediaSrv interface {
	SrvInsertMedia(ctx context.Context, MediaIn modelMedia.SocialMediaCreate) (MediaRes modelMedia.SocialMediaRes, err error)
	SrvUpdateMedia(ctx context.Context, MediaIn modelMedia.SocialMediaUpdate) (MediaRes modelMedia.SocialMediaRes, err error)
	SrvFindByIdMedia(ctx context.Context, MediaId uint64) (MediaRes modelMedia.SocialMediaRes, err error)
	SrvFindAllMedia(ctx context.Context) (MediaAllRes []modelMedia.SocialMediaRes, err error)
	SrvDeleteMedia(ctx context.Context, MediaId uint64, user_id uint64) (MediaRes modelMedia.SocialMediaRes, err error)
}
