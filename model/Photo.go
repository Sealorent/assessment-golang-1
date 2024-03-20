package model

import (
	"errors"
	"time"
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

	var stringError string = ""

	if p.Title == "" {
		stringError += "title is required. "
	}

	if p.PhotoUrl == "" {
		stringError += "photo url is required. "
	}

	if p.UserID == "" {
		stringError += "please sign-in. "
	}

	if stringError != "" {
		return errors.New(stringError)
	}

	return nil
}
