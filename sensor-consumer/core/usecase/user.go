package usecase

import (
	"context"
	"sensor-consumer/core/dto"
	"sensor-consumer/core/entity"
	"sensor-consumer/core/infrastructure/auth"
	"sensor-consumer/core/repository"
)

type userUsecase struct {
	userRepo       repository.UserRepository
	passwordHasher auth.PasswordHasher
}

type UserUsecase interface {
	ValidateUser(ctx context.Context, dto dto.LoginRequest) (entity.User, bool, error)
}

func NewUserUsecase(userRepo repository.UserRepository, passwordHasher auth.PasswordHasher) UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (u *userUsecase) ValidateUser(ctx context.Context, dto dto.LoginRequest) (entity.User, bool, error) {
	user, err := u.userRepo.GetByUsername(ctx, dto.Username)
	if err != nil {
		return user, false, err
	}

	err = u.passwordHasher.CheckPasswordHash(user.Password, dto.Password)
	return user, user.ID != 0 && dto.Username == user.Username && err == nil, err
}
