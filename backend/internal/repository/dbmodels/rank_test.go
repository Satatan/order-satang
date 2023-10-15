package dbmodels

import (
	"order_satang/internal/models"
	"reflect"
	"testing"
)

func TestUserRank_TableName(t *testing.T) {
	tests := []struct {
		name string
		u    UserRank
		want string
	}{
		{
			name: "success",
			u:    UserRank{},
			want: "orders",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.TableName(); got != tt.want {
				t.Errorf("UserRank.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRank_ToEntityModel(t *testing.T) {
	tests := []struct {
		name string
		ur   *UserRank
		want *models.UserRank
	}{
		{
			name: "success",
			ur: &UserRank{
				CustomerID: "foo",
				Total:      1,
			},
			want: &models.UserRank{
				UserID: "foo",
				Total:  1,
				Ranks:  "silver",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ur.ToEntityModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRank.ToEntityModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
