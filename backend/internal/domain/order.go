package domain

import "order_satang/internal/models"

type OrderUsecaseInterface interface {
	CreateOrder(data models.Order) (*models.Order, error)
	UpdateOrder(data models.Order) (*models.Order, error)
}

type OrderRepositoryInterface interface {
	CreateOrder(data models.Order) (*models.Order, error)
	UpdateOrder(data models.Order) (*models.Order, error)
}
