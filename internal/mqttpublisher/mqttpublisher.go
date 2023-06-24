package mqttpublisher

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"time"
)

func PublishCurrentPower(currentPower float32) error {
	log.Debug().Float32("currentPower", currentPower).Msg("Publishing current power")

	config, err := getMqttConfig()
	if err != nil {
		return err
	}

	opts := createClientOptions(config)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return fmt.Errorf("failed to connect to MQTT server- %s", token.Error())
	}

	payload := fmt.Sprintf("%f", currentPower)
	token := client.Publish(
		config.Topic,
		1,
		false,
		payload,
	)

	ok := token.WaitTimeout(10 * time.Second)
	if !ok {
		return fmt.Errorf("failed to publish current power (%f) - %s", currentPower, token.Error())
	}

	defer disconnectClient(client)

	return nil
}

func createClientOptions(config *mqttConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%s", config.Server, config.Port))
	opts.SetClientID("SolisCurrentPowerPublisher")
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)

	return opts
}

func disconnectClient(client mqtt.Client) {
	client.Disconnect(0)
	log.Info().Msg("Disconnected MQTT client")
}
