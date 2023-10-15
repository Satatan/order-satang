package dto

import (
	"order_satang/internal/models"
	"reflect"
	"testing"
)

func TestOrderRequest_ToEntityModel(t *testing.T) {
	tests := []struct {
		name string
		o    *OrderRequest
		want *models.Order
	}{
		{
			name: "success",
			o:    &OrderRequest{},
			want: &models.Order{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.ToEntityModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRequest.ToEntityModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_ToEntityCreateResponse(t *testing.T) {
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
			tt.o.ToEntityCreateResponse(tt.args.data)
		})
	}
}

func TestOrder_ToEntityResponse(t *testing.T) {
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
			tt.o.ToEntityResponse(tt.args.data)
		})
	}
}
