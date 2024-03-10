package repository

import (
	"fmt"
	"project-pertama/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) Create(newOrder model.Order) (model.Order, error) {

	// Assign values to the existing newOrder variable
	newsOrder := model.Order{
		CustomerName: newOrder.CustomerName,
		OrderedAt:    newOrder.OrderedAt,
	}

	// Store the order in the database
	if err := or.db.Create(&newsOrder).Error; err != nil {
		// Handle error
		return model.Order{}, err
	}

	// Prepare items for the order
	var items []model.Item
	var resultItems []model.Item
	for _, itemDTO := range newOrder.Item {
		item := model.Item{
			ItemCode:    itemDTO.ItemCode,
			Description: itemDTO.Description,
			Quantity:    itemDTO.Quantity,
			OrderID:     newsOrder.ID,
		}
		items = append(items, item)
	}

	fmt.Println(items)

	// Store the items in the database
	for _, item := range items {
		if err := or.db.Create(&item).Error; err != nil {
			// Handle error
			return model.Order{}, err
		}
		resultItems = append(resultItems, item)
	}

	// Prepare the response
	orderResponse := model.Order{
		ID:           newsOrder.ID,
		CustomerName: newsOrder.CustomerName,
		OrderedAt:    newsOrder.OrderedAt,
		Status:       newsOrder.Status,
		Item:         resultItems,
	}

	// Return the response
	return orderResponse, nil
}

func (or *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	// Get all orders from the database and preload the items where the status is true
	if err := or.db.Where("status = ?", true).Preload("Item", "status = ?", true).Find(&orders).Error; err != nil {
		// Handle error
		return []model.Order{}, err
	}

	// Return the response
	return orders, nil
}

func (or *OrderRepository) Update(id uint, updatedOrder model.Order) (model.Order, error) {

	// Get the order from the database
	var existingOrder model.Order
	if err := or.db.Where("id = ?", id).Preload("Item", "status = ?", true).First(&existingOrder).Error; err != nil {
		// Handle error
		return model.Order{}, err
	}

	// update the order
	existingOrder.CustomerName = updatedOrder.CustomerName
	existingOrder.OrderedAt = updatedOrder.OrderedAt
	// update the order in the database
	if err := or.db.Save(&existingOrder).Error; err != nil {
		// Handle error
		return model.Order{}, err
	}

	// make for loop to update the items
	var updateItems []model.Item
	for _, item := range updatedOrder.Item {
		// check if the item is already exist
		var existingItem model.Item
		if err := or.db.Where("id = ?", item.ID).First(&existingItem).Error; err != nil {
			// Handle error
			return model.Order{}, err
		}
		// update the item

		if item.ItemCode != "" && item.ItemCode != existingItem.ItemCode {
			existingItem.ItemCode = item.ItemCode
		}

		if item.Description != "" && item.Description != existingItem.Description {
			existingItem.Description = item.Description
		}

		if item.Quantity != existingItem.Quantity && item.Quantity != 0 {
			existingItem.Quantity = item.Quantity
		}

		// update the item in the database
		result := or.db.Model(&existingItem).Select("item_code", "description", "quantity").Updates(&existingItem)
		if result.Error != nil {
			// Handle error
			return model.Order{}, result.Error
		}

		updateItems = append(updateItems, existingItem)

	}

	// Prepare the response
	existingOrder.Item = updateItems

	// Return the response
	return existingOrder, nil
}

func (or *OrderRepository) Delete(id uint) (model.Order, error) {
	// Get the order from the database
	var existingOrder model.Order
	if err := or.db.Where("id = ?", id).Preload("Item", "status = ?", true).First(&existingOrder).Error; err != nil {
		// Handle error
		return model.Order{}, err
	}

	// update the status of the order to false
	existingOrder.Status = false
	// update the order in the database
	if err := or.db.Save(&existingOrder).Error; err != nil {
		// Handle error
		return model.Order{}, err
	}

	// make for loop to update the items
	var deleteItems []model.Item
	for _, item := range existingOrder.Item {
		// update the status of the item to false
		item.Status = false
		// update the item in the database
		if err := or.db.Save(&item).Error; err != nil {
			// Handle error
			return model.Order{}, err
		}

		deleteItems = append(deleteItems, item)
	}

	// Prepare the response
	existingOrder.Item = deleteItems

	// Return the response
	return existingOrder, nil
}
