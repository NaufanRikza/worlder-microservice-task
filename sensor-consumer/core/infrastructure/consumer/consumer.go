package consumer

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type consumer struct {
	mqttClient mqtt.Client
	topic      string
}

type Consumer interface {
	Consume(handler func(client mqtt.Client, msg mqtt.Message)) error
	Close()
}

func NewConsumer(mqttClient mqtt.Client, topic string) Consumer {
	return &consumer{
		mqttClient: mqttClient,
		topic:      topic,
	}
}

func (c *consumer) Consume(handler func(client mqtt.Client, msg mqtt.Message)) error {
	token := c.mqttClient.Subscribe(c.topic, 1, handler)
	token.Wait()
	return token.Error()
}

func (c *consumer) Close() {
	c.mqttClient.Unsubscribe(c.topic)
	c.mqttClient.Disconnect(250)
}
