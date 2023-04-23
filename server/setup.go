package server

import (
	"github.com/Billy278/MyGram/db"
	ctrlComment "github.com/Billy278/MyGram/module/controller/comment"
	ctrlPhoto "github.com/Billy278/MyGram/module/controller/photo"
	ctrlMedia "github.com/Billy278/MyGram/module/controller/socialmedia"
	ctrlUser "github.com/Billy278/MyGram/module/controller/user"
	repositoryComment "github.com/Billy278/MyGram/module/repository/comment"
	repositoryPhoto "github.com/Billy278/MyGram/module/repository/photo"
	repositoryMedia "github.com/Billy278/MyGram/module/repository/socialmedia"
	repositoryUser "github.com/Billy278/MyGram/module/repository/user"
	servicevComment "github.com/Billy278/MyGram/module/service/comment"
	servicePhoto "github.com/Billy278/MyGram/module/service/photo"
	serviceMedia "github.com/Billy278/MyGram/module/service/socialmedia"
	serviceUser "github.com/Billy278/MyGram/module/service/user"
	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	UserCtrl    ctrlUser.UserCtrl
	MediaCtrl   ctrlMedia.MediaCtrl
	PhotoCtrl   ctrlPhoto.PhotoCtrl
	CommentCtrl ctrlComment.CommentCtrl
}

func NewSetup(validate *validator.Validate) Controllers {
	datastore := db.NewDBPostges()
	//user
	repoUser := repositoryUser.NewUserRepoImpl(datastore)
	srvUser := serviceUser.NewUSerServImpl(repoUser)
	ctrUser := ctrlUser.NewUserCtrlImpl(srvUser, validate)

	//photo
	repoPhoto := repositoryPhoto.NewPhotoRepoImpl(datastore)
	srvPhoto := servicePhoto.NewPhotoSrvImpl(repoPhoto)
	ctrPhoto := ctrlPhoto.NewPhotoCtrlImpl(srvPhoto, validate)

	//social media
	repoMedia := repositoryMedia.NewSocialMediaRepoImpl(datastore)
	srvMedia := serviceMedia.NewSocialMediaSrvImp(repoMedia)
	ctrMedia := ctrlMedia.NewMediaCtrlImpl(srvMedia, validate)

	//comment
	repoComment := repositoryComment.NewCommentRepoImpl(datastore)
	srvComment := servicevComment.NewCommentSrvImpl(repoComment, srvPhoto)
	ctrComment := ctrlComment.NewCommentCtrlImpl(srvComment, validate)

	return Controllers{
		UserCtrl:    ctrUser,
		PhotoCtrl:   ctrPhoto,
		MediaCtrl:   ctrMedia,
		CommentCtrl: ctrComment,
	}
}
