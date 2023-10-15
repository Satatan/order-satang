package delivery

import (
	"net/http"
	"order_satang/internal/delivery/dto"
	"order_satang/internal/domain"
	"order_satang/util"
	"order_satang/util/enum"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type orderHandler struct {
	OrderUsecase domain.OrderUsecaseInterface
}

type OrderDependencies struct {
	E            *echo.Echo
	OrderUsecase domain.OrderUsecaseInterface
}

func NewOrderHandler(d OrderDependencies) {
	handler := &orderHandler{
		OrderUsecase: d.OrderUsecase,
	}

	api := d.E.Group("/api")
	v1 := api.Group("/v1")

	orderGroup := v1.Group("/orders")
	orderGroup.POST("", handler.CreateOrder)
	orderGroup.PUT("/:id", handler.UpdateOrder)
}

func (h *orderHandler) CreateOrder(c echo.Context) error {

	req := dto.OrderRequest{}

	err := c.Bind(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/delivery/order_handler.go",
			"func": "CreateOrder",
		}).Error(err.Error())

		return c.JSON(http.StatusBadRequest, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  err.Error(),
		})
	}

	errValidate := util.ValidateInputs(&req)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  errValidate,
		})
	}

	data := req.ToEntityModel()
	order, err := h.OrderUsecase.CreateOrder(*data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  err.Error(),
		})

	}

	result := dto.Order{}
	result.ToEntityCreateResponse(*order)

	res := dto.CoreResponse{
		Message: enum.MessageSuccess,
		Result:  result,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *orderHandler) UpdateOrder(c echo.Context) error {

	req := dto.OrderRequest{}

	err := c.Bind(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/delivery/order_handler.go",
			"func": "CreateOrder",
		}).Error(err.Error())

		return c.JSON(http.StatusBadRequest, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  err.Error(),
		})
	}

	errValidate := util.ValidateInputs(&req)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  errValidate,
		})
	}

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"path": "internal/delivery/order_handler.go",
			"func": "CreateOrder",
		}).Error(err.Error())

		return c.JSON(http.StatusBadRequest, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  err.Error(),
		})
	}

	data := req.ToEntityModel()
	data.ID = id
	order, err := h.OrderUsecase.UpdateOrder(*data)
	if err != nil {
		if err.Error() == enum.MessageStatusForbidden {
			return c.JSON(http.StatusForbidden, dto.CoreResponse{
				Message: enum.MessageError,
				Result:  err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  err.Error(),
		})
	}

	result := dto.Order{}
	result.ToEntityResponse(*order)

	res := dto.CoreResponse{
		Message: enum.MessageSuccess,
		Result:  result,
	}

	return c.JSON(http.StatusOK, res)
}
