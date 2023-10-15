package domain

import "order_satang/internal/models"

type RankUsecaseInterface interface {
	GetUserRank() ([]models.UserRank, error)
}

type RankRepositoryInterface interface {
	GetUserRank() ([]models.UserRank, error)
}
