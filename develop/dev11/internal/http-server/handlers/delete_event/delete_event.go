package delete_event

import (
	resp "dev11/internal/lib/api/response"
	"dev11/internal/lib/logger/sl"
	"dev11/internal/strct"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	resp.Response
	Result string `json:"result,omitempty"`
}

type EventDeleter interface {
	DeleteEvent(date time.Time, event_name string) (int64, error)
}

func New(log *slog.Logger, ed EventDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.delete_event.New"
		log = log.With(slog.String("op", op))

		if r.Method != http.MethodPost {
			log.Error("invalid http method, need POST")
			http.Error(w, "invalid http method, need POST", http.StatusBadRequest)
			return
		} else {
			eve := new(strct.Event)
			if err := json.NewDecoder(r.Body).Decode(eve); err != nil {
				log.Error("failed to decode request body", sl.Err(err))
				http.Error(w, "failed to decode request", http.StatusBadRequest)
				return
			}

			if err := validator.New().Struct(eve); err != nil {
				validateErr := err.(validator.ValidationErrors)
				log.Error("invalid request", sl.Err(err))
				http.Error(w, fmt.Sprintf("request validation error: %s", validateErr), http.StatusBadRequest)
				return
			}

			_, err := ed.DeleteEvent(eve.Date, eve.Name)
			if err != nil {
				log.Error("failed to delete event", sl.Err(err))
				http.Error(w, "failed to delete event", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(Response{Response: resp.OK(), Result: "success"}); err != nil {
				log.Error("failed to encode response", sl.Err(err))
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
				return
			}
		}
	}
}
