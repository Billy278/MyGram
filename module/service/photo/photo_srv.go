package photo

import (
	"context"

	modelPhoto "github.com/Billy278/MyGram/module/model/photo"
)

type PhotoSrv interface {
	SrvInsertPhoto(ctx context.Context, photoIn modelPhoto.PhotoCreate) (photoRes modelPhoto.PhotoRes, err error)
	SrvUpdatePhoto(ctx context.Context, photoIn modelPhoto.PhotoUpdate) (photoRes modelPhoto.PhotoRes, err error)
	SrvFindByIdPhoto(ctx context.Context, photoId uint64) (photoRes modelPhoto.PhotoRes, err error)
	SrvFindAllPhoto(ctx context.Context) (AllphotoRes []modelPhoto.PhotoRes, err error)
	SrvDeletePhoto(ctx context.Context, photoId uint64, user_id uint64) (photoRes modelPhoto.PhotoRes, err error)
}
