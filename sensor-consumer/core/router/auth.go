package router

import (
	"sensor-consumer/core/handler/http"
	"github.com/labstack/echo/v4"
)

type authRouter struct {
	authHandler http.AuthHandler
}

type AuthRouter interface {
	RegisterRoutes(e *echo.Echo)
}

func NewAuthRouter(authHandler http.AuthHandler) AuthRouter {
	return &authRouter{
		authHandler: authHandler,
	}
}

func (r *authRouter) RegisterRoutes(e *echo.Echo) {
	e.POST("/login", r.authHandler.Login)
}


