package env

import (
	"log"
	"os"
	"strconv"
)

func GetStr(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Missing environment variable: %s", key)
	}

	return value
}

func GetInt(key string) int {
	strValue := GetStr(key)

	intValue, err := strconv.ParseInt(strValue, 10, 0)
	if err != nil {
		log.Fatal("Invalid string value", err)
	}

	return int(intValue)
}
