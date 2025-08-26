package router

import (
	"sensor-producer/core/handler/http"

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
	e.POST("/sensor/:frequency", r.sensorHandler.ChangeFrequency)
}
