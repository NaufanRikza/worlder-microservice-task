package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorize(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			allowedSet := make(map[string]struct{}, len(roles))
			for _, r := range roles {
				allowedSet[r] = struct{}{}
			}

			claims := GetClaims(c)
			if claims == nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
			}

			for _, role := range claims.Roles {
				if _, ok := allowedSet[role]; ok {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{"message": "forbidden"})
		}
	}
}
