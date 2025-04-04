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
	page := 1
	pageSize := 10

	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		if pageVal, err := strconv.Atoi(pageParam); err == nil && pageVal > 0 {
			page = pageVal
		}
	}

	if sizeParam := r.URL.Query().Get("size"); sizeParam != "" {
		if sizeVal, err := strconv.Atoi(sizeParam); err == nil && sizeVal > 0 {
			pageSize = sizeVal
		}
	}

	users, total, err := h.repo.FindAll(page, pageSize)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data": users,
		"pagination": map[string]interface{}{
			"page":  page,
			"size":  pageSize,
			"total": total,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
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
