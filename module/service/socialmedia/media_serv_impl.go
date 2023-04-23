package socialmedia

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	modelMedia "github.com/Billy278/MyGram/module/model/socialmedia"
	repoMedia "github.com/Billy278/MyGram/module/repository/socialmedia"
	helper "github.com/Billy278/MyGram/pkg/model"
)

type SocialMediaSrvImp struct {
	RepoMedia repoMedia.SocialMediaRepo
}

func NewSocialMediaSrvImp(repomedia repoMedia.SocialMediaRepo) SocialMediaSrv {
	return &SocialMediaSrvImp{
		RepoMedia: repomedia,
	}
}

func (s_serv *SocialMediaSrvImp) SrvInsertMedia(ctx context.Context, MediaIn modelMedia.SocialMediaCreate) (MediaRes modelMedia.SocialMediaRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvInsertMedia", s_serv)
	log.Printf("%v invoked logCtx", logCtx)
	tnow := time.Now()
	mediareq := modelMedia.SocialMedia{
		Name:             MediaIn.Name,
		User_id:          MediaIn.User_id,
		Social_media_url: MediaIn.Social_media_url,
		Created_at:       &tnow,
	}
	mediares, err := s_serv.RepoMedia.InsertMedia(ctx, mediareq)
	if err != nil {
		log.Printf("[ERROR] error create Media :%v\n", err)
		return
	}
	MediaRes.Name = mediares.Name
	return
}
func (s_serv *SocialMediaSrvImp) SrvUpdateMedia(ctx context.Context, MediaIn modelMedia.SocialMediaUpdate) (MediaRes modelMedia.SocialMediaRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvUpdateMedia", s_serv)
	log.Printf("%v invoked logCtx", logCtx)
	mediareq, err := s_serv.RepoMedia.FindByIdMedia(ctx, MediaIn.Id)
	if err != nil {
		log.Printf("[INFO] SrvUpdateMedia :%v\n", err)
		return
	}

	if mediareq.User_id != MediaIn.User_id {
		return MediaRes, errors.New("UNAUTHORIZED")
	}
	tnow := time.Now()
	mediareq.Name = MediaIn.Name
	mediareq.Social_media_url = MediaIn.Social_media_url
	mediareq.Updated_at = &tnow
	res, err := s_serv.RepoMedia.UpdateMedia(ctx, mediareq)
	if err != nil {
		log.Printf("[ERROR] error Update media :%v\n", err)
		return
	}
	MediaRes.Name = res.Name
	return
}
func (s_serv *SocialMediaSrvImp) SrvFindByIdMedia(ctx context.Context, MediaId uint64) (MediaRes modelMedia.SocialMediaRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvFindByIdMedia", s_serv)
	log.Printf("%v invoked logCtx", logCtx)
	media, err := s_serv.RepoMedia.FindByIdMedia(ctx, MediaId)
	if err != nil {
		log.Printf("[INFO] SrvFindByIdMedia :%v\n", err)
		return
	}

	return helper.ToMediaResponse(media), err
}
func (s_serv *SocialMediaSrvImp) SrvFindAllMedia(ctx context.Context) (MediaAllRes []modelMedia.SocialMediaRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvFindAllMedia", s_serv)
	log.Printf("%v invoked logCtx", logCtx)
	allMedia, err := s_serv.RepoMedia.FindAllMedia(ctx)
	if err != nil {
		log.Printf("[INFO] SrvFindAllMedia :%v\n", err)
		return
	}

	return helper.ToMediaResponses(allMedia), err
}
func (s_serv *SocialMediaSrvImp) SrvDeleteMedia(ctx context.Context, MediaId uint64, user_id uint64) (MediaRes modelMedia.SocialMediaRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvDeleteMedia", s_serv)
	log.Printf("%v invoked logCtx", logCtx)
	media, err := s_serv.RepoMedia.FindByIdMedia(ctx, MediaId)
	if err != nil {
		log.Printf("[INFO] SrvDeleteMedia :%v\n", err)
		return
	}
	if media.User_id != user_id {
		return MediaRes, errors.New("UNAUTHORIZED")
	}
	mediares, err := s_serv.RepoMedia.DeleteMedia(ctx, MediaId)
	if err != nil {
		log.Printf("[INFO]  SrvDeleteMedia :%v\n", err)
		return
	}
	MediaRes.Id = mediares.Id
	return
}
