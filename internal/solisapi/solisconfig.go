package solisapi

import (
	"fmt"
	"os"
	"strings"
)

type SolisConfig struct {
	SolisApiId     string
	SolisApiSecret string
	SolisApiUrl    string
}

func GetSolisApiConfig() (*SolisConfig, error) {
	apiId, err := lookupEnv("SOLISAPI_ID")
	if err != nil {
		return nil, err
	}

	apiSecret, err := lookupEnv("SOLISAPI_SECRET")
	if err != nil {
		return nil, err
	}

	url, err := lookupEnv("SOLISAPI_URL")
	if err != nil {
		return nil, err
	}

	config := SolisConfig{
		SolisApiId:     *apiId,
		SolisApiSecret: *apiSecret,
		SolisApiUrl:    *url,
	}

	return &config, nil
}

func IsValid(config *SolisConfig) bool {
	if config == nil {
		return false
	}

	if strings.TrimSpace(config.SolisApiId) == "" {
		return false
	}

	if strings.TrimSpace(config.SolisApiSecret) == "" {
		return false
	}

	if strings.TrimSpace(config.SolisApiUrl) == "" {
		return false
	}

	return true
}

func lookupEnv(key string) (*string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return nil, fmt.Errorf("no value for %s", key)
	}

	return &val, nil
}
