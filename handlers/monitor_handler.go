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

	monitors, total, err := h.repo.FindAll(page, pageSize)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data": monitors,
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

func (h *MonitorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
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
