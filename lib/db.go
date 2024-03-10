package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	connectionString := "host=173.212.232.47 port=2670 user=postgres password=postgres dbname=db-golang sslmode=disable"
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
