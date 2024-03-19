package user_repo

import (
	"final_project/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Register(newUser model.User) (model.User, error) {

	if err := newUser.Validate(); err != nil {
		return model.User{}, err // Return empty User struct and validation error
	}

	// Proceed with registration logic
	tx := ur.db.Create(&newUser)
	if tx.Error != nil {
		return model.User{}, tx.Error // Return empty User struct and database error
	}

	return newUser, nil // Ret

}

func (ur *userRepository) UserByEmail(email string) (model.User, error) {
	var user model.User
	tx := ur.db.Where("status = ?", true).First(&user, "email = ?", email)
	return user, tx.Error
}

func (ur *userRepository) UpdateUser(updateUser model.User, id string) (model.User, error) {

	tx := ur.db.Model(&updateUser).Where("id = ?", id).Updates(updateUser)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}

	var updatedUser model.User
	err := ur.db.Where("id = ?", id).First(&updatedUser).Error
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil

}

func (ur *userRepository) DeleteUser(id string) error {
	tx := ur.db.Model(&model.User{}).Where("id = ?", id).Update("status", false)
	return tx.Error
}
