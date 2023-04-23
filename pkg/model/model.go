package model

import (
	modelComment "github.com/Billy278/MyGram/module/model/comment"
	modelPhoto "github.com/Billy278/MyGram/module/model/photo"
	modelMedia "github.com/Billy278/MyGram/module/model/socialmedia"
)

func ToPhotoResponse(photo modelPhoto.Photo) modelPhoto.PhotoRes {
	return modelPhoto.PhotoRes{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_url: photo.Photo_url,
		User_id:   photo.User_id,
	}
}

func ToPhotoResponses(photos []modelPhoto.Photo) (resPhoto []modelPhoto.PhotoRes) {
	for _, photo := range photos {
		resPhoto = append(resPhoto, ToPhotoResponse(photo))
	}
	return
}

func ToMediaResponse(media modelMedia.SocialMedia) modelMedia.SocialMediaRes {
	return modelMedia.SocialMediaRes{
		Id:               media.Id,
		Name:             media.Name,
		User_id:          media.User_id,
		Social_media_url: media.Social_media_url,
	}
}

func ToMediaResponses(mediaAll []modelMedia.SocialMedia) (resMedia []modelMedia.SocialMediaRes) {
	for _, media := range mediaAll {
		resMedia = append(resMedia, ToMediaResponse(media))
	}
	return
}

func ToCommentResponse(comment modelComment.Comment) modelComment.CommentRes {
	return modelComment.CommentRes{
		Id:       comment.Id,
		User_id:  comment.User_id,
		Photo_id: comment.Photo_id,
		Message:  comment.Message,
	}
}

func ToCommentResponses(commentAll []modelComment.Comment) (resComment []modelComment.CommentRes) {
	for _, comment := range commentAll {
		resComment = append(resComment, ToCommentResponse(comment))
	}
	return
}
