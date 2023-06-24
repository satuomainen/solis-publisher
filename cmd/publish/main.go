package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"solis-publisher/internal/mqttpublisher"
	"strconv"
)

// main function for publish publishes the value given as command line
// parameter to the configured MQTT topic
func main() {
	power, err := getFirstArg()
	if err != nil {
		log.Fatal().Err(err).Msg("Power value must be given as command line argument")
		os.Exit(1)
	}

	err = mqttpublisher.PublishCurrentPower(power)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to publish current power")
		os.Exit(1)
	} else {
		log.Info().Msg("Successfully published current power")
	}
}

func getFirstArg() (float32, error) {
	if len(os.Args) < 2 {
		log.Fatal().Msg("Value must be provided as command line argument")
	}

	power, err := strconv.ParseFloat(os.Args[1], 32)

	return float32(power), err
}
