package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Username  string    `json:"username" gorm:"unique;not null" validate:"required"`
	Password  string    `json:"password" validate:"required,min=6"`
	Age       uint      `json:"age" validate:"required,min=9"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status" gorm:"default:true"`
}

// create uuid for user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

// Validate validates the User struct
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
