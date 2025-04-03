package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/jszymanowski/alive/models"
	"github.com/jszymanowski/alive/repositories"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _, err := h.repo.FindAll(1, 100)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(users)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "ID")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		if render.Render(w, r, ErrInvalidRequest(err)) != nil {
			http.Error(w, "Failed to bind request", http.StatusBadRequest)
		}
		return
	}

	user, err := h.repo.FindByID(uint(idUint))
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(user)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

type UserPayload struct {
	*models.User
}

func (u *UserPayload) Bind(r *http.Request) error {
	return nil
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := &UserPayload{}

	if err := render.Bind(r, data); err != nil {
		if render.Render(w, r, ErrInvalidRequest(err)) != nil {
			http.Error(w, "Failed to bind request", http.StatusBadRequest)
		}
		return
	}

	user := data.User

	createdUser, err := h.repo.Create(user)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(createdUser)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
