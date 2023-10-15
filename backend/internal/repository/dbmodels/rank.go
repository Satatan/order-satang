package dbmodels

import (
	"order_satang/internal/models"
	"order_satang/util/enum"
)

type UserRank struct {
	CustomerID string
	Total      int
}

func (UserRank) TableName() string {
	return "orders"
}

func (ur *UserRank) ToEntityModel() *models.UserRank {
	result := models.UserRank{}
	result.UserID = ur.CustomerID
	result.Total = ur.Total
	result.Ranks = enum.GetRankByTotalOrder(ur.Total)
	return &result
}
