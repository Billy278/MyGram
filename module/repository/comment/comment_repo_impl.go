package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	modelComment "github.com/Billy278/MyGram/module/model/comment"
)

type CommentRepoImpl struct {
	DB *sql.DB
}

func NewCommentRepoImpl(db *sql.DB) CommentRepo {
	return &CommentRepoImpl{
		DB: db,
	}
}

func (c_repo *CommentRepoImpl) InsertComment(ctx context.Context, CommentIn modelComment.Comment) (Comment modelComment.Comment, err error) {
	logCtx := fmt.Sprintf("%T - CreateComment", c_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "INSERT INTO data_comment(user_id,photo_id,message,created_at) VALUES($1,$2,$3,$4) RETURNING message"
	row, err := c_repo.DB.QueryContext(ctx, sql, CommentIn.User_id, CommentIn.Photo_id, CommentIn.Message, CommentIn.Created_at)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&Comment.Message)
		if err != nil {
			return
		}
		return
	} else {
		return Comment, errors.New("FAILED TO CREATE COMMENT")
	}
}
func (c_repo *CommentRepoImpl) UpdateComment(ctx context.Context, CommentIn modelComment.Comment) (Comment modelComment.Comment, err error) {
	logCtx := fmt.Sprintf("%T - UpdateComment", c_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "UPDATE data_comment SET photo_id=$1,message=$2,updated_at=$3 WHERE id=$4 RETURNING message"
	row, err := c_repo.DB.QueryContext(ctx, sql, CommentIn.Photo_id, CommentIn.Message, CommentIn.Updated_at, CommentIn.Id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&Comment.Message)
		if err != nil {
			return
		}
		return
	} else {
		return Comment, errors.New("FAILED TO UPDATE COMMENT")
	}
}

func (c_repo *CommentRepoImpl) FindByIdComment(ctx context.Context, CommentId uint64) (Comment modelComment.Comment, err error) {
	logCtx := fmt.Sprintf("%T - FindByIdComment", c_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,user_id,photo_id,message FROM data_comment WHERE id=$1"
	row, err := c_repo.DB.QueryContext(ctx, sql, CommentId)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&Comment.Id, &Comment.User_id, &Comment.Photo_id, &Comment.Message)
		if err != nil {
			return
		}
		return
	} else {
		return Comment, errors.New("COMMENT IS NOT FOUND")
	}
}
func (c_repo *CommentRepoImpl) FindAllComment(ctx context.Context) (CommentAll []modelComment.Comment, err error) {
	logCtx := fmt.Sprintf("%T - FindByAllComment", c_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "SELECT id,user_id,photo_id,message FROM data_comment"
	row, err := c_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer row.Close()
	Comment := modelComment.Comment{}
	for row.Next() {
		err = row.Scan(&Comment.Id, &Comment.User_id, &Comment.Photo_id, &Comment.Message)
		if err != nil {
			return
		}
		CommentAll = append(CommentAll, Comment)
	}
	return
}
func (c_repo *CommentRepoImpl) DeleteComment(ctx context.Context, CommentId uint64) (Comment modelComment.Comment, err error) {
	logCtx := fmt.Sprintf("%T - DeleteComment", c_repo)
	log.Printf("%v invoked logCtx", logCtx)
	sql := "DELETE FROM data_comment WHERE id=$1"
	_, err = c_repo.DB.QueryContext(ctx, sql, CommentId)
	if err != nil {
		return
	}
	Comment.Id = CommentId
	return
}
