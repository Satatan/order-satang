package usecase

import (
	"order_satang/internal/domain"
	"order_satang/internal/models"
)

type rankUsecase struct {
	RankRepo domain.RankRepositoryInterface
}

type RankDependencies struct {
	RankRepo domain.RankRepositoryInterface
}

func NewRankUsecase(d RankDependencies) domain.RankUsecaseInterface {
	return &rankUsecase{
		RankRepo: d.RankRepo,
	}
}

func (u *rankUsecase) GetUserRank() ([]models.UserRank, error) {
	result, err := u.RankRepo.GetUserRank()
	if err != nil {
		return nil, err
	}
	return result, nil
}
