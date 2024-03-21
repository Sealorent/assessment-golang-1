package social_media_repo

import (
	"errors"
	"final_project/model"

	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (ur *socialMediaRepository) Create(socialMedia model.SocialMedia) (model.SocialMedia, error) {

	err := socialMedia.Validate()
	if err != nil {
		return model.SocialMedia{}, err
	}

	tx := ur.db.Create(&socialMedia)
	if tx.Error != nil {
		return model.SocialMedia{}, tx.Error
	}

	return socialMedia, nil

}

func (ur *socialMediaRepository) FindAll() ([]model.SocialMedia, error) {

	var socialMedia []model.SocialMedia
	tx := ur.db.Where("status = ?", true).Preload("User").Find(&socialMedia)
	if tx.Error != nil {
		return []model.SocialMedia{}, tx.Error
	}

	return socialMedia, nil
}

func (ur *socialMediaRepository) FindOne(socialMediaId string) (model.SocialMedia, error) {

	var socialMedia model.SocialMedia
	tx := ur.db.Where("id = ?", socialMediaId).Preload("User").First(&socialMedia)
	if tx.Error != nil {
		return model.SocialMedia{}, tx.Error
	}

	if !socialMedia.Status {
		return model.SocialMedia{}, errors.New("social media not found")
	}

	return socialMedia, nil

}

func (ur *socialMediaRepository) UpdateOne(socialMedia model.SocialMedia, socialMediaId string, userId string) (model.SocialMedia, error) {

	tx := ur.db.Begin()

	var existingSocialMedia model.SocialMedia
	tx.Where("id = ?", socialMediaId).First(&existingSocialMedia)

	// check if the social media status is false
	if !existingSocialMedia.Status {
		tx.Rollback()
		return model.SocialMedia{}, errors.New("social media not found")
	}
	// check if the user is the owner of the social media
	if existingSocialMedia.UserId != userId {
		tx.Rollback()
		return model.SocialMedia{}, errors.New("you don't have permission to update this social media")
	}

	// update the social media
	existingSocialMedia.Name = socialMedia.Name
	existingSocialMedia.SocialMediaUrl = socialMedia.SocialMediaUrl
	if err := tx.Save(&existingSocialMedia).Error; err != nil {
		tx.Rollback()
		return model.SocialMedia{}, err
	}

	// commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return model.SocialMedia{}, err
	}

	return existingSocialMedia, nil

}

func (ur *socialMediaRepository) Delete(socialMediaId string, userId string) error {
	tx := ur.db.Begin()

	var existingSocialMedia model.SocialMedia
	tx.Where("id = ?", socialMediaId).First(&existingSocialMedia)
	if !existingSocialMedia.Status {
		tx.Rollback()
		return errors.New("social media not found")
	}
	if existingSocialMedia.UserId != userId {
		tx.Rollback()
		return errors.New("you don't have permission to delete this social media")
	}

	existingSocialMedia.Status = false
	if err := tx.Save(&existingSocialMedia).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}
