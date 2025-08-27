package router

import (
	"sensor-consumer/core/handler/http"

	"github.com/labstack/echo/v4"
)

type sensorRouter struct {
	sensorHandler http.SensorHandler
}

type SensorRouter interface {
	RegisterRoutes(g *echo.Group)
}

func NewSensorRouter(sensorHandler http.SensorHandler) SensorRouter {
	return &sensorRouter{
		sensorHandler: sensorHandler,
	}
}

func (r *sensorRouter) RegisterRoutes(g *echo.Group) {
	g.GET("/sensor", r.sensorHandler.GetSensorData)
	g.DELETE("/sensor/:id", r.sensorHandler.DeleteSensorData)
	g.PATCH("/sensor/:id", r.sensorHandler.UpdateSensorData)
}
