package app

import (
	"os"
)

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	return fallback
}