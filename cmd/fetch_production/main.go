package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"solis-publisher/internal/solisapi"
)

// main function for fetch_production connects to the Solis API and fetches the
// current power production in kWh
func main() {
	production, err := solisapi.FetchProduction()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch current power yield from Solis API")
		os.Exit(1)
	}

	fmt.Printf("%f", *production)
}
