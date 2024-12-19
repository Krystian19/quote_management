package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	REDIS_TEST_URL string

	DB_URL      string
	DB_TEST_URL string

	EXTERNAL_BFF_PORT string
	INTERNAL_BFF_PORT int
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		if value != "" {
			return value
		}
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

func GetEnv() Env {
	// ignore file if we are currently in the CI container
	if !getBoolEnv("IS_CI", false) {
		godotenv.Load("../../../.env")
	}

	return Env{
		REDIS_TEST_URL: getEnv("REDIS_TEST_URL", "quote.redis:6379"),

		DB_URL:      getEnv("DB_URL", "host=quote.postgres port=5432 user=postgres password=pass dbname=quote sslmode=disable"),
		DB_TEST_URL: getEnv("DB_TEST_URL", "host=quote.postgres port=5432 user=postgres password=pass dbname=quote_test sslmode=disable"),

		EXTERNAL_BFF_PORT: getEnv("EXTERNAL_BFF_PORT", "4000"),
		INTERNAL_BFF_PORT: getIntEnv("INTERNAL_BFF_PORT", 9040),
	}
}
