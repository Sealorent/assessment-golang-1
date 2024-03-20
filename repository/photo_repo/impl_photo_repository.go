package photo_repo

import "final_project/model"

type IPhotoRepository interface {
	Create(photo model.Photo) (model.Photo, error)
	FindAll() ([]model.Photo, error)
}
