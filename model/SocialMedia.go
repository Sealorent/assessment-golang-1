// id (Primary Key)
// name (String/ varchar)
// social_media_url (String/ Text)
// UserId(Foreign Key Of User Table)
// created_at (Date)
// updated_at (Date)

// SocialMedia is a struct that represents the social_media model
package model

import (
	"errors"
	"time"
)

type SocialMedia struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null" validate:"required"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"not null" validate:"required"`
	UserId         string    `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Status         bool      `json:"status" gorm:"default:true"`
}

// Validate validates the SocialMedia struct
func (s *SocialMedia) Validate() error {

	var stringError string = ""

	if s.Name == "" {
		stringError += "name is required. "
	}

	if s.SocialMediaUrl == "" {
		stringError += "social media url is required. "
	}

	if stringError != "" {
		return errors.New(stringError)
	}

	return nil
}

type SocialMediaResultCreated struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         string `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}

type SocialMediaResult struct {
	ID             uint                 `json:"id"`
	Name           string               `json:"name"`
	SocialMediaUrl string               `json:"social_media_url"`
	UserId         string               `json:"user_id"`
	CreatedAt      string               `json:"created_at"`
	UpdatedAt      string               `json:"updated_at"`
	User           UserReferSocialMedia `json:"user"`
}

type SocialMediaResultUpdated struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         string `json:"user_id"`
	UpdatedAt      string `json:"updated_at"`
}
