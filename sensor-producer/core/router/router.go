package router

import (
	"sensor-producer/core/handler/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type router struct {
	sensorHandler http.SensorHandler
}

type Router interface {
	RegisterRoutes(e *echo.Group)
}

func NewRouter(sensorHandler http.SensorHandler) Router {
	return &router{
		sensorHandler: sensorHandler,
	}
}

func (r *router) RegisterRoutes(e *echo.Group) {
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte("your-secret-key"),
		TokenLookup: "header:Authorization",
	}))

	e.POST("/sensor/:frequency", r.sensorHandler.ChangeFrequency)
}
