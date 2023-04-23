package comment

import (
	"context"

	modelComment "github.com/Billy278/MyGram/module/model/comment"
)

type CommentSrv interface {
	SrvInsertComment(ctx context.Context, CommentIn modelComment.CommentCreate) (CommentRes modelComment.CommentRes, err error)
	SrvUpdateComment(ctx context.Context, CommentIn modelComment.CommentUpdate) (CommentRes modelComment.CommentRes, err error)
	SrvFindByIdComment(ctx context.Context, CommentId uint64) (CommentRes modelComment.CommentRes, err error)
	SrvFindAllComment(ctx context.Context) (CommentAll []modelComment.CommentRes, err error)
	SrvDeleteComment(ctx context.Context, CommentId uint64, user_id uint64) (CommentRes modelComment.CommentRes, err error)
}
