package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtManager struct {
	secretKey string
}

type JWTManager interface {
	Generate(id uint64) (string, error)
	Validate(token string) (uint, error)
}

func NewJWTManager(secretKey string) JWTManager {
	return &jwtManager{
		secretKey: secretKey,
	}
}

func (j *jwtManager) Generate(id uint64) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtManager) Validate(token string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	id, ok := claims["id"].(uint)
	if !ok {
		return 0, errors.New("invalid token")
	}

	exp, ok := claims["exp"].(int64)
	if !ok || exp < time.Now().Unix() {
		return 0, errors.New("token expired")
	}

	return id, nil
}
