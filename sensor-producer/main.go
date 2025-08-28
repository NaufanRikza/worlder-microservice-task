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

	freqChannel := make(chan uint, 1)
	ctx, cancel := context.WithCancel(context.Background())

	//start mqtt client
	mqttClient := cmd.StartMQTTClient(config.MqttConfig)
	publisher := infrastructure.NewPublisher(mqttClient)
	sensorUsecase := usecase.NewSensorUsecase(
		publisher,
		config.MqttConfig.Topic,
		config.AppConfig.DataGenerationFrequency,
		freqChannel,
	)

	go sensorUsecase.Start(ctx, config.AppConfig.SensorType, config.AppConfig.SensorID)

	e := echo.New()
	sensorHandler := http.NewSensorHandler(sensorUsecase)
	router := router.NewRouter(sensorHandler)
	group := e.Group("/api/v1")
	router.RegisterRoutes(group, config.JWTConfig.SecretKey)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down...")
		cancel()
	}()
	// Start HTTP server
	cmd.StartHTTPServer(ctx, e, config.AppConfig.ProducerHTTPPort)

	<-ctx.Done()
	fmt.Println("Cleanup complete.")
}
