package enum

import (
	"reflect"
	"testing"
)

func TestGetRankByTotalOrder(t *testing.T) {
	type args struct {
		total int
	}
	tests := []struct {
		name string
		args args
		want Rank
	}{
		{
			name: "rank platinum",
			args: args{
				total: 100001,
			},
			want: Platinum,
		},
		{
			name: "rank gold",
			args: args{
				total: 30000,
			},
			want: Gold,
		},
		{
			name: "rank silver",
			args: args{
				total: 29999,
			},
			want: Silver,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRankByTotalOrder(tt.args.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRankByTotalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
