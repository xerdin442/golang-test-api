package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/config"
	repo "github.com/xerdin442/api-practice/internal/repository"
	"github.com/xerdin442/api-practice/internal/service"
)

type application struct {
	port       int
	services   *service.Manager
	cache      *cache.Cache
	tasksQueue *asynq.Client
	cfg        *config.Config
}

func main() {
	// Initialize logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Load environment variables
	cfg := config.Load()

	// Improve readability of the logs in development
	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	// Validate connection string
	db, err := sql.Open("mysql", cfg.GooseDbString)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid database connection string")
	}

	// Connect to database
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Database connection failed")
	}
	defer db.Close()

	// Initialize cache, repositories and services
	cache := cache.New(cfg)
	registry := repo.NewRegistry(db)
	services := service.NewManager(registry, cfg)

	// Initialize task queue
	tasksQueue := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     cfg.RedisAddr,
			Password: cfg.RedisPassword,
		},
	)

	app := &application{
		port:       cfg.AppPort,
		services:   services,
		cache:      cache,
		tasksQueue: tasksQueue,
		cfg:        cfg,
	}

	// Start the http server
	if err := app.serve(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
