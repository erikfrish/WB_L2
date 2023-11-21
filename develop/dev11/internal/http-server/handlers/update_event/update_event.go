package update_event

import (
	"log/slog"
	"net/http"
	"time"
)

type EventUpdater interface {
	UpdateEvent(date time.Time, event_name string) (int64, error)
}

func New(log *slog.Logger, eu EventUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
