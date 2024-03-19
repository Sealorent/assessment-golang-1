package lib

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	fmt.Println("DB_USER: ", os.Getenv("DB_USER"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DBNAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	port, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, port, dbUser, dbPassword, dbName)
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})

}
