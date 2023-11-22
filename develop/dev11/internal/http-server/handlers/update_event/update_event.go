package update_event

import (
	resp "dev11/internal/lib/api/response"
	"dev11/internal/lib/logger/sl"
	"dev11/internal/strct"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	resp.Response
	Result string `json:"result,omitempty"`
}

type EventUpdater interface {
	UpdateEvent(date, event_name, event_description string) (int64, error)
}

func New(log *slog.Logger, eu EventUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.update_event.New"
		log = log.With(slog.String("op", op))

		if r.Method != http.MethodPost {
			log.Error("invalid http method, need POST")
			http.Error(w, "invalid http method, need POST", http.StatusBadRequest)
			return
		} else {
			eve := new(strct.Event)
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			form := r.Form
			eve.Date = form.Get("date")
			eve.Name = form.Get("name")
			eve.Desc = form.Get("desc")
			if len(eve.Date) == 0 || len(eve.Name) == 0 {
				log.Error("you have to specify both name and date of event")
				http.Error(w, "you have to specify both name and date of event", http.StatusInternalServerError)
				return
			}

			validate := validator.New()
			err := validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
				field := fl.Field().String()
				match, _ := regexp.MatchString(`[0-9]{4}-[0-9]{2}-[0-9]{2}`, field)
				return match
			})
			if err != nil {
				log.Error("can't register validation func", sl.Err(err))
				http.Error(w, fmt.Sprintf("can't register validation func: %s", err), http.StatusInternalServerError)
				return
			}
			if err := validate.Struct(eve); err != nil {
				validateErr := err.(validator.ValidationErrors)
				log.Error("invalid request", sl.Err(err))
				http.Error(w, fmt.Sprintf("request validation error: %s", validateErr), http.StatusBadRequest)
				return
			}

			_, err = eu.UpdateEvent(eve.Date, eve.Name, eve.Desc)
			if err != nil {
				log.Error("failed to update event", sl.Err(err))
				http.Error(w, "failed to save event", http.StatusInternalServerError)
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
