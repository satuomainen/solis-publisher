package util

import (
	"fmt"
	"os"
)

func LookupEnv(key string) (*string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return nil, fmt.Errorf("no value in environment for %s", key)
	}

	return &val, nil
}
