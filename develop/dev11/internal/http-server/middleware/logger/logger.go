package logger

import (
	"net/http"
	"time"

	"log/slog"
)

func New(log *slog.Logger) func(next http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		log = log.With(
			slog.String("component", "middleware/logger"),
		)

		log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
			)
			ww := NewLogResponseWriter(w)

			t1 := time.Now()
			defer func() {
				entry.Info("request completed",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.Size()),
					slog.String("duration", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

// LogResponseWriter wraps the standard http.ResponseWriter allowing for more
// verbose logging
type LogResponseWriter struct {
	status int
	size   int
	http.ResponseWriter
}

func NewLogResponseWriter(res http.ResponseWriter) *LogResponseWriter {
	// Default the status code to 200
	return &LogResponseWriter{status: 200, ResponseWriter: res}
}

// Status provides an easy way to retrieve the status code
func (w *LogResponseWriter) Status() int {
	return w.status
}

// Size provides an easy way to retrieve the response size in bytes
func (w *LogResponseWriter) Size() int {
	return w.size
}

// Header returns & satisfies the http.ResponseWriter interface
func (w *LogResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write satisfies the http.ResponseWriter interface and
// captures data written, in bytes
func (w *LogResponseWriter) Write(data []byte) (int, error) {

	written, err := w.ResponseWriter.Write(data)
	w.size += written

	return written, err
}

// WriteHeader satisfies the http.ResponseWriter interface and
// allows us to catch the status code
func (w *LogResponseWriter) WriteHeader(statusCode int) {

	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
