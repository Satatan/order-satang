package repository

import (
	"order_satang/database"
	"order_satang/internal/domain"
	"order_satang/internal/models"
	"order_satang/internal/repository/dbmodels"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type orderRepository struct {
	DB database.CustomGorm
}

type OrderDependencies struct {
	DB database.CustomGorm
}

func NewOrderRepository(d OrderDependencies) domain.OrderRepositoryInterface {
	return &orderRepository{
		DB: d.DB,
	}
}

func (r *orderRepository) CreateOrder(data models.Order) (*models.Order, error) {
	model := dbmodels.Order{}
	model.ToDBmodel(data)

	if err := r.DB.Create(&model).Error(); err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/repository/order_repo.go",
			"func": "CreateOrder",
		}).Error(err.Error())

		return nil, errors.WithStack(err)
	}

	result := model.ToEntityModel()

	return result, nil
}

func (r *orderRepository) UpdateOrder(data models.Order) (*models.Order, error) {

	model := &dbmodels.Order{}
	if err := r.DB.First(model, data.ID).Error(); err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/repository/order_repo.go",
			"func": "UpdateOrder",
		}).Error(err.Error())

		return nil, errors.WithStack(err)
	}

	model.ToDBmodel(data)

	if err := r.DB.Save(model).Error(); err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/repository/order_repo.go",
			"func": "UpdateOrder",
		}).Error(err.Error())

		return nil, errors.WithStack(err)
	}

	result := model.ToEntityModel()

	return result, nil
}
