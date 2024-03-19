// id (Primary Key)
// name (String/ varchar)
// social_media_url (String/ Text)
// UserId(Foreign Key Of User Table)
// created_at (Date)
// updated_at (Date)

// SocialMedia is a struct that represents the social_media model
package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMedia struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null" validate:"required"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"not null" validate:"required"`
	UserId         uint      `json:"user_id" gorm:"foreignKey:UserRefer"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Validate validates the SocialMedia struct
func (s *SocialMedia) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
