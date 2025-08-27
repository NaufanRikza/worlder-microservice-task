package cmd

import (
	"context"

	"github.com/labstack/echo/v4"
)

func StartHTTPServer(ctx context.Context, e *echo.Echo) error {
	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			panic("Failed to start HTTP server: " + err.Error())
		}
	}()

	<-ctx.Done()
	return e.Shutdown(ctx)
}
