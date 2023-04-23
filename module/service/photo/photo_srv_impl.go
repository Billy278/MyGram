package photo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	modelPhoto "github.com/Billy278/MyGram/module/model/photo"
	repoPhoto "github.com/Billy278/MyGram/module/repository/photo"
	helper "github.com/Billy278/MyGram/pkg/model"
)

type PhotoSrvImpl struct {
	PhotoSrv repoPhoto.PhotoRepo
}

func NewPhotoSrvImpl(photosrv repoPhoto.PhotoRepo) PhotoSrv {
	return &PhotoSrvImpl{
		PhotoSrv: photosrv,
	}
}
func (p_srv *PhotoSrvImpl) SrvInsertPhoto(ctx context.Context, photoIn modelPhoto.PhotoCreate) (photoRes modelPhoto.PhotoRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvCreateUser", p_srv)
	log.Printf("%v invoked logCtx", logCtx)
	tnow := time.Now()
	photoreq := modelPhoto.Photo{
		Title:      photoIn.Title,
		Caption:    photoIn.Caption,
		Photo_url:  photoIn.Photo_url,
		User_id:    photoIn.User_id,
		Created_at: &tnow,
	}
	photores, err := p_srv.PhotoSrv.InsertPhoto(ctx, photoreq)
	if err != nil {
		log.Printf("[ERROR] error create Photo :%v\n", err)
		return
	}
	photoRes.Title = photores.Title
	return
}
func (p_srv *PhotoSrvImpl) SrvUpdatePhoto(ctx context.Context, photoIn modelPhoto.PhotoUpdate) (photoRes modelPhoto.PhotoRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvUpdatePhoto", p_srv)
	log.Printf("%v invoked logCtx", logCtx)
	photoreq, err := p_srv.PhotoSrv.FindByIdPhoto(ctx, photoIn.Id)
	if err != nil {
		log.Printf("[INFO] error SrvUpdatePhoto :%v\n", err)
		return
	}
	if photoreq.User_id != photoIn.User_id {
		return photoRes, errors.New("UNAUTHORIZED")
	}
	tnow := time.Now()
	photoreq.Title = photoIn.Title
	photoreq.Caption = photoIn.Caption
	photoreq.Photo_url = photoIn.Photo_url
	photoreq.Updated_at = &tnow
	fmt.Println(photoreq)
	photores, err := p_srv.PhotoSrv.UpdatePhoto(ctx, photoreq)
	if err != nil {
		log.Printf("[ERROR] error Update Photo :%v\n", err)
		return
	}
	photoRes.Title = photores.Title
	return
}
func (p_srv *PhotoSrvImpl) SrvFindByIdPhoto(ctx context.Context, photoId uint64) (photoRes modelPhoto.PhotoRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvFindByIdPhoto", p_srv)
	log.Printf("%v invoked logCtx", logCtx)
	photo, err := p_srv.PhotoSrv.FindByIdPhoto(ctx, photoId)
	if err != nil {
		log.Printf("[INFO] SrvFindByIdPhoto :%v\n", err)
		return
	}

	return helper.ToPhotoResponse(photo), err
}
func (p_srv *PhotoSrvImpl) SrvFindAllPhoto(ctx context.Context) (AllphotoRes []modelPhoto.PhotoRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvFindAllPhoto", p_srv)
	log.Printf("%v invoked logCtx", logCtx)
	allPhoto, err := p_srv.PhotoSrv.FindAllPhoto(ctx)
	if err != nil {
		log.Printf("[INFO] SrvFindAllPhoto :%v\n", err)
		return
	}

	return helper.ToPhotoResponses(allPhoto), err
}
func (p_srv *PhotoSrvImpl) SrvDeletePhoto(ctx context.Context, photoId uint64, user_id uint64) (photoRes modelPhoto.PhotoRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvDeletePhoto", p_srv)
	log.Printf("%v invoked logCtx", logCtx)
	photo, err := p_srv.PhotoSrv.FindByIdPhoto(ctx, photoId)
	if err != nil {
		log.Printf("[INFO] SrvDeletePhoto :%v\n", err)
		return
	}
	if photo.User_id != int64(user_id) {
		return photoRes, errors.New("UNAUTHORIZED")
	}
	photores, err := p_srv.PhotoSrv.DeletePhoto(ctx, photoId)
	if err != nil {
		log.Printf("[INFO] SrvDeletePhoto :%v\n", err)
		return
	}
	photoRes.Id = photores.Id
	return
}
