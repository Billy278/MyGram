package server

import (
	docs "github.com/Billy278/MyGram/docs"
	router "github.com/Billy278/MyGram/module/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(v *validator.Validate) {
	ctrl := NewSetup(v)
	rt := gin.Default()
	//init middleware
	rt.Use(gin.Recovery(), gin.Logger())
	docs.SwaggerInfo.BasePath = "/api/v1"
	// swagger
	rt.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NewRouter(rt, ctrl.UserCtrl, ctrl.CommentCtrl, ctrl.PhotoCtrl, ctrl.MediaCtrl)
	rt.Run("0.0.0.0:8080")
}
