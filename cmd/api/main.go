package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynq"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/env"
	repo "github.com/xerdin442/api-practice/internal/repository"
	"github.com/xerdin442/api-practice/internal/service"
)

type application struct {
	port       int
	services   *service.Manager
	cache      *cache.Redis
	tasksQueue *asynq.Client
}

func main() {
	// Initialize logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Improve readability of the logs in development
	if env.GetStr("NODE_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	// Validate connection string
	db, err := sql.Open("mysql", env.GetStr("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid database connection string")
	}

	// Connect to database
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Database connection failed")
	}
	defer db.Close()

	// Initialize cache, repositories and services
	cache := cache.NewRedis()
	registry := repo.NewRegistry(db)
	services := service.NewManager(registry)

	// Initialize task queue
	tasksQueue := asynq.NewClient(
		asynq.RedisClientOpt{Addr: env.GetStr("REDIS_ADDR"), Password: env.GetStr("REDIS_PASSWORD")},
	)

	app := &application{
		port:       env.GetInt("APP_PORT"),
		services:   services,
		cache:      cache,
		tasksQueue: tasksQueue,
	}

	// Start the http server
	if err := app.serve(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
