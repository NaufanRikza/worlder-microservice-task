package usecase

import (
	"sensor-consumer/core/infrastructure/auth"
)

type authUsecase struct {
	jwtManager auth.JWTManager
}

type AuthUsecase interface {
	GenerateToken(id uint64, roles []string) (string, error)
	ValidateToken(token string) (uint, error)
}

func NewAuthUsecase(jwtManager auth.JWTManager) AuthUsecase {
	return &authUsecase{
		jwtManager: jwtManager,
	}
}

func (a *authUsecase) GenerateToken(id uint64, roles []string) (string, error) {
	return a.jwtManager.Generate(id, roles)
}

func (a *authUsecase) ValidateToken(token string) (uint, error) {
	return a.jwtManager.Validate(token)
}
