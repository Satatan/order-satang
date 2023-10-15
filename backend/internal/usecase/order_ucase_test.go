package usecase

import (
	"errors"
	domainMocks "order_satang/internal/domain/mocks"
	"order_satang/internal/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewOrderUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		OrderRepository         *domainMocks.MockOrderRepositoryInterface
		OrderRepositoryBehavior func(*domainMocks.MockOrderRepositoryInterface)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				OrderRepository:         domainMocks.NewMockOrderRepositoryInterface(ctrl),
				OrderRepositoryBehavior: func(meui *domainMocks.MockOrderRepositoryInterface) {},
			},
		},
	}

	for _, tt := range tests {
		tt.args.OrderRepositoryBehavior(tt.args.OrderRepository)

		t.Run(tt.name, func(t *testing.T) {
			NewOrderUsecase(OrderDependencies{
				OrderRepo: tt.args.OrderRepository,
			})
		})
	}
}

func Test_orderUsecase_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		OrderRepository         *domainMocks.MockOrderRepositoryInterface
		OrderRepositoryBehavior func(*domainMocks.MockOrderRepositoryInterface)
	}
	type args struct {
		data models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Order
		wantErr bool
	}{
		{
			name: "error with creating",
			fields: fields{
				OrderRepository: domainMocks.NewMockOrderRepositoryInterface(ctrl),
				OrderRepositoryBehavior: func(meui *domainMocks.MockOrderRepositoryInterface) {
					meui.EXPECT().CreateOrder(gomock.Any()).Return(nil, fooErr)
				},
			},
			args: args{
				data: models.Order{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				OrderRepository: domainMocks.NewMockOrderRepositoryInterface(ctrl),
				OrderRepositoryBehavior: func(meui *domainMocks.MockOrderRepositoryInterface) {
					meui.EXPECT().CreateOrder(gomock.Any()).Return(
						&models.Order{
							ID:       1,
							UserID:   "foo",
							Product:  "foo",
							Price:    1,
							Quantity: 1,
						}, nil)
				},
			},
			args: args{
				data: models.Order{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			want: &models.Order{
				ID:       1,
				UserID:   "foo",
				Product:  "foo",
				Price:    1,
				Quantity: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.OrderRepositoryBehavior(tt.fields.OrderRepository)

		t.Run(tt.name, func(t *testing.T) {
			u := orderUsecase{
				OrderRepo: tt.fields.OrderRepository,
			}

			got, err := u.CreateOrder(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderUsecase.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderUsecase.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderUsecase_UpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		OrderRepository         *domainMocks.MockOrderRepositoryInterface
		OrderRepositoryBehavior func(*domainMocks.MockOrderRepositoryInterface)
	}
	type args struct {
		data models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Order
		wantErr bool
	}{
		{
			name: "error with updating",
			fields: fields{
				OrderRepository: domainMocks.NewMockOrderRepositoryInterface(ctrl),
				OrderRepositoryBehavior: func(meui *domainMocks.MockOrderRepositoryInterface) {
					meui.EXPECT().UpdateOrder(gomock.Any()).Return(nil, fooErr)
				},
			},
			args: args{
				data: models.Order{
					ID:       1,
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				OrderRepository: domainMocks.NewMockOrderRepositoryInterface(ctrl),
				OrderRepositoryBehavior: func(meui *domainMocks.MockOrderRepositoryInterface) {
					meui.EXPECT().UpdateOrder(gomock.Any()).Return(
						&models.Order{
							ID:       1,
							UserID:   "foo",
							Product:  "foo",
							Price:    1,
							Quantity: 1,
						}, nil)
				},
			},
			args: args{
				data: models.Order{
					ID:       1,
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			want: &models.Order{
				ID:       1,
				UserID:   "foo",
				Product:  "foo",
				Price:    1,
				Quantity: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.OrderRepositoryBehavior(tt.fields.OrderRepository)

		t.Run(tt.name, func(t *testing.T) {
			u := orderUsecase{
				OrderRepo: tt.fields.OrderRepository,
			}

			got, err := u.UpdateOrder(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderUsecase.UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderUsecase.UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
