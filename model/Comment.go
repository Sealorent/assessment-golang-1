// id (Primary Key)
// user_id (Foreign Key Of User Table)
// photo_id  (Foreign Key Of Photo Table)
// message (string)
// created_at (Date)
// updated_at (Date)
package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"foreignKey:UserRefer" validate:"required"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validates the Comment struct
func (c *Comment) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
