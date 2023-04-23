package photo

import "github.com/gin-gonic/gin"

type PhotoCtrl interface {
	CtrlInsertPhoto(ctx *gin.Context)
	CtrlUpdatePhoto(ctx *gin.Context)
	CtrlFindByIdPhoto(ctx *gin.Context)
	CtrlFindAllPhoto(ctx *gin.Context)
	CtrlDeletePhoto(ctx *gin.Context)
}
