package usecase

import (
	"context"
	"sensor-consumer/core/dto"
	"sensor-consumer/core/entity"
	"sensor-consumer/core/repository"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

type UserUsecase interface {
	ValidateUser(ctx context.Context, dto dto.LoginRequest) (entity.User, bool, error)
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) ValidateUser(ctx context.Context, dto dto.LoginRequest) (entity.User, bool, error) {
	user, err := u.userRepo.GetByUsername(ctx, dto.Username)
	if err != nil {
		return entity.User{}, false, err
	}
	return user, user.ID != 0 && user.Password == dto.Password && dto.Username == user.Username, nil
}
