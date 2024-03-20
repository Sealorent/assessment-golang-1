package main

import (
	"final_project/controller/photo_control"
	"final_project/controller/user_control"
	"final_project/lib"
	"final_project/middleware"
	"final_project/repository/photo_repo"
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
	// err = db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})
	// if err != nil {
	// 	panic(err)
	// }

	userRepository := user_repo.NewUserRepository(db)
	userController := user_control.NewUserController(userRepository)

	photoRepository := photo_repo.NewPhotoRepository(db)
	photoController := photo_control.NewPhotoController(photoRepository)

	ginEngine := gin.Default()
	ginEngine.POST("/auth/register", userController.Register)
	ginEngine.POST("/auth/login", userController.Login)

	userGroup := ginEngine.Group("/users", middleware.AuthMiddleware)

	userGroup.PUT("", userController.UpdateUser)
	userGroup.DELETE("", userController.DeleteUser)

	photoGroup := ginEngine.Group("/photos", middleware.AuthMiddleware)
	photoGroup.POST("", photoController.Create)
	photoGroup.GET("", photoController.FindAll)

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
