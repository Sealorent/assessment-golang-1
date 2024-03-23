package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "final_project/docs"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Load .env file
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

	}

	// Setup Database
	db := SetupDB()

	// Setup features
	features := NewSetupFeatures(db)

	// Setup routes
	ginEngine := SetupRoutes(features)

	// Host and port
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("PORT")
	port, err := strconv.Atoi(serverPort)
	if err != nil {
		panic(err)
	}

	// Address
	addr := fmt.Sprintf("%s:%d", serverHost, port)

	// Swagger
	docs.SwaggerInfo.Title = "MyGram"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = addr
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	ginEngine.Run(addr)

}
