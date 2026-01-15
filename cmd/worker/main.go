package main

import (
	"github.com/hibiken/asynq"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/xerdin442/api-practice/internal/env"
	"github.com/xerdin442/api-practice/internal/tasks"
)

func main() {
	addr := env.GetStr("REDIS_ADDR")
	password := env.GetStr("REDIS_PASSWORD")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: addr, Password: password},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("email_queue", tasks.HandleEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal().Err(err).Msg("Worker initialization failed")
	}
}
