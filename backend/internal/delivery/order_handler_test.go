package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"order_satang/internal/delivery/dto"
	domainMocks "order_satang/internal/domain/mocks"
	"order_satang/internal/models"
	"order_satang/util/enum"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		E                    *echo.Echo
		OrderUsecase         *domainMocks.MockOrderUsecaseInterface
		OrderUsecaseBehavior func(*domainMocks.MockOrderUsecaseInterface)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				E:                    echo.New(),
				OrderUsecase:         domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {},
			},
		},
	}

	for _, tt := range tests {
		tt.args.OrderUsecaseBehavior(tt.args.OrderUsecase)

		t.Run(tt.name, func(t *testing.T) {
			NewOrderHandler(OrderDependencies{
				E:            tt.args.E,
				OrderUsecase: tt.args.OrderUsecase,
			})
		})
	}
}

func Test_orderHandler_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		OrderUsecase         *domainMocks.MockOrderUsecaseInterface
		OrderUsecaseBehavior func(*domainMocks.MockOrderUsecaseInterface)
	}
	type httpBuilder struct {
		body interface{}
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name        string
		fields      fields
		httpBuilder httpBuilder
		args        args
		statusCode  int
	}{
		{
			name: "error with binding",
			fields: fields{
				OrderUsecase:         domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {},
			},
			httpBuilder: httpBuilder{
				body: "foo",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error with validatide",
			fields: fields{
				OrderUsecase:         domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:  "foo",
					Product: "foo",
					Price:   1,
				},
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error with creating",
			fields: fields{
				OrderUsecase: domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {
					meui.EXPECT().CreateOrder(gomock.Any()).Return(nil, fooErr)
				},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			fields: fields{
				OrderUsecase: domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {
					meui.EXPECT().CreateOrder(gomock.Any()).Return(
						&models.Order{
							ID: 1,
						}, nil)
				},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {

		tt.fields.OrderUsecaseBehavior(tt.fields.OrderUsecase)

		e := echo.New()
		b, err := json.Marshal(tt.httpBuilder.body)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		tt.args.c = e.NewContext(req, rec)

		t.Run(tt.name, func(t *testing.T) {
			h := orderHandler{
				OrderUsecase: tt.fields.OrderUsecase,
			}

			err := h.CreateOrder(tt.args.c)
			if err != nil {
				httpError := err.(*echo.HTTPError)
				assert.Equal(t, tt.statusCode, httpError.Code)
			} else {
				assert.Equal(t, tt.statusCode, rec.Code)
			}
		})
	}
}

func Test_orderHandler_UpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		OrderUsecase         *domainMocks.MockOrderUsecaseInterface
		OrderUsecaseBehavior func(*domainMocks.MockOrderUsecaseInterface)
	}
	type httpBuilder struct {
		body interface{}
		id   string
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name        string
		fields      fields
		httpBuilder httpBuilder
		args        args
		statusCode  int
	}{
		{
			name: "error with binding",
			fields: fields{
				OrderUsecase:         domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {},
			},
			httpBuilder: httpBuilder{
				body: "foo",
				id:   "1",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error with validatide",
			fields: fields{
				OrderUsecase:         domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:  "foo",
					Product: "foo",
					Price:   1,
				},
				id: "1",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error with id is not number",
			fields: fields{
				OrderUsecase:         domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
				id: "foo",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "error no permission",
			fields: fields{
				OrderUsecase: domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {
					meui.EXPECT().UpdateOrder(gomock.Any()).Return(nil, errors.New(enum.MessageStatusForbidden))
				},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
				id: "1",
			},
			statusCode: http.StatusForbidden,
		},
		{
			name: "error with updating",
			fields: fields{
				OrderUsecase: domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {
					meui.EXPECT().UpdateOrder(gomock.Any()).Return(nil, fooErr)
				},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
				id: "1",
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "success",
			fields: fields{
				OrderUsecase: domainMocks.NewMockOrderUsecaseInterface(ctrl),
				OrderUsecaseBehavior: func(meui *domainMocks.MockOrderUsecaseInterface) {
					meui.EXPECT().UpdateOrder(gomock.Any()).Return(
						&models.Order{
							ID: 1,
						}, nil)
				},
			},
			httpBuilder: httpBuilder{
				body: dto.OrderRequest{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
				id: "1",
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {

		tt.fields.OrderUsecaseBehavior(tt.fields.OrderUsecase)

		e := echo.New()
		b, err := json.Marshal(tt.httpBuilder.body)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/orders/:id", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		tt.args.c = e.NewContext(req, rec)

		tt.args.c.SetParamNames("id")
		tt.args.c.SetParamValues(tt.httpBuilder.id)

		t.Run(tt.name, func(t *testing.T) {
			h := orderHandler{
				OrderUsecase: tt.fields.OrderUsecase,
			}

			err := h.UpdateOrder(tt.args.c)
			if err != nil {
				httpError := err.(*echo.HTTPError)
				assert.Equal(t, tt.statusCode, httpError.Code)
			} else {
				assert.Equal(t, tt.statusCode, rec.Code)
			}
		})
	}
}
