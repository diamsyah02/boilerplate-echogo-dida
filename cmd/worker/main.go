package main

import (
	emailTask "boilerplate-echogo-dida/pkg/tasks/email"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Gagal memuat .env")
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDRESS")},
		asynq.Config{
			Concurrency: 10,
		},
	)

	mux := asynq.NewServeMux()
	// Tambahkan handler task di sini
	mux.HandleFunc(emailTask.TypeEmailWelcome, emailTask.HandleEmailTask) // email task

	log.Info().Msg("👷 Worker berjalan dan mendengarkan task...")
	if err := srv.Run(mux); err != nil {
		log.Fatal().Err(err).Msg("Gagal menjalankan worker")
	}
}
