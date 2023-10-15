package repository

import (
	"order_satang/database"
	"order_satang/internal/domain"
	"order_satang/internal/models"
	"order_satang/internal/repository/dbmodels"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type rankRepository struct {
	DB database.CustomGorm
}

type RankDependencies struct {
	DB database.CustomGorm
}

func NewRankRepository(d RankDependencies) domain.RankRepositoryInterface {
	return &rankRepository{
		DB: d.DB,
	}
}

func (r *rankRepository) GetUserRank() ([]models.UserRank, error) {
	userRankModels := []dbmodels.UserRank{}

	db := r.DB.Group("customer_id").
		Select("customer_id, sum(price*quantity) as total")

	err := db.Find(&userRankModels).Error()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/repository/rank_repo.go",
			"func": "GetUserRank",
		}).Error(err.Error())

		return nil, errors.WithStack(err)
	}

	result := []models.UserRank{}
	for _, record := range userRankModels {
		result = append(result, *record.ToEntityModel())
	}

	return result, nil
}
