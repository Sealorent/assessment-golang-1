// routes.go
package main

import (
	"final_project/middleware"
	"final_project/views"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	controllers *SetupFeatures,
) *gin.Engine {
	ginEngine := gin.Default()

	//INDEX
	ginEngine.GET("/", func(c *gin.Context) {
		// Call the Index function directly
		views.Index(c.Writer, c.Request)
	})

	// AUTH
	ginEngine.POST("/auth/register", controllers.UserController.Register)
	ginEngine.POST("/auth/login", controllers.UserController.Login)

	// USER
	userGroup := ginEngine.Group("/users", middleware.AuthMiddleware)
	userGroup.PUT("/:userId", controllers.UserController.UpdateUser)
	userGroup.DELETE("", controllers.UserController.DeleteUser)

	// PHOTO
	photoGroup := ginEngine.Group("/photos", middleware.AuthMiddleware)
	photoGroup.POST("", controllers.PhotoController.Create)
	photoGroup.GET("", controllers.PhotoController.FindAll)
	photoGroup.GET("/:photoId", controllers.PhotoController.FindOne)
	photoGroup.PUT("/:photoId", controllers.PhotoController.UpdateOne)
	photoGroup.DELETE("/:photoId", controllers.PhotoController.Delete)

	// COMMENT
	commentGroup := ginEngine.Group("/comments", middleware.AuthMiddleware)
	commentGroup.POST("", controllers.CommentController.CreateComment)
	commentGroup.GET("", controllers.CommentController.GetAll)
	commentGroup.GET("/:commentId", controllers.CommentController.GetOne)
	commentGroup.PUT("/:commentId", controllers.CommentController.Update)
	commentGroup.DELETE("/:commentId", controllers.CommentController.Delete)

	// SOCIAL MEDIA
	socialMediaGroup := ginEngine.Group("/social-media", middleware.AuthMiddleware)
	socialMediaGroup.POST("", controllers.SocialMediaController.Create)
	socialMediaGroup.GET("", controllers.SocialMediaController.FindAll)
	socialMediaGroup.GET("/:socialMediaId", controllers.SocialMediaController.FindOne)
	socialMediaGroup.PUT("/:socialMediaId", controllers.SocialMediaController.UpdateOne)
	socialMediaGroup.DELETE("/:socialMediaId", controllers.SocialMediaController.Delete)

	return ginEngine
}
