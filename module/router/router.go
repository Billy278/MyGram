package router

import (
	ctrlComment "github.com/Billy278/MyGram/module/controller/comment"
	ctrlPhoto "github.com/Billy278/MyGram/module/controller/photo"
	ctrlMedia "github.com/Billy278/MyGram/module/controller/socialmedia"
	ctrlUser "github.com/Billy278/MyGram/module/controller/user"
	middleware "github.com/Billy278/MyGram/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, UserCtrl ctrlUser.UserCtrl, CommentCtrl ctrlComment.CommentCtrl, PhotoCtrl ctrlPhoto.PhotoCtrl, MediaCtrl ctrlMedia.MediaCtrl) {
	//user
	r.POST("/register", UserCtrl.Registration)
	r.POST("/login", UserCtrl.LoginUser)
	group := r.Group("/user", middleware.BearerOAuth())
	//photo
	group.GET("/photo", PhotoCtrl.CtrlFindAllPhoto)
	group.GET("/photo/:id", PhotoCtrl.CtrlFindByIdPhoto)
	group.POST("/photo", PhotoCtrl.CtrlInsertPhoto)
	group.PUT("/photo/:id", PhotoCtrl.CtrlUpdatePhoto)
	group.DELETE("/photo/:id", PhotoCtrl.CtrlDeletePhoto)

	//socialmedia
	group.GET("/media", MediaCtrl.CtlFindAllMedia)
	group.GET("/media/:id", MediaCtrl.CtlFindByIdMedia)
	group.POST("/media", MediaCtrl.CtlInsertMedia)
	group.PUT("/media/:id", MediaCtrl.CtlUpdateMedia)
	group.DELETE("/media/:id", MediaCtrl.CtlDeleteMedia)

	//comment
	group.GET("/comment", CommentCtrl.CtrlFindAllComment)
	group.GET("/comment/:id", CommentCtrl.CtrlFindByIdComment)
	group.POST("/comment", CommentCtrl.CtrlInsertComment)
	group.PUT("/comment/:id", CommentCtrl.CtrlUpdateComment)
	group.DELETE("/comment/:id", CommentCtrl.CtrlDeleteComment)
}
