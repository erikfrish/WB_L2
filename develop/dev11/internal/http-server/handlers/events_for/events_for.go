package events_for

import (
	"log/slog"
	"net/http"
	"time"
)

type EventGetter interface {
	GetForDay(date time.Time) (int64, error)
	GetForWeek(date time.Time) (int64, error)
	GetForMonth(date time.Time) (int64, error)
}

func NewForDay(log *slog.Logger, eg EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func NewForWeek(log *slog.Logger, eg EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func NewForMonth(log *slog.Logger, eg EventGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
