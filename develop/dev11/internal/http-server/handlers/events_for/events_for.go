package events_for

import (
	resp "dev11/internal/lib/api/response"
	"dev11/internal/lib/logger/sl"
	"dev11/internal/strct"
	"encoding/json"

	// "fmt"
	"log/slog"
	"net/http"
	"time"
	// "github.com/go-playground/validator/v10"
)

type Response struct {
	resp.Response
	Result any `json:"result,omitempty"`
}

type EventGetter interface {
	GetForDay(date time.Time) ([]strct.Event, error)
	GetForWeek(date time.Time) ([]strct.Event, error)
	GetForMonth(date time.Time) ([]strct.Event, error)
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
			eve := new(strct.Event)
			que := r.URL.Query()
			result := make([][]strct.Event, 0, 1)
			for _, date := range que["date"] {
				var err error
				eve.Date, err = time.Parse(time.DateOnly, date)
				if err != nil {
					log.Error(`error parsing date, please use "2000-10-10" formatting`)
					http.Error(w, `error parsing date, please use "2000-10-10" formatting`, http.StatusBadRequest)
					return
				}
				eves, err := eg.GetForDay(eve.Date)
				if err != nil {
					log.Error("failed to create event", sl.Err(err))
					http.Error(w, "failed to create event", http.StatusInternalServerError)
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
			eve := new(strct.Event)
			que := r.URL.Query()
			result := make([][]strct.Event, 0, 1)
			for _, date := range que["date"] {
				var err error
				eve.Date, err = time.Parse(time.DateOnly, date)
				if err != nil {
					log.Error(`error parsing date, please use "2000-10-10" formatting`)
					http.Error(w, `error parsing date, please use "2000-10-10" formatting`, http.StatusBadRequest)
					return
				}
				eves, err := eg.GetForWeek(eve.Date)
				if err != nil {
					log.Error("failed to create event", sl.Err(err))
					http.Error(w, "failed to create event", http.StatusInternalServerError)
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
			eve := new(strct.Event)
			que := r.URL.Query()
			result := make([][]strct.Event, 0, 1)
			for _, date := range que["date"] {
				var err error
				eve.Date, err = time.Parse(time.DateOnly, date)
				if err != nil {
					log.Error(`error parsing date, please use "2000-10-10" formatting`)
					http.Error(w, `error parsing date, please use "2000-10-10" formatting`, http.StatusBadRequest)
					return
				}
				eves, err := eg.GetForMonth(eve.Date)
				if err != nil {
					log.Error("failed to create event", sl.Err(err))
					http.Error(w, "failed to create event", http.StatusInternalServerError)
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
