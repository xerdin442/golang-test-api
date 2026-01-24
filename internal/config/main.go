package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

type Config struct {
	JwtSecret           string
	AppName             string
	AppPort             int
	Environment         string
	GooseDbString       string
	RedisAddr           string
	RedisPassword       string
	ResendEmailApiKey   string
	AppEmail            string
	CloudinaryName      string
	CloudinarySecret    string
	CloudinaryApiKey    string
	DefaultProfileImage string
}

func Load() *Config {
	return &Config{
		JwtSecret:           GetStr("JWT_SECRET"),
		AppName:             GetStr("APP_NAME"),
		AppEmail:            GetStr("APP_EMAIL"),
		AppPort:             GetInt("APP_PORT"),
		Environment:         GetStr("ENVIRONMENT"),
		GooseDbString:       GetStr("GOOSE_DBSTRING"),
		RedisAddr:           GetStr("REDIS_ADDR"),
		RedisPassword:       GetStr("REDIS_PASSWORD"),
		ResendEmailApiKey:   GetStr("RESEND_EMAIL_API_KEY"),
		CloudinaryName:      GetStr("CLOUDINARY_NAME"),
		CloudinarySecret:    GetStr("CLOUDINARY_SECRET"),
		CloudinaryApiKey:    GetStr("CLOUDINARY_API_KEY"),
		DefaultProfileImage: GetStr("DEFAULT_PROFILE_IMAGE"),
	}
}

func GetStr(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatal().Msgf("Missing environment variable: %s", key)
	}

	return value
}

func GetInt(key string) int {
	strValue := GetStr(key)

	intValue, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid string value")
	}

	return int(intValue)
}
