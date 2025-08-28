package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	UserID uint64   `json:"user_id"`
	Exp    int64    `json:"exp"`
	Roles  []string `json:"role"`
	jwt.RegisteredClaims
}
