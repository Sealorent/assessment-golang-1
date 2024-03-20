// id (Primary key)

// title (string)

// caption (string)

// photo_url (string)

// user_id (Foreign Key Of User Table)

// created_at (Date)

// updated_at (Date)

// Photo is a struct that represents the photo model
package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Photo struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null" validate:"required"`
	UserId    string    `json:"user_id" gorm:"foreignKey:UserRefer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validates the Photo struct
func (p *Photo) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
