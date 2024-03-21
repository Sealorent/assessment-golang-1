package user_repo

import (
	"final_project/model"
)

type IUserRepository interface {
	Register(newUser model.User) (model.User, error)
	UserByEmail(string) (model.User, error)
	UpdateUser(updateUser model.User, id string) (model.User, error)
	DeleteUser(id string) error
}
