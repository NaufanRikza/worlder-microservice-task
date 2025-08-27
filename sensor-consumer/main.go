package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sensor-consumer/cmd"
	"sensor-consumer/config"
	http_handler "sensor-consumer/core/handler/http"
	"sensor-consumer/core/repository"
	"sensor-consumer/core/router"
	"sensor-consumer/core/usecase"
	"syscall"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down...")
		cancel()
	}()

	databaseConfig, err := config.GetDatabaseConfig()
	if err != nil {
		fmt.Println("Error loading database config:", err)
		panic(err)
	}

	db, err := cmd.NewDatabaseInstance(databaseConfig)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		panic(err)
	}

	e := echo.New()

	sensorRepository := repository.NewSensorRepository(db)
	// userRepository := repository.NewUserRepository(db)

	sensorUsecase := usecase.NewSensorUsecase(sensorRepository)

	sensorHandler := http_handler.NewSensorHandler(sensorUsecase)

	sensorRouter := router.NewSensorRouter(sensorHandler)

	groupV1 := e.Group("/api/v1")
	sensorRouter.RegisterRoutes(groupV1)

	cmd.StartHTTPServer(ctx, e)

	<-ctx.Done()
	fmt.Println("Cleanup complete.")
}
