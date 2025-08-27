package http

import (
	"net/http"
	"sensor-consumer/core/repository"
	"sensor-consumer/core/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type sensorHandler struct {
	UserUsecase usecase.SensorUsecase
}

type SensorHandler interface {
	GetSensorData(c echo.Context) error
	DeleteSensorData(c echo.Context) error
	UpdateSensorData(c echo.Context) error
}

func NewSensorHandler(userUsecase usecase.SensorUsecase) SensorHandler {
	return &sensorHandler{
		UserUsecase: userUsecase,
	}
}

func (h *sensorHandler) GetSensorData(c echo.Context) error {
	req := repository.SensorRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	data, err := h.UserUsecase.GetSensorData(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": data, "message": "success"})
}

func (h *sensorHandler) DeleteSensorData(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err := h.UserUsecase.DeleteSensorData(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func (h *sensorHandler) UpdateSensorData(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	body := repository.UpdateSensorBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err := h.UserUsecase.UpdateSensorData(id, body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
