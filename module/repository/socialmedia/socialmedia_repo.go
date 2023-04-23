package socialmedia

import (
	"context"

	modelMedia "github.com/Billy278/MyGram/module/model/socialmedia"
)

type SocialMediaRepo interface {
	InsertMedia(ctx context.Context, MediaIn modelMedia.SocialMedia) (Media modelMedia.SocialMedia, err error)
	UpdateMedia(ctx context.Context, MediaIn modelMedia.SocialMedia) (Media modelMedia.SocialMedia, err error)
	FindByIdMedia(ctx context.Context, MediaId uint64) (Media modelMedia.SocialMedia, err error)
	FindAllMedia(ctx context.Context) (MediaAll []modelMedia.SocialMedia, err error)
	DeleteMedia(ctx context.Context, MediaId uint64) (Media modelMedia.SocialMedia, err error)
}
