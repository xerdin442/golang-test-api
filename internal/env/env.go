package env

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

func GetStr(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatal().Msgf("Missing environment variable: %s", key)
	}

	return value
}

func GetInt(key string) int {
	strValue := GetStr(key)

	intValue, err := strconv.ParseInt(strValue, 10, 0)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid string value")
	}

	return int(intValue)
}
