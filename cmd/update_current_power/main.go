package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"solis-publisher/internal/mqttpublisher"
	"solis-publisher/internal/solisapi"
)

// main function for update current power gets the current power yield from
// Solis API and publishes it to the configure MQTT topic
func main() {
	log.Info().Msg("Fetching and updating")

	power, err := solisapi.FetchProduction()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch current power from Solis API")
		os.Exit(1)
	}

	err = mqttpublisher.PublishCurrentPower(*power)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to publish current power")
		os.Exit(1)
	}

	log.Info().Float32("power", *power).Msg("Successfully published current power")
}
