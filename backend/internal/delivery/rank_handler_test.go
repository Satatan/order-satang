package delivery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	domainMocks "order_satang/internal/domain/mocks"
	"order_satang/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewRankHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		E                   *echo.Echo
		RankUsecase         *domainMocks.MockRankUsecaseInterface
		RankUsecaseBehavior func(*domainMocks.MockRankUsecaseInterface)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				E:                   echo.New(),
				RankUsecase:         domainMocks.NewMockRankUsecaseInterface(ctrl),
				RankUsecaseBehavior: func(meui *domainMocks.MockRankUsecaseInterface) {},
			},
		},
	}

	for _, tt := range tests {
		tt.args.RankUsecaseBehavior(tt.args.RankUsecase)

		t.Run(tt.name, func(t *testing.T) {
			NewRankHandler(RankDependencies{
				E:           tt.args.E,
				RankUsecase: tt.args.RankUsecase,
			})
		})
	}
}

func Test_rankHandler_GetUserRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		RankUsecase         *domainMocks.MockRankUsecaseInterface
		RankUsecaseBehavior func(*domainMocks.MockRankUsecaseInterface)
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		statusCode int
	}{
		{
			name: "error with getting",
			fields: fields{
				RankUsecase: domainMocks.NewMockRankUsecaseInterface(ctrl),
				RankUsecaseBehavior: func(meui *domainMocks.MockRankUsecaseInterface) {
					meui.EXPECT().GetUserRank().Return(nil, fooErr)
				},
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			fields: fields{
				RankUsecase: domainMocks.NewMockRankUsecaseInterface(ctrl),
				RankUsecaseBehavior: func(meui *domainMocks.MockRankUsecaseInterface) {
					meui.EXPECT().GetUserRank().Return(
						[]models.UserRank{
							{
								UserID: "foo",
								Total:  1,
								Ranks:  "silver",
							},
						}, nil)
				},
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {

		tt.fields.RankUsecaseBehavior(tt.fields.RankUsecase)

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/ranks", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		tt.args.c = e.NewContext(req, rec)

		t.Run(tt.name, func(t *testing.T) {
			h := rankHandler{
				RankUsecase: tt.fields.RankUsecase,
			}

			err := h.GetUserRank(tt.args.c)
			if err != nil {
				httpError := err.(*echo.HTTPError)
				assert.Equal(t, tt.statusCode, httpError.Code)
			} else {
				assert.Equal(t, tt.statusCode, rec.Code)
			}
		})
	}
}
