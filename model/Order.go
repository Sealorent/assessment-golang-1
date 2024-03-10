package model

import (
	"time"
)

type Order struct {
	ID           uint      `json:"id" gorm:"primary_key;type:serial;AUTO_INCREMENT"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Item         []Item    `json:"items" gorm:"foreignKey:OrderID;references:ID"`
	Status       bool      `json:"status" gorm:"default:true"`
}

func OrderHandler() {

	type request struct {
		CustomerName string    `json:"customer_name"`
		OrderedAt    string    `json:"ordered_at"`
		Item         []ItemDTO `json:"items"`
	}

	type response struct {
		ID           uint   `json:"id"`
		CustomerName string `json:"customer_name"`
		OrderedAt    string `json:"ordered_at"`
		Item         []Item `json:"items"`
	}
}
