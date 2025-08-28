package router

import (
	"sensor-consumer/core/handler/http"
	"sensor-consumer/core/infrastructure/middleware"

	"github.com/labstack/echo/v4"
)

type sensorRouter struct {
	sensorHandler http.SensorHandler
}

type SensorRouter interface {
	RegisterRoutes(g *echo.Group, jwtSecretKey string)
}

func NewSensorRouter(sensorHandler http.SensorHandler) SensorRouter {
	return &sensorRouter{
		sensorHandler: sensorHandler,
	}
}

func (r *sensorRouter) RegisterRoutes(g *echo.Group, jwtSecretKey string) {
	g.Use(middleware.JWTMiddleware(jwtSecretKey))

	g.GET("/sensor", r.sensorHandler.GetSensorData)
	g.DELETE("/sensor/:id", r.sensorHandler.DeleteSensorData)
	g.PATCH("/sensor/:id", r.sensorHandler.UpdateSensorData)
}
