package dto

import (
	"order_satang/internal/models"
)

type OrderRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Product  string `json:"product" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type Order struct {
	ID       uint64 `json:"id,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	Product  string `json:"product,omitempty"`
	Price    int    `json:"price,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

func (o *OrderRequest) ToEntityModel() *models.Order {
	result := models.Order{}
	result.UserID = o.UserID
	result.Product = o.Product
	result.Price = o.Price
	result.Quantity = o.Quantity
	return &result
}

func (o *Order) ToEntityCreateResponse(data models.Order) {
	o.ID = data.ID
}

func (o *Order) ToEntityResponse(data models.Order) {
	o.ID = data.ID
	o.UserID = data.UserID
	o.Product = data.Product
	o.Price = data.Price
	o.Quantity = data.Quantity
}
