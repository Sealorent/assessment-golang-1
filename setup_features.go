package main

import (
	"final_project/controller/comment_control"
	"final_project/controller/photo_control"
	"final_project/controller/social_media_control"
	"final_project/controller/user_control"
	"final_project/repository/comment_repo"
	"final_project/repository/photo_repo"
	"final_project/repository/social_media_repo"
	"final_project/repository/user_repo"

	"gorm.io/gorm"
)

type SetupFeatures struct {
	UserController        *user_control.UserController
	CommentController     *comment_control.CommentController
	SocialMediaController *social_media_control.SocialMediaController
	PhotoController       *photo_control.PhotoController
}

func NewSetupFeatures(db *gorm.DB) *SetupFeatures {
	
	userRepository := user_repo.NewUserRepository(db)
	userController := user_control.NewUserController(userRepository)

	photoRepository := photo_repo.NewPhotoRepository(db)
	photoController := photo_control.NewPhotoController(photoRepository)

	commentRepository := comment_repo.NewCommentRepository(db)
	commentController := comment_control.NewCommentController(commentRepository)

	socialMediaRepository := social_media_repo.NewSocialMediaRepository(db)
	socialMediaController := social_media_control.NewSocialMediaController(socialMediaRepository)

	return &SetupFeatures{
		UserController:        userController,
		CommentController:     commentController,
		SocialMediaController: socialMediaController,
		PhotoController:       photoController,
	}
}
