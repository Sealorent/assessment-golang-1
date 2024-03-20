package photo_repo

import (
	"final_project/model"

	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		db: db,
	}
}

func (ur *photoRepository) Create(photo model.Photo) (model.Photo, error) {

	if err := photo.Validate(); err != nil {
		return model.Photo{}, err
	}

	tx := ur.db.Create(&photo)
	if tx.Error != nil {
		return model.Photo{}, tx.Error
	}

	return photo, nil

}

func (ur *photoRepository) FindAll() ([]model.Photo, error) {

	var photos []model.Photo
	tx := ur.db.Preload("User").Find(&photos)
	if tx.Error != nil {
		return []model.Photo{}, tx.Error
	}
	return photos, nil
}
