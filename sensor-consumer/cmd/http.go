package cmd

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartHTTPServer(ctx context.Context, e *echo.Echo) error {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			panic("Failed to start HTTP server: " + err.Error())
		}
	}()

	<-ctx.Done()
	return e.Shutdown(ctx)
}
