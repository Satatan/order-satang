package dto

import (
	"order_satang/internal/models"
	"testing"
)

func TestUserRank_ToEntityResponse(t *testing.T) {
	type args struct {
		data models.UserRank
	}
	tests := []struct {
		name string
		ur   *UserRank
		args args
	}{
		{
			name: "success",
			ur:   &UserRank{},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ur.ToEntityResponse(tt.args.data)
		})
	}
}
