package handlers

import (
	"encoding/json"
	"github.com/bonus2k/lab8/internal/models"
	"github.com/bonus2k/lab8/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

var (
	handlers UserHandler
)

type UserHandlerImpl struct {
	service services.UserService
}

func (u UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	if "application/json" != r.Header.Get("Content-Type") {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	var user models.UserReq
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf, &user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	res, err := u.service.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (u UserHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	res, err := u.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (u UserHandlerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.service.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func Init(service *services.UserService) UserHandler {
	handlers = UserHandlerImpl{service: *service}
	return handlers
}
