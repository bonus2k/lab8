package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

var (
	h UserHandler
)

func UserRouter(handler *UserHandler) *chi.Mux {
	h = *handler
	return route()
}

func route() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(3 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User service"))
	})

	// RESTy routes for "User" resource
	r.Route("/api", func(r chi.Router) {

		r.Post("/user", h.CreateUser)
		r.Get("/user/{userID}", h.GetUser)
		r.Get("/user", h.GetAllUsers)

	})
	return r
}
