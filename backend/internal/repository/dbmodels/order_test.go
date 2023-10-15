package dbmodels

import (
	"order_satang/internal/models"
	"reflect"
	"testing"
)

func TestOrder_TableName(t *testing.T) {
	tests := []struct {
		name string
		o    Order
		want string
	}{
		{
			name: "success",
			o:    Order{},
			want: "orders",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.TableName(); got != tt.want {
				t.Errorf("Order.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_ToDBmodel(t *testing.T) {
	type args struct {
		data models.Order
	}
	tests := []struct {
		name string
		o    *Order
		args args
	}{
		{
			name: "success",
			o:    &Order{},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.ToDBmodel(tt.args.data)
		})
	}
}

func TestOrder_ToEntityModel(t *testing.T) {
	tests := []struct {
		name string
		o    *Order
		want *models.Order
	}{
		{
			name: "success",
			o:    &Order{},
			want: &models.Order{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.ToEntityModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.ToEntityModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
