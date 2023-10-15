package factory

import (
	"order_satang/database"
	delivery "order_satang/internal/delivery"
	repository "order_satang/internal/repository"
	usercase "order_satang/internal/usecase"

	"github.com/labstack/echo/v4"
)

func DependencyResolve(e *echo.Echo, db database.CustomGorm) {

	orderRepo := repository.NewOrderRepository(repository.OrderDependencies{
		DB: db,
	})
	rankRepo := repository.NewRankRepository(repository.RankDependencies{
		DB: db,
	})

	orderUsecase := usercase.NewOrderUsecase(usercase.OrderDependencies{
		OrderRepo: orderRepo,
	})
	rankUsecase := usercase.NewRankUsecase(usercase.RankDependencies{
		RankRepo: rankRepo,
	})

	delivery.NewOrderHandler(delivery.OrderDependencies{
		E:            e,
		OrderUsecase: orderUsecase,
	})
	delivery.NewRankHandler(delivery.RankDependencies{
		E:           e,
		RankUsecase: rankUsecase,
	})
}
