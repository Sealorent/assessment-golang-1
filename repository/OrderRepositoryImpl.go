package repository

import "project-pertama/model"

type OrderRepositoryImpl interface {
	Create(newOrder model.Order) (model.Order, error)
	GetAll() ([]model.Order, error)
	Update(id uint, order model.Order) (model.Order, error)
	Delete(id uint) (model.Order, error)
}
