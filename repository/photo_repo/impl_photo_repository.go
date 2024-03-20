package photo_repo

import "final_project/model"

type IPhotoRepository interface {
	Create(photo model.Photo) (model.Photo, error)
	FindAll() ([]model.Photo, error)
	UpdateOne(photo model.Photo, photoId string, userId string) (model.Photo, error)
	Delete(photoId string, userId string) error
}
