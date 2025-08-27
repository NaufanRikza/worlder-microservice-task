package http

import (
	"net/http"
	"sensor-consumer/core/dto"
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
	req := dto.SensorRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	ctx := c.Request().Context()

	data, err := h.UserUsecase.GetSensorData(ctx, req)
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

	ctx := c.Request().Context()

	if err := h.UserUsecase.DeleteSensorData(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func (h *sensorHandler) UpdateSensorData(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	body := dto.UpdateSensorBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	ctx := c.Request().Context()

	if err := h.UserUsecase.UpdateSensorData(ctx, id, body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
