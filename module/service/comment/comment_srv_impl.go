package comment

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	modelComment "github.com/Billy278/MyGram/module/model/comment"
	repoComent "github.com/Billy278/MyGram/module/repository/comment"
	servPhoto "github.com/Billy278/MyGram/module/service/photo"
	helper "github.com/Billy278/MyGram/pkg/model"
)

type CommentSrvImpl struct {
	RepoComment repoComent.CommentRepo
	ServPhoto   servPhoto.PhotoSrv
}

func NewCommentSrvImpl(repo repoComent.CommentRepo, serv servPhoto.PhotoSrv) CommentSrv {
	return &CommentSrvImpl{
		RepoComment: repo,
		ServPhoto:   serv,
	}
}

func (c_srv *CommentSrvImpl) SrvInsertComment(ctx context.Context, CommentIn modelComment.CommentCreate) (CommentRes modelComment.CommentRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvInsertComment", c_srv)
	log.Printf("%v invoked logCtx", logCtx)
	_, err = c_srv.ServPhoto.SrvFindByIdPhoto(ctx, CommentIn.Photo_id)
	if err != nil {
		log.Printf("[INFO] SrvInsertComment :%v\n", err)
		return
	}
	tnow := time.Now()
	commentreq := modelComment.Comment{
		User_id:    CommentIn.User_id,
		Photo_id:   CommentIn.Photo_id,
		Message:    CommentIn.Message,
		Created_at: &tnow,
	}

	commentres, err := c_srv.RepoComment.InsertComment(ctx, commentreq)
	CommentRes.Message = commentres.Message
	return
}

func (c_srv *CommentSrvImpl) SrvUpdateComment(ctx context.Context, CommentIn modelComment.CommentUpdate) (CommentRes modelComment.CommentRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvUpdateComment", c_srv)
	log.Printf("%v invoked logCtx", logCtx)
	_, err = c_srv.ServPhoto.SrvFindByIdPhoto(ctx, CommentIn.Photo_id)
	if err != nil {
		log.Printf("[INFO] SrvUpdateComment :%v\n", err)
		return
	}
	commentreq, err := c_srv.RepoComment.FindByIdComment(ctx, CommentIn.Id)
	if err != nil {
		log.Printf("[INFO] SrvUpdateComment :%v\n", err)
		return
	}
	if commentreq.User_id != CommentIn.User_id {
		return CommentRes, errors.New("UNAUTHORIZED")
	}

	tnow := time.Now()
	commentreq.Photo_id = CommentIn.Photo_id
	commentreq.Message = CommentIn.Message
	commentreq.Updated_at = &tnow
	commentres, err := c_srv.RepoComment.UpdateComment(ctx, commentreq)
	if err != nil {
		log.Printf("[INFO] SrvUpdateComment :%v\n", err)
		return
	}
	CommentRes.Message = commentres.Message
	return
}
func (c_srv *CommentSrvImpl) SrvFindByIdComment(ctx context.Context, CommentId uint64) (CommentRes modelComment.CommentRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvFindByIdComment", c_srv)
	log.Printf("%v invoked logCtx", logCtx)
	commentreq, err := c_srv.RepoComment.FindByIdComment(ctx, CommentId)
	if err != nil {
		log.Printf("[INFO] SrvFindByIdComment :%v\n", err)
		return
	}

	return helper.ToCommentResponse(commentreq), err
}
func (c_srv *CommentSrvImpl) SrvFindAllComment(ctx context.Context) (CommentAll []modelComment.CommentRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvFindAllComment", c_srv)
	log.Printf("%v invoked logCtx", logCtx)
	commentreq, err := c_srv.RepoComment.FindAllComment(ctx)
	if err != nil {
		log.Printf("[INFO] SrvFindAllComment :%v\n", err)
		return
	}

	return helper.ToCommentResponses(commentreq), err

}
func (c_srv *CommentSrvImpl) SrvDeleteComment(ctx context.Context, CommentId uint64, user_id uint64) (CommentRes modelComment.CommentRes, err error) {
	logCtx := fmt.Sprintf("%T - SrvDeleteComment", c_srv)
	log.Printf("%v invoked logCtx", logCtx)
	commentreq, err := c_srv.RepoComment.FindByIdComment(ctx, CommentId)
	if err != nil {
		log.Printf("[INFO] SrvDeleteComment :%v\n", err)
		return
	}
	if commentreq.User_id != user_id {
		return CommentRes, errors.New("UNAUTHORIZED")
	}
	res, err := c_srv.RepoComment.DeleteComment(ctx, CommentId)
	if err != nil {
		log.Printf("[INFO] SrvDeleteComment :%v\n", err)
		return
	}
	CommentRes.Id = res.Id
	return
}
