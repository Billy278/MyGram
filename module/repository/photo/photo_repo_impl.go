package photo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	modelPhoto "github.com/Billy278/MyGram/module/model/photo"
)

type PhotoRepoImpl struct {
	DB *sql.DB
}

func NewPhotoRepoImpl(db *sql.DB) PhotoRepo {
	return &PhotoRepoImpl{
		DB: db,
	}
}

func (p_repo *PhotoRepoImpl) InsertPhoto(ctx context.Context, photoIn modelPhoto.Photo) (photo modelPhoto.Photo, err error) {
	logCtx := fmt.Sprintf("%T - CreatePhoto", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_photo(title,caption,photo_url,user_id,created_at) VALUES($1,$2,$3,$4,$5) RETURNING title"
	row, err := p_repo.DB.QueryContext(ctx, sql, photoIn.Title, photoIn.Caption, photoIn.Photo_url, photoIn.User_id, photoIn.Created_at)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&photo.Title)
		if err != nil {
			return
		}
		return
	} else {
		return photo, errors.New("FAILED TO CREATE PHOTO")
	}
}
func (p_repo *PhotoRepoImpl) UpdatePhoto(ctx context.Context, photoIn modelPhoto.Photo) (photo modelPhoto.Photo, err error) {
	logCtx := fmt.Sprintf("%T - UpdatePhoto", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "UPDATE data_photo SET title=$1,caption=$2,photo_url=$3,updated_at=$4 WHERE id=$5 RETURNING title"
	row, err := p_repo.DB.QueryContext(ctx, sql, photoIn.Title, photoIn.Caption, photoIn.Photo_url, photoIn.Updated_at, photoIn.Id)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&photo.Title)
		if err != nil {
			return
		}
		return
	} else {
		return photo, errors.New("FAILED TO UPDATE PHOTO")
	}
}
func (p_repo *PhotoRepoImpl) FindByIdPhoto(ctx context.Context, photoId uint64) (photo modelPhoto.Photo, err error) {
	logCtx := fmt.Sprintf("%T - FindByIdPhoto", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,title,caption,photo_url,user_id FROM data_photo WHERE id=$1"
	row, err := p_repo.DB.QueryContext(ctx, sql, photoId)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Photo_url, &photo.User_id)
		if err != nil {
			return
		}
		return
	} else {
		return photo, errors.New("PHOTO IS NOT FOUND")
	}
}
func (p_repo *PhotoRepoImpl) FindAllPhoto(ctx context.Context) (photos []modelPhoto.Photo, err error) {
	logCtx := fmt.Sprintf("%T - FindAllPhoto", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,title,caption,photo_url,user_id FROM data_photo"
	row, err := p_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer row.Close()
	photo := modelPhoto.Photo{}
	for row.Next() {
		err = row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Photo_url, &photo.User_id)
		if err != nil {
			return
		}
		photos = append(photos, photo)
	}
	return
}
func (p_repo *PhotoRepoImpl) DeletePhoto(ctx context.Context, photoId uint64) (photo modelPhoto.Photo, err error) {
	logCtx := fmt.Sprintf("%T - DeleteByIdPhoto", p_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "DELETE FROM data_photo WHERE id=$1"
	_, err = p_repo.DB.ExecContext(ctx, sql, photoId)
	if err != nil {
		return
	}
	photo.Id = photoId
	return
}
