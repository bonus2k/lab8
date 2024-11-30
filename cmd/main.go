package main

import (
	"context"
	"github.com/bonus2k/lab8/internal/models"
	"github.com/bonus2k/lab8/internal/repository"
	"log/slog"
	"strconv"
	"sync"
)

var log *slog.Logger

func main() {
	log = slog.Default()
	ctx := context.Background()

	repo, err := repository.Init("test.db")
	if err != nil {
		log.
			With(slog.AnyValue(err)).
			Error("repository.Init")
	}

	log.Info("Application start")
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(int2 int) {
			user := models.User{
				Name:     "test" + strconv.Itoa(int2),
				Email:    "<EMAIL>",
				Password: "<PASSWORD>",
			}
			err = repo.CreateUser(ctx, &user)
			if err != nil {
				log.
					With(slog.AnyValue(err)).
					Error("repo.CreateUser")
			}
			log.
				With(slog.Any("user", user)).
				Info("user created")
			wg.Done()
		}(i)
	}
	wg.Wait()
}
