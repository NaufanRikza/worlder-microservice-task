package http

import (
	"net/http"
	"sensor-producer/core/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type sensorHandler struct {
	SensorUsecase usecase.SensorUsecase
}

type SensorHandler interface {
	ChangeFrequency(c echo.Context) error
}

func NewSensorHandler(sensorUsecase usecase.SensorUsecase) SensorHandler {
	return &sensorHandler{
		SensorUsecase: sensorUsecase,
	}
}

func (h *sensorHandler) ChangeFrequency(c echo.Context) error {
	freq, err := strconv.ParseUint(c.Param("frequency"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid frequency parameter"})
	}
	err = h.SensorUsecase.ChangeFrequency(uint(freq))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to change frequency"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "frequency changed successfully"})
}
