package model

import (
	"errors"
	"strings"
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

type UserRefer struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserReferComment struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserReferSocialMedia struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// create uuid for user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

// Validate validates the User struct
func (u *User) Validate() error {
	var validationErrors []string
	validate := validator.New()

	if err := validate.Var(u.Username, "required"); err != nil {
		validationErrors = append(validationErrors, "Username is required")
	}

	if err := validate.Var(u.Email, "required,email"); err != nil {
		validationErrors = append(validationErrors, "Email is required and must be valid")
	}

	if err := validate.Var(u.Password, "required,min=6"); err != nil {
		validationErrors = append(validationErrors, "Password is required and must have at least 6 characters")
	}

	if err := validate.Var(u.Age, "required,min=9"); err != nil {
		validationErrors = append(validationErrors, "Age is required and must be at least 8")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
