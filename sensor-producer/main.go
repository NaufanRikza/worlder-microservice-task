package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sensor-producer/cmd"
	"sensor-producer/config"
	"syscall"
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

	go cmd.StartHTTPServer()
	go cmd.StartMQTTClient(config.MqttConfig, freqChannel, ctx)

	<-ctx.Done()
	fmt.Println("Cleanup complete.")
}
