package model

type Item struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	Status      bool   `json:"status" gorm:"default:true"`
}

type ItemDTO struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
