package server

import (
	router "github.com/Billy278/MyGram/module/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServer(v *validator.Validate) {
	ctrl := NewSetup(v)
	rt := gin.Default()
	//init middleware
	rt.Use(gin.Recovery(), gin.Logger())
	router.NewRouter(rt, ctrl.UserCtrl, ctrl.CommentCtrl, ctrl.PhotoCtrl, ctrl.MediaCtrl)
	rt.Run("0.0.0.0:8080")
}
