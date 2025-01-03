package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vlamat/currency-service/internal/repository"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAllRatesHandler(w http.ResponseWriter, r *http.Request) {
	rates, err := repository.GetAllRates(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetRatesByDateHandler(w http.ResponseWriter, r *http.Request) {
	// Предположим, мы получаем дату через query param ?date=2025-01-03
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "date parameter is required", http.StatusBadRequest)
		return
	}

	dateParsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	rates, err := repository.GetRatesByDate(h.DB, dateParsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
