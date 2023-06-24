package solisapi

import (
	"solis-publisher/internal/util"
	"strings"
)

type SolisConfig struct {
	SolisApiId     string
	SolisApiSecret string
	SolisApiUrl    string
}

func GetSolisApiConfig() (*SolisConfig, error) {
	apiId, err := util.LookupEnv("SOLISAPI_ID")
	if err != nil {
		return nil, err
	}

	apiSecret, err := util.LookupEnv("SOLISAPI_SECRET")
	if err != nil {
		return nil, err
	}

	url, err := util.LookupEnv("SOLISAPI_URL")
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
