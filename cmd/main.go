package main

import (
	"github.com/bonus2k/lab8/internal/handlers"
	"github.com/bonus2k/lab8/internal/repositories"
	"github.com/bonus2k/lab8/internal/services"
	"log/slog"
	"net/http"
	"os"
)

var log *slog.Logger

func main() {
	log = slog.Default()

	userRepo, err := repositories.Init("test.db")
	if err != nil {
		log.
			With(slog.AnyValue(err)).
			Error("repository.Init")
		os.Exit(1)
	}

	userService := services.Init(&userRepo)
	userHandler := handlers.Init(&userService)
	userRouter := handlers.UserRouter(&userHandler)

	log.Info("Application starting")
	err = http.ListenAndServe(":8080", userRouter)
	if err != nil {
		log.
			With(slog.AnyValue(err)).
			Error("Application start")
	}
}
