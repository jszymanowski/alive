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

type MonitorHandler struct {
	repo *repositories.MonitorRepository
}

func NewMonitorHandler(repo *repositories.MonitorRepository) *MonitorHandler {
	return &MonitorHandler{repo: repo}
}

func (h *MonitorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	monitors, _, err := h.repo.FindAll(1, 100)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(monitors)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *MonitorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "ID")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		if render.Render(w, r, ErrInvalidRequest(err)) != nil {
			http.Error(w, "Failed to bind request", http.StatusBadRequest)
		}
		return
	}

	monitor, err := h.repo.FindByID(uint(idUint))
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(monitor)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

type MonitorPayload struct {
	*models.Monitor
}

func (u *MonitorPayload) Bind(r *http.Request) error {
	return nil
}

func (h *MonitorHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := &MonitorPayload{}

	if err := render.Bind(r, data); err != nil {
		if render.Render(w, r, ErrInvalidRequest(err)) != nil {
			http.Error(w, "Failed to bind request", http.StatusBadRequest)
		}
		return
	}

	monitor := data.Monitor

	createdMonitor, err := h.repo.Create(monitor)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(createdMonitor)
	if encodeErr != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
