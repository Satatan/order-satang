package dto

import (
	"order_satang/internal/models"
	"order_satang/util/enum"
)

type UserRank struct {
	UserID string    `json:"user_id"`
	Total  int       `json:"total"`
	Ranks  enum.Rank `json:"ranks"`
}

func (ur *UserRank) ToEntityResponse(data models.UserRank) {
	ur.UserID = data.UserID
	ur.Total = data.Total
	ur.Ranks = data.Ranks
}
