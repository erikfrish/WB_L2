package delete_event

import (
	"log/slog"
	"net/http"
	"time"
)

type EventDeleter interface {
	DeleteEvent(date time.Time, event_name string) (int64, error)
}

func New(log *slog.Logger, ed EventDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
