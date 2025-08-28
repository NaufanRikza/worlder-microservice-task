package router

import (
	"sensor-producer/core/handler/http"
	"sensor-producer/core/infrastructure/middleware"

	"github.com/labstack/echo/v4"
)

type router struct {
	sensorHandler http.SensorHandler
}

type Router interface {
	RegisterRoutes(e *echo.Group, jwtSecret string)
}

func NewRouter(sensorHandler http.SensorHandler) Router {
	return &router{
		sensorHandler: sensorHandler,
	}
}

func (r *router) RegisterRoutes(e *echo.Group, jwtSecret string) {
	e.Use(middleware.JWTMiddleware(jwtSecret))

	e.POST("/sensor/:frequency", r.sensorHandler.ChangeFrequency, middleware.Authorize("admin"))
}
