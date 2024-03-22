// id (Primary Key)
// user_id (Foreign Key Of User Table)
// photo_id  (Foreign Key Of Photo Table)
// message (string)
// created_at (Date)
// updated_at (Date)
package model

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	PhotoId   uint      `json:"photo_id"`
	Photo     Photo     `gorm:"foreignKey:PhotoId"`
	Message   string    `json:"message" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status" gorm:"default:true"`
}

type CommentResult struct {
	ID        uint              `json:"id"`
	Message   string            `json:"message"`
	PhotoId   uint              `json:"photo_id"`
	UserID    string            `json:"user_id"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	User      UserReferComment  `json:"user"`
	Photo     PhotoReferComment `json:"photo"`
}

type CommentResultCreate struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	PhotoId   uint   `json:"photo_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type CommentResultUpdate struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    string `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}

// Validate validates the Comment struct
func (c *Comment) Validate() error {
	var validationErrors []string
	validate := validator.New()

	if err := validate.Var(c.Message, "required"); err != nil {
		validationErrors = append(validationErrors, "Message is required")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
