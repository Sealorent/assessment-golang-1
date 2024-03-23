package main

import (
	"final_project/lib"
	"final_project/model"

	"gorm.io/gorm"
)

func SetupDB() (db *gorm.DB) {
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

	return db
	//
}
