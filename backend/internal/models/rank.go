package models

import "order_satang/util/enum"

type UserRank struct {
	UserID string
	Total  int
	Ranks  enum.Rank
}
