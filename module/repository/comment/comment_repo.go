package comment

import (
	"context"

	modelComment "github.com/Billy278/MyGram/module/model/comment"
)

type CommentRepo interface {
	InsertComment(ctx context.Context, CommentIn modelComment.Comment) (Comment modelComment.Comment, err error)
	UpdateComment(ctx context.Context, CommentIn modelComment.Comment) (Comment modelComment.Comment, err error)
	FindByIdComment(ctx context.Context, CommentId uint64) (Comment modelComment.Comment, err error)
	FindAllComment(ctx context.Context) (CommentAll []modelComment.Comment, err error)
	DeleteComment(ctx context.Context, CommentId uint64) (Comment modelComment.Comment, err error)
}
