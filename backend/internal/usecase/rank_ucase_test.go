package usecase

import (
	"errors"
	domainMocks "order_satang/internal/domain/mocks"
	"order_satang/internal/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewRankUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		RankRepository         *domainMocks.MockRankRepositoryInterface
		RankRepositoryBehavior func(*domainMocks.MockRankRepositoryInterface)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				RankRepository:         domainMocks.NewMockRankRepositoryInterface(ctrl),
				RankRepositoryBehavior: func(meui *domainMocks.MockRankRepositoryInterface) {},
			},
		},
	}

	for _, tt := range tests {
		tt.args.RankRepositoryBehavior(tt.args.RankRepository)

		t.Run(tt.name, func(t *testing.T) {
			NewRankUsecase(RankDependencies{
				RankRepo: tt.args.RankRepository,
			})
		})
	}
}

func Test_rankUsecase_GetUserRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		RankRepository         *domainMocks.MockRankRepositoryInterface
		RankRepositoryBehavior func(*domainMocks.MockRankRepositoryInterface)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.UserRank
		wantErr bool
	}{
		{
			name: "error with gettting",
			fields: fields{
				RankRepository: domainMocks.NewMockRankRepositoryInterface(ctrl),
				RankRepositoryBehavior: func(meui *domainMocks.MockRankRepositoryInterface) {
					meui.EXPECT().GetUserRank().Return(nil, fooErr)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				RankRepository: domainMocks.NewMockRankRepositoryInterface(ctrl),
				RankRepositoryBehavior: func(meui *domainMocks.MockRankRepositoryInterface) {
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
			want: []models.UserRank{
				{
					UserID: "foo",
					Total:  1,
					Ranks:  "silver",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.RankRepositoryBehavior(tt.fields.RankRepository)

		t.Run(tt.name, func(t *testing.T) {
			u := rankUsecase{
				RankRepo: tt.fields.RankRepository,
			}

			got, err := u.GetUserRank()
			if (err != nil) != tt.wantErr {
				t.Errorf("orderUsecase.GetUserRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderUsecase.GetUserRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
