package mqttpublisher

import (
	"os"
	"solis-publisher/internal/util"
)

type mqttConfig struct {
	Username string
	Password string
	Server   string
	Port     string
	Topic    string
}

func getMqttConfig() (*mqttConfig, error) {
	username, err := util.LookupEnv("MQTT_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := util.LookupEnv("MQTT_PASSWORD")
	if err != nil {
		return nil, err
	}

	server, err := util.LookupEnv("MQTT_SERVER")
	if err != nil {
		return nil, err
	}

	config := mqttConfig{
		Username: *username,
		Password: *password,
		Server:   *server,
		Port:     "8883",
		Topic:    "solis/yieldkw",
	}

	port, ok := os.LookupEnv("MQTT_PORT")
	if ok {
		config.Port = port
	}

	topic, ok := os.LookupEnv("MQTT_TOPIC")
	if ok {
		config.Topic = topic
	}

	return &config, nil
}
