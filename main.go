package main

import (
	"final_project/controller/comment_control"
	"final_project/controller/photo_control"
	"final_project/controller/social_media_control"
	"final_project/controller/user_control"
	"final_project/lib"
	"final_project/middleware"
	"final_project/model"
	"final_project/repository/comment_repo"
	"final_project/repository/photo_repo"
	"final_project/repository/social_media_repo"
	"final_project/repository/user_repo"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// gorm connection
	db, err := lib.InitDB()
	if err != nil {
		panic(err)
	}

	// migrate the schema
	err = db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})
	if err != nil {
		panic(err)
	}
	// USER
	userRepository := user_repo.NewUserRepository(db)
	userController := user_control.NewUserController(userRepository)
	// PHOTO
	photoRepository := photo_repo.NewPhotoRepository(db)
	photoController := photo_control.NewPhotoController(photoRepository)
	// COMMENT
	commentRepository := comment_repo.NewCommentRepository(db)
	commentController := comment_control.NewCommentController(commentRepository)
	// SOCIAL MEDIA
	socialMediaRepository := social_media_repo.NewSocialMediaRepository(db)
	socialMediaController := social_media_control.NewSocialMediaController(socialMediaRepository)

	// ROUTES
	ginEngine := gin.Default()
	ginEngine.POST("/auth/register", userController.Register)
	ginEngine.POST("/auth/login", userController.Login)

	// USER
	userGroup := ginEngine.Group("/users", middleware.AuthMiddleware)
	userGroup.PUT("/:userId", userController.UpdateUser)
	userGroup.DELETE("", userController.DeleteUser)

	// PHOTO
	photoGroup := ginEngine.Group("/photos", middleware.AuthMiddleware)
	photoGroup.POST("", photoController.Create)
	photoGroup.GET("", photoController.FindAll)
	photoGroup.GET("/:photoId", photoController.FindOne)
	photoGroup.PUT("/:photoId", photoController.UpdateOne)
	photoGroup.DELETE("/:photoId", photoController.Delete)

	// COMMENT
	commentGroup := ginEngine.Group("/comments", middleware.AuthMiddleware)
	commentGroup.POST("", commentController.CreateComment)
	commentGroup.GET("", commentController.GetAll)
	commentGroup.GET("/:commentId", commentController.GetOne)
	commentGroup.PUT("/:commentId", commentController.Update)
	commentGroup.DELETE("/:commentId", commentController.Delete)

	// SOCIAL MEDIA
	socialMediaGroup := ginEngine.Group("/social-media", middleware.AuthMiddleware)
	socialMediaGroup.POST("", socialMediaController.Create)
	socialMediaGroup.GET("", socialMediaController.FindAll)
	socialMediaGroup.GET("/:socialMediaId", socialMediaController.FindOne)
	socialMediaGroup.PUT("/:socialMediaId", socialMediaController.UpdateOne)
	socialMediaGroup.DELETE("/:socialMediaId", socialMediaController.Delete)

	// Host and port
	serverHost := os.Getenv("APP_HOST")
	serverPort := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(serverPort)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%d", serverHost, port)
	ginEngine.Run(addr)

}
