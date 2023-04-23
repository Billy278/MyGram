package photo

import (
	"context"

	modelPhoto "github.com/Billy278/MyGram/module/model/photo"
)

type PhotoRepo interface {
	InsertPhoto(ctx context.Context, photoIn modelPhoto.Photo) (photo modelPhoto.Photo, err error)
	UpdatePhoto(ctx context.Context, photoIn modelPhoto.Photo) (photo modelPhoto.Photo, err error)
	FindByIdPhoto(ctx context.Context, photoId uint64) (photo modelPhoto.Photo, err error)
	FindAllPhoto(ctx context.Context) (photos []modelPhoto.Photo, err error)
	DeletePhoto(ctx context.Context, photoId uint64) (photo modelPhoto.Photo, err error)
}
