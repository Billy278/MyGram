package socialmedia

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	modelMedia "github.com/Billy278/MyGram/module/model/socialmedia"
)

type SocialMediaRepoImpl struct {
	DB *sql.DB
}

func NewSocialMediaRepoImpl(db *sql.DB) SocialMediaRepo {
	return &SocialMediaRepoImpl{
		DB: db,
	}
}

func (s_repo *SocialMediaRepoImpl) InsertMedia(ctx context.Context, MediaIn modelMedia.SocialMedia) (Media modelMedia.SocialMedia, err error) {
	logCtx := fmt.Sprintf("%T - CreateMedia", s_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_media(name,social_media_url,user_id,created_at) VALUES($1,$2,$3,$4) RETURNING name"
	row, err := s_repo.DB.QueryContext(ctx, sql, MediaIn.Name, MediaIn.Social_media_url, MediaIn.User_id, MediaIn.Created_at)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&Media.Name)
		if err != nil {
			return
		}
		return
	} else {
		return Media, errors.New("FAILED TO CREATE SOCIAL MEDIA")
	}
}
func (s_repo *SocialMediaRepoImpl) UpdateMedia(ctx context.Context, MediaIn modelMedia.SocialMedia) (Media modelMedia.SocialMedia, err error) {
	logCtx := fmt.Sprintf("%T - UpdateMedia", s_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "UPDATE data_media SET name=$1,social_media_url=$2,updated_at=$3 WHERE id=$4 RETURNING name"
	row, err := s_repo.DB.QueryContext(ctx, sql, MediaIn.Name, MediaIn.Social_media_url, MediaIn.Updated_at, MediaIn.Id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&Media.Name)
		if err != nil {
			return
		}
		return
	} else {
		return Media, errors.New("FAILED TO UPDATE SOCIAL MEDIA")
	}
}
func (s_repo *SocialMediaRepoImpl) FindByIdMedia(ctx context.Context, MediaId uint64) (Media modelMedia.SocialMedia, err error) {
	logCtx := fmt.Sprintf("%T - FIndByIdMedia", s_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,name,social_media_url,user_id FROM data_media WHERE id=$1"
	row, err := s_repo.DB.QueryContext(ctx, sql, MediaId)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&Media.Id, &Media.Name, &Media.Social_media_url, &Media.User_id)
		if err != nil {
			return
		}
		return
	} else {
		return Media, errors.New(" ID SOCIAL MEDIA IS NOT FOUND")
	}
}
func (s_repo *SocialMediaRepoImpl) FindAllMedia(ctx context.Context) (MediaAll []modelMedia.SocialMedia, err error) {
	logCtx := fmt.Sprintf("%T - FIndAllMedia", s_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,name,social_media_url,user_id FROM data_media"
	row, err := s_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer row.Close()
	Media := modelMedia.SocialMedia{}
	for row.Next() {
		err = row.Scan(&Media.Id, &Media.Name, &Media.Social_media_url, &Media.User_id)
		if err != nil {
			return
		}
		MediaAll = append(MediaAll, Media)
	}
	return
}
func (s_repo *SocialMediaRepoImpl) DeleteMedia(ctx context.Context, MediaId uint64) (Media modelMedia.SocialMedia, err error) {
	logCtx := fmt.Sprintf("%T - DeleteMedia", s_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "DELETE FROM data_media WHERE id=$1"
	_, err = s_repo.DB.ExecContext(ctx, sql, MediaId)
	if err != nil {
		return
	}
	Media.Id = MediaId
	return
}
