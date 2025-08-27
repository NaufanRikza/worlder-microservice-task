package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sensor-producer/cmd"
	"sensor-producer/config"
	"sensor-producer/core/handler/http"
	"sensor-producer/core/infrastructure"
	"sensor-producer/core/router"
	"sensor-producer/core/usecase"
	"syscall"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic("Error loading config:")
	}

	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	freqChannel := make(chan uint, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down...")
		cancel()
	}()

	//start mqtt client
	mqttClient := cmd.StartMQTTClient(config.MqttConfig)
	publisher := infrastructure.NewPublisher(mqttClient)
	sensorUsecase := usecase.NewSensorUsecase(
		publisher,
		config.MqttConfig.Topic,
		config.AppConfig.DataGenerationFrequency,
		freqChannel,
	)

	go sensorUsecase.Start(ctx)

	e := echo.New()
	sensorHandler := http.NewSensorHandler(sensorUsecase)
	router := router.NewRouter(sensorHandler)
	group := e.Group("/api/v1")
	router.RegisterRoutes(group)

	// Start HTTP server
	cmd.StartHTTPServer(ctx, e)

	<-ctx.Done()
	fmt.Println("Cleanup complete.")
}
