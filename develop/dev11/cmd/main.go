package main

/*
=== HTTP server ===
Реализовать HTTP-сервер для работы с календарем.
В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:

1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
2. Реализовать вспомогательные функции для парсинга
и валидации параметров методов /create_event и /update_event.
3. Реализовать HTTP обработчики для каждого из методов API,
используя вспомогательные функции и объекты доменной области.
4. Реализовать middleware для логирования запросов

Методы API:

- POST /create_event
- POST /update_event
- POST /delete_event
- GET /events_for_day
- GET /events_for_week
- GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.

В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."}
в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:

1. Реализовать все методы.
2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту
указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

import (
	"dev11/internal/config"
	"dev11/internal/http-server/handlers/create_event"
	"dev11/internal/http-server/handlers/delete_event"
	"dev11/internal/http-server/handlers/events_for"
	"dev11/internal/http-server/handlers/update_event"
	mwLogger "dev11/internal/http-server/middleware/logger"
	"dev11/internal/lib/logger/handlers/slogpretty"
	"dev11/internal/lib/logger/sl"
	"dev11/internal/storage/sqlite"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting dev11 on ", slog.String("address", cfg.Address))
	log.Info("", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage

	router := http.NewServeMux()

	mwLogger := mwLogger.New(log)
	router.HandleFunc("/create_event", mwLogger(create_event.New(log, storage))) // хэндлеры
	router.HandleFunc("/update_event", mwLogger(update_event.New(log, storage)))
	router.HandleFunc("/delete_event", mwLogger(delete_event.New(log, storage)))
	router.HandleFunc("/events_for_day", mwLogger(events_for.NewForDay(log, storage)))
	router.HandleFunc("/events_for_week", mwLogger(events_for.NewForWeek(log, storage)))
	router.HandleFunc("/events_for_month", mwLogger(events_for.NewForMonth(log, storage)))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case "local":
		log = setupPrettySlog()
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
