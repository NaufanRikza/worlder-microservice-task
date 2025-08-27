package usecase

import (
	"sensor-consumer/core/infrastructure/auth"
)

type authUsecase struct {
	jwtManager     auth.JWTManager
	passwordHasher auth.PasswordHasher
}

type AuthUsecase interface {
	GenerateToken(id uint64) (string, error)
	ValidateToken(token string) (uint, error)
}

func NewAuthUsecase(jwtManager auth.JWTManager, passwordHasher auth.PasswordHasher) AuthUsecase {
	return &authUsecase{
		jwtManager:     jwtManager,
		passwordHasher: passwordHasher,
	}
}

func (a *authUsecase) GenerateToken(id uint64) (string, error) {
	return a.jwtManager.Generate(id)
}

func (a *authUsecase) ValidateToken(token string) (uint, error) {
	return a.jwtManager.Validate(token)
}
