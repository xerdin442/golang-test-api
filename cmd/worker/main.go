package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hibiken/asynq"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xerdin442/api-practice/internal/env"
	"github.com/xerdin442/api-practice/internal/tasks"
)

func main() {
	// Improve readability of the logs in development
	if env.GetStr("NODE_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	redisConn := asynq.RedisClientOpt{
		Addr:     env.GetStr("REDIS_ADDR"),
		Password: env.GetStr("REDIS_PASSWORD"),
	}

	// Initialize the worker server
	srv := asynq.NewServer(
		redisConn,
		asynq.Config{Concurrency: 10},
	)

	// Define tasks handlers
	mux := asynq.NewServeMux()
	mux.HandleFunc("email_queue", tasks.HandleEmailTask)

	// Initialize the scheduler
	scheduler := asynq.NewScheduler(redisConn, nil)

	// Register cron jobs
	scheduler.Register("0 * * * *", asynq.NewTask("type:cleanup", nil))

	// Start the worker server
	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatal().Err(err).Msg("Worker initialization failed")
		}
	}()

	// Start the scheduler
	go func() {
		if err := scheduler.Run(); err != nil {
			log.Fatal().Err(err).Msg("Scheduler initialization failed")
		}
	}()

	log.Info().Msg("Worker and Scheduler are running...")

	// Keep the server running unless interrupted
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Graceful shutdown
	scheduler.Shutdown()
	srv.Shutdown()
	log.Warn().Msg("Shutdown complete.")
}
