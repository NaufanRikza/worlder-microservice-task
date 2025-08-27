package http

import (
	"net/http"
	"sensor-consumer/core/dto"
	"sensor-consumer/core/usecase"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authUsecase usecase.AuthUsecase
	userUsecase usecase.UserUsecase
}

type AuthHandler interface {
	Login(c echo.Context) error
}

func NewAuthHandler(authUsecase usecase.AuthUsecase, userUsecase usecase.UserUsecase) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
		userUsecase: userUsecase,
	}
}

func (h *authHandler) Login(c echo.Context) error {
	req := dto.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	ctx := c.Request().Context()

	user, isValid, err := h.userUsecase.ValidateUser(ctx, req)
	if err != nil || !isValid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "username or password is incorrect"})
	}

	token, err := h.authUsecase.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
