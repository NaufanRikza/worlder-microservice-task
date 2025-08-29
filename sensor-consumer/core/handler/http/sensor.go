package http

import (
	"net/http"
	"sensor-consumer/core/dto"
	"sensor-consumer/core/usecase"
	"strconv"
	"time"

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
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	if (req.TimeEnd != nil && !req.TimeEnd.IsZero()) && (req.TimeStart == nil || req.TimeStart.IsZero()) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	ctx := c.Request().Context()

	data, err := h.UserUsecase.GetSensorData(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error : failed to get sensor data"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": data, "message": "success"})
}

func (h *sensorHandler) DeleteSensorData(c echo.Context) error {
	req := dto.DeleteSensorRequest{}

	id := uint64(0)
	id, _ = strconv.ParseUint(c.Param("id"), 10, 64)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	req.ID = id
	if req == (dto.DeleteSensorRequest{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	if (req.TimeEnd != nil && !req.TimeEnd.IsZero()) && (req.TimeStart == nil || req.TimeStart.IsZero()) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	ctx := c.Request().Context()

	if err := h.UserUsecase.DeleteSensorData(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error : failed to delete sensor data"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func (h *sensorHandler) UpdateSensorData(c echo.Context) error {
	id := uint64(0)
	id, _ = strconv.ParseUint(c.Param("id"), 10, 64)

	//parse body
	body := dto.UpdateSensorBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	if err := c.Validate(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	//parse params
	req := dto.UpdateSensorRequest{}
	req.ID = id
	req.ID1 = c.QueryParam("id1")
	req.ID2, _ = strconv.ParseUint(c.QueryParam("id2"), 10, 64)
	timeStart, _ := time.Parse(time.RFC3339, c.QueryParam("time_start"))
	timeEnd, _ := time.Parse(time.RFC3339, c.QueryParam("time_end"))

	if !timeStart.IsZero() {
		req.TimeStart = &timeStart
	}

	if !timeEnd.IsZero() {
		req.TimeEnd = &timeEnd
	}

	if (req.TimeEnd != nil && !req.TimeEnd.IsZero()) && (req.TimeStart == nil || req.TimeStart.IsZero()) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	ctx := c.Request().Context()

	if err := h.UserUsecase.UpdateSensorData(ctx, req, body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error : failed to update sensor data"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
