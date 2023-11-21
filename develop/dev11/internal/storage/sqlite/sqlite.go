package sqlite

// TODO: переделать
import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY,
		date DATE NOT NULL UNIQUE,
		event_name TEXT NOT NULL,
		event_description TEXT NOT NULL 
		);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateEvent(date time.Time, event_name, event_description string) (int64, error) {
	const op = "storage.sqlite.CreateEvent"

	stmt, err := s.db.Prepare("INSERT INTO events(date, event_name, event_description) VALUES(?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement %w", op, err)
	}
	res, err := stmt.Exec(date, event_name, event_description)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last inserted id: %w", op, err)
	}
	return id, nil
}

func (s *Storage) DeleteEvent(date time.Time, event_name string) (int64, error)
func (s *Storage) UpdateEvent(date time.Time, event_name string) (int64, error)
func (s *Storage) GetForDay(date time.Time) (int64, error)
func (s *Storage) GetForWeek(date time.Time) (int64, error)
func (s *Storage) GetForMonth(date time.Time) (int64, error)

// func (s *Storage) GetURL(alias string) (string, error) {
// 	const op = "storage.sqlite.GetURL"

// 	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias = '?'")
// 	if err != nil {
// 		return "", fmt.Errorf("%s: prepare statement %w", op, err)
// 	}
// 	var resURL string
// 	err = stmt.QueryRow(alias).Scan((&resURL))
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return "", fmt.Errorf("storage.ErrURLNotFound")
// 	}
// 	if err != nil {
// 		return "", fmt.Errorf("%s: execute statement: %w", op, err)
// 	}

// 	return resURL, nil
// }
