package photo_repo

import (
	"errors"
	"final_project/model"
	"time"

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

	err := photo.Validate()
	if err != nil {
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
	// add where status = true with preloading user
	tx := ur.db.Where("status = ?", true).Preload("User").Find(&photos)
	if tx.Error != nil {
		return []model.Photo{}, tx.Error
	}

	return photos, nil
}

func (ur *photoRepository) UpdateOne(photo model.Photo, photoId string, userId string) (model.Photo, error) {

	// Make begin transaction
	tx := ur.db.Begin()

	// 1. Find photo by id and user_id
	var existingPhoto model.Photo
	if err := tx.Where("id = ? AND user_id = ? AND status = true", photoId, userId).First(&existingPhoto).Error; err != nil {
		tx.Rollback()
		return model.Photo{}, errors.New("photo not found or you don't have permission to update this photo")
	}

	// 2. Update photo
	existingPhoto.Title = photo.Title
	existingPhoto.Caption = photo.Caption
	existingPhoto.PhotoUrl = photo.PhotoUrl
	existingPhoto.UpdatedAt = time.Now() // You may want to update the 'UpdatedAt' field
	if err := tx.Save(&existingPhoto).Error; err != nil {
		tx.Rollback()
		return model.Photo{}, err
	}

	// 3. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return model.Photo{}, err
	}

	// 4. Return updated photo
	return existingPhoto, nil

}

func (ur *photoRepository) Delete(photoId string, userId string) error {

	tx := ur.db.Begin()

	// 1. Find photo by id and user_id
	var existingPhoto model.Photo
	if err := tx.Where("id = ? AND user_id = ?", photoId, userId).First(&existingPhoto).Error; err != nil {
		tx.Rollback()
		return errors.New("photo not found or you don't have permission to delete this photo")
	}

	// 2. Delete photo using soft delete
	existingPhoto.Status = false
	if err := tx.Save(&existingPhoto).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}
