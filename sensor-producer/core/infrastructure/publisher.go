package infrastructure

import mqtt "github.com/eclipse/paho.mqtt.golang"

type publisher struct {
	mqttClient mqtt.Client
}

type Publisher interface {
	Publish(topic string, payload []byte) error
	Disconnect()
}

func NewPublisher(mqttClient mqtt.Client) Publisher {
	return &publisher{
		mqttClient: mqttClient,
	}
}

func (p *publisher) Publish(topic string, payload []byte) error {
	token := p.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func (p *publisher) Disconnect() {
	p.mqttClient.Disconnect(250)
}
