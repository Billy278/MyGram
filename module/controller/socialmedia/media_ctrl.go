package socialmedia

import "github.com/gin-gonic/gin"

type MediaCtrl interface {
	CtlInsertMedia(ctx *gin.Context)
	CtlUpdateMedia(ctx *gin.Context)
	CtlFindByIdMedia(ctx *gin.Context)
	CtlFindAllMedia(ctx *gin.Context)
	CtlDeleteMedia(ctx *gin.Context)
}
