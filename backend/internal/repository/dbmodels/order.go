package dbmodels

import (
	"order_satang/internal/models"
)

type Order struct {
	ID         uint64
	CustomerID string
	Product    string
	Price      int
	Quantity   int
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) ToDBmodel(data models.Order) {
	o.CustomerID = data.UserID
	o.Product = data.Product
	o.Price = data.Price
	o.Quantity = data.Quantity
}

func (o *Order) ToEntityModel() *models.Order {
	result := models.Order{}
	result.ID = o.ID
	result.UserID = o.CustomerID
	result.Product = o.Product
	result.Price = o.Price
	result.Quantity = o.Quantity
	return &result
}
