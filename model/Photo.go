package model

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Photo struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null" validate:"required"`
	UserID    string    `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status" gorm:"default:true"`
}

type PhotoResult struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	User      UserRefer `json:"user"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"updated_at"`
}

type PhotoReferComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   string `json:"user_id"`
}

// Validate validates the Photo struct
func (p *Photo) Validate() error {
	var validationErrors []string
	validate := validator.New()

	if err := validate.Var(p.Title, "required"); err != nil {
		validationErrors = append(validationErrors, "Title is required")
	}

	if err := validate.Var(p.PhotoUrl, "required"); err != nil {
		validationErrors = append(validationErrors, "Photo url is required")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}

type PhotoCreateRequestSwaggo struct {
	Title    string `json:"title" example:"example"`
	Caption  string `json:"caption" example:"example"`
	PhotoUrl string `json:"photo_url" example:"www.example.com"`
}

type PhotoUpdateRequestSwaggo struct {
	Title    string `json:"title" example:"example"`
	Caption  string `json:"caption" example:"example"`
	PhotoUrl string `json:"photo_url" example:"www.example.com"`
}
