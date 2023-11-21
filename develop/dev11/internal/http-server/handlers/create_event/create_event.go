package create_event

import (
	"log/slog"
	"net/http"
	"time"
)

type EventCreator interface {
	CreateEvent(date time.Time, event_name, event_description string) (int64, error)
}

func New(log *slog.Logger, ec EventCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
