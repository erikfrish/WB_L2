package events_for

import (
	resp "dev11/internal/lib/api/response"
	"dev11/internal/lib/logger/sl"
	"dev11/internal/strct"
	"encoding/json"

	// "fmt"
	"log/slog"
	"net/http"
	// "github.com/go-playground/validator/v10"
)

type Response struct {
	resp.Response
	Result any `json:"result,omitempty"`
}

type EventGetter interface {
	GetForDay(date string) ([]strct.Event, error)
	GetForWeek(date string) ([]strct.Event, error)
	GetForMonth(date string) ([]strct.Event, error)
}

func NewForDay(log *slog.Logger, eg EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.events_for.NewForDay"
		log = log.With(slog.String("op", op))

		if r.Method != http.MethodGet {
			log.Error("invalid http method, need GET")
			http.Error(w, "invalid http method, need GET", http.StatusBadRequest)
			return
		} else {
			que := r.URL.Query()
			result := make([][]strct.Event, 0, 1)
			for _, date := range que["date"] {
				eves, err := eg.GetForDay(date)
				if err != nil {
					log.Error("failed to get events for day", sl.Err(err))
					http.Error(w, "failed to get events for day", http.StatusInternalServerError)
					return
				}
				result = append(result, eves)
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(Response{Response: resp.OK(), Result: result}); err != nil {
				log.Error("failed to encode response", sl.Err(err))
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
				return
			}

		}

	}
}
func NewForWeek(log *slog.Logger, eg EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.events_for.NewForWeek"
		log = log.With(slog.String("op", op))

		if r.Method != http.MethodGet {
			log.Error("invalid http method, need GET")
			http.Error(w, "invalid http method, need GET", http.StatusBadRequest)
			return
		} else {
			que := r.URL.Query()
			result := make([][]strct.Event, 0, 1)
			for _, date := range que["date"] {
				eves, err := eg.GetForWeek(date)
				if err != nil {
					log.Error("failed to get events for week", sl.Err(err))
					http.Error(w, "failed to get events for week", http.StatusInternalServerError)
					return
				}
				result = append(result, eves)
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(Response{Response: resp.OK(), Result: result}); err != nil {
				log.Error("failed to encode response", sl.Err(err))
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
				return
			}
		}

	}
}
func NewForMonth(log *slog.Logger, eg EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.events_for.NewForMonth"
		log = log.With(slog.String("op", op))

		if r.Method != http.MethodGet {
			log.Error("invalid http method, need GET")
			http.Error(w, "invalid http method, need GET", http.StatusBadRequest)
			return
		} else {
			que := r.URL.Query()
			result := make([][]strct.Event, 0, 1)
			for _, date := range que["date"] {
				eves, err := eg.GetForMonth(date)
				if err != nil {
					log.Error("failed to get events for month", sl.Err(err))
					http.Error(w, "failed to get events for month", http.StatusInternalServerError)
					return
				}
				result = append(result, eves)
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(Response{Response: resp.OK(), Result: result}); err != nil {
				log.Error("failed to encode response", sl.Err(err))
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
				return
			}
		}

	}
}
