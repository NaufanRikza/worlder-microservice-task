package usecase

import (
	"sensor-consumer/core/repository"
)

type userRoleUsecase struct {
	userRoleRepository repository.UserRoleRepository
}

type UserRoleUsecase interface {
	GetRoleByUserID(userID uint64) ([]string, error)
}

func NewUserRoleUsecase(repo repository.UserRoleRepository) UserRoleUsecase {
	return &userRoleUsecase{
		userRoleRepository: repo,
	}
}

func (u *userRoleUsecase) GetRoleByUserID(userID uint64) ([]string, error) {
	return u.userRoleRepository.GetRoleByUserID(userID)
}
