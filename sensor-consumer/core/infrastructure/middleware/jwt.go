package middleware

import (
	"net/http"
	"sensor-consumer/core/infrastructure/auth"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	ClaimsKey = "claims"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing Authorization header"})
			}

			// Check "Bearer" prefix
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid Authorization header format"})
			}
			tokenString := parts[1]

			claims := &auth.JWTClaim{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
			}

			if claims.Exp < time.Now().Unix() {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "token expired"})
			}

			c.Set(ClaimsKey, claims)


			return next(c)
		}
	}
}

func GetClaims(c echo.Context) *auth.JWTClaim {
	v := c.Get(ClaimsKey)
	if v == nil {
		return nil
	}
	if mc, ok := v.(*auth.JWTClaim); ok {
		return mc
	}
	return nil
}
