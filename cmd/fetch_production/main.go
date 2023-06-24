package main

import (
	"fmt"
	"os"
	"solis-publisher/internal/solisapi"
)

// main function for fetch_production connects to the Solis API and fetches the
// current power production in kWh
func main() {
	config, err := solisapi.GetSolisApiConfig()
	if err != nil {
		fmt.Println("solis-publisher/fetch_production: Configuration error", err)
		os.Exit(1)
	}

	production, err := solisapi.FetchProduction(config)
	if err != nil {
		fmt.Println("solis-publisher/fetch_production: API error", err)
	}

	fmt.Printf("%s", *production)
}
