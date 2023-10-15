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

func TestNewRankRepository(t *testing.T) {
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
			got := NewRankRepository(RankDependencies{
				DB: tt.args.db,
			})
			assert.NotNil(t, got)
		})
	}
}

func Test_rankRepository_GetUserRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		db         *databaseMock.MockCustomGorm
		dbBehavior func(*databaseMock.MockCustomGorm)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.UserRank
		wantErr bool
	}{
		{
			name: "error with getting",
			fields: fields{
				db: databaseMock.NewMockCustomGorm(ctrl),
				dbBehavior: func(mg *databaseMock.MockCustomGorm) {
					mg.EXPECT().Group(gomock.Any()).Return(mg)
					mg.EXPECT().Select(gomock.Any()).Return(mg)
					mg.EXPECT().Find(gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(fooErr)
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
					mg.EXPECT().Group(gomock.Any()).Return(mg)
					mg.EXPECT().Select(gomock.Any()).Return(mg)
					mg.EXPECT().Find(gomock.Any()).Return(mg)
					mg.EXPECT().Error().Return(nil)
				},
			},
			want:    []models.UserRank{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.dbBehavior(tt.fields.db)

		t.Run(tt.name, func(t *testing.T) {
			repo := &rankRepository{
				DB: tt.fields.db,
			}
			got, err := repo.GetUserRank()
			if (err != nil) != tt.wantErr {
				t.Errorf("rankRepository.GetUserRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rankRepository.GetUserRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
