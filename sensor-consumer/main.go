package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sensor-consumer/cmd"
	"sensor-consumer/config"
	http_handler "sensor-consumer/core/handler/http"
	"sensor-consumer/core/infrastructure/auth"
	mqtt_consumer "sensor-consumer/core/infrastructure/consumer"
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

	config, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		panic(err)
	}

	db, err := cmd.NewDatabaseInstance(&config.DatabaseConfig)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		panic(err)
	}

	e := echo.New()

	//repository
	sensorRepository := repository.NewSensorRepository(db)
	userRepository := repository.NewUserRepository(db)

	jwtManager := auth.NewJWTManager("")
	passwordHasher := auth.NewPasswordHasher()
	mqttClient := cmd.StartMQTTClient(config.MqttConfig)
	consumer := mqtt_consumer.NewConsumer(mqttClient, config.MqttConfig.Topic)

	//usecase
	authUsecase := usecase.NewAuthUsecase(
		jwtManager,
		passwordHasher,
	)
	sensorUsecase := usecase.NewSensorUsecase(
		sensorRepository,
		consumer,
	)
	userUseCase := usecase.NewUserUsecase(userRepository)

	//handler
	authHandler := http_handler.NewAuthHandler(
		authUsecase,
		userUseCase,
	)
	sensorHandler := http_handler.NewSensorHandler(sensorUsecase)

	//router
	authRouter := router.NewAuthRouter(authHandler)
	sensorRouter := router.NewSensorRouter(sensorHandler)

	authRouter.RegisterRoutes(e)

	groupV1 := e.Group("/api/v1")
	sensorRouter.RegisterRoutes(groupV1, config.JWTConfig.SecretKey)

	cmd.StartMQTTSubscriber(sensorUsecase)
	cmd.StartHTTPServer(ctx, e)

	<-ctx.Done()
	fmt.Println("Cleanup complete.")
}
