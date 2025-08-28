package consumer

import mqtt "github.com/eclipse/paho.mqtt.golang"

type consumer struct {
	mqttClient mqtt.Client
	topic      string
}

type Consumer interface {
	Consume(handler func(client mqtt.Client, msg mqtt.Message)) error
}

func NewConsumer(mqttClient mqtt.Client, topic string) Consumer {
	return &consumer{
		mqttClient: mqttClient,
		topic:      topic,
	}
}

func (c *consumer) Consume(handler func(client mqtt.Client, msg mqtt.Message)) error {
	token := c.mqttClient.Subscribe(c.topic, 0, handler)
	token.Wait()
	return token.Error()
}
