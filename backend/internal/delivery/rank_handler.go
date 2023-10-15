package delivery

import (
	"net/http"
	"order_satang/internal/delivery/dto"
	"order_satang/internal/domain"
	"order_satang/util/enum"

	"github.com/labstack/echo/v4"
)

type rankHandler struct {
	RankUsecase domain.RankUsecaseInterface
}

type RankDependencies struct {
	E           *echo.Echo
	RankUsecase domain.RankUsecaseInterface
}

func NewRankHandler(d RankDependencies) {
	handler := &rankHandler{
		RankUsecase: d.RankUsecase,
	}

	api := d.E.Group("/api")
	v1 := api.Group("/v1")

	rankGroup := v1.Group("/ranks")
	rankGroup.GET("", handler.GetUserRank)
}

func (h *rankHandler) GetUserRank(c echo.Context) error {

	userRanks, err := h.RankUsecase.GetUserRank()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.CoreResponse{
			Message: enum.MessageError,
			Result:  err.Error(),
		})

	}

	result := []dto.UserRank{}
	for _, record := range userRanks {
		entity := dto.UserRank{}
		entity.ToEntityResponse(record)
		result = append(result, entity)
	}

	res := dto.CoreResponse{
		Message: enum.MessageSuccess,
		Result:  result,
	}

	return c.JSON(http.StatusOK, res)
}
