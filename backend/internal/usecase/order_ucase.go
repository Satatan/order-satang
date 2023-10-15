package usecase

import (
	"order_satang/internal/domain"
	"order_satang/internal/models"
)

type orderUsecase struct {
	OrderRepo domain.OrderRepositoryInterface
}

type OrderDependencies struct {
	OrderRepo domain.OrderRepositoryInterface
}

func NewOrderUsecase(d OrderDependencies) domain.OrderUsecaseInterface {
	return &orderUsecase{
		OrderRepo: d.OrderRepo,
	}
}

func (u *orderUsecase) CreateOrder(data models.Order) (*models.Order, error) {
	result, err := u.OrderRepo.CreateOrder(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *orderUsecase) UpdateOrder(data models.Order) (*models.Order, error) {
	result, err := u.OrderRepo.UpdateOrder(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
