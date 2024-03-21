package social_media_repo

import "final_project/model"

type ISocialMediaRepository interface {
	Create(SocialMedia model.SocialMedia) (model.SocialMedia, error)
	FindAll() ([]model.SocialMedia, error)
	FindOne(SocialMediaId string) (model.SocialMedia, error)
	UpdateOne(SocialMedia model.SocialMedia, SocialMediaId string, userId string) (model.SocialMedia, error)
	Delete(SocialMediaId string, userId string) error
}
