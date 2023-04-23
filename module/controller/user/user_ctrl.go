package user

import "github.com/gin-gonic/gin"

type UserCtrl interface {
	Registration(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}
