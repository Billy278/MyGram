package comment

import "github.com/gin-gonic/gin"

type CommentCtrl interface {
	CtrlInsertComment(ctx *gin.Context)
	CtrlUpdateComment(ctx *gin.Context)
	CtrlFindByIdComment(ctx *gin.Context)
	CtrlFindAllComment(ctx *gin.Context)
	CtrlDeleteComment(ctx *gin.Context)
}
