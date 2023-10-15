package repository

import (
	databaseMock "order_satang/database/mocks"
	"order_satang/internal/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		db         *databaseMock.MockCustomGorm
		dbBehavior func(*databaseMock.MockCustomGorm)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				db:         databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := NewOrderRepository(OrderDependencies{
				DB: tt.args.db,
			})
			assert.NotNil(t, got)
		})
	}
}

func Test_orderRepository_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		db         *databaseMock.MockCustomGorm
		dbBehavior func(*databaseMock.MockCustomGorm)
	}
	type args struct {
		req models.Order
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
				db: databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {
					mg.EXPECT().Create(gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(fooErr)
				},
			},
			args: args{
				req: models.Order{
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
				db: databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {
					mg.EXPECT().Create(gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(nil)
				},
			},
			args: args{
				req: models.Order{
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			want: &models.Order{
				ID:       0,
				UserID:   "foo",
				Product:  "foo",
				Price:    1,
				Quantity: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.dbBehavior(tt.fields.db)

		t.Run(tt.name, func(t *testing.T) {
			repo := &orderRepository{
				DB: tt.fields.db,
			}
			got, err := repo.CreateOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderRepository.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderRepository.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderRepository_UpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		db         *databaseMock.MockCustomGorm
		dbBehavior func(*databaseMock.MockCustomGorm)
	}
	type args struct {
		req models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Order
		wantErr bool
	}{
		{
			name: "error with find first",
			fields: fields{
				db: databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {
					mg.EXPECT().First(gomock.Any(), gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(fooErr)
				},
			},
			args: args{
				req: models.Order{
					ID:       0,
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
			name: "error with saving",
			fields: fields{
				db: databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {
					mg.EXPECT().First(gomock.Any(), gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(nil)

					mg.EXPECT().Save(gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(fooErr)
				},
			},
			args: args{
				req: models.Order{
					ID:       0,
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
				db: databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {
					mg.EXPECT().First(gomock.Any(), gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(nil)

					mg.EXPECT().Save(gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(nil)
				},
			},
			args: args{
				req: models.Order{
					ID:       0,
					UserID:   "foo",
					Product:  "foo",
					Price:    1,
					Quantity: 1,
				},
			},
			want: &models.Order{
				ID:       0,
				UserID:   "foo",
				Product:  "foo",
				Price:    1,
				Quantity: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.dbBehavior(tt.fields.db)

		t.Run(tt.name, func(t *testing.T) {
			repo := &orderRepository{
				DB: tt.fields.db,
			}
			got, err := repo.UpdateOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderRepository.UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderRepository.UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
