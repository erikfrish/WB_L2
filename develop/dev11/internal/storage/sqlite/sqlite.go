package sqlite

import (
	"database/sql"
	"dev11/internal/strct"
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
		date TEXT NOT NULL,
		name TEXT NOT NULL,
		description TEXT NOT NULL 
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

func (s *Storage) CreateEvent(event_date time.Time, event_name, event_description string) (int64, error) {
	const op = "storage.sqlite.CreateEvent"

	stmt, err := s.db.Prepare("INSERT INTO events(date, name, description) VALUES(?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement %w", op, err)
	}
	res, err := stmt.Exec(event_date.Format(time.DateOnly), event_name, event_description)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last inserted id: %w", op, err)
	}
	return id, nil
}

func (s *Storage) UpdateEvent(event_date time.Time, event_name, event_description string) (int64, error) {
	const op = "storage.sqlite.UpdateEvent"

	stmt, err := s.db.Prepare("UPDATE events SET description=? WHERE date=? AND name=?")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement %w", op, err)
	}
	res, err := stmt.Exec(event_description, event_date.Format(time.DateOnly), event_name)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get rows affected ra: %w", op, err)
	}
	return ra, nil
}
func (s *Storage) DeleteEvent(event_date time.Time, event_name string) (int64, error) {
	const op = "storage.sqlite.DeleteEvent"

	stmt, err := s.db.Prepare("DELETE FROM events WHERE WHERE date=? AND name=?")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement %w", op, err)
	}
	res, err := stmt.Exec(event_date.Format(time.DateOnly), event_name)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get rows affected ra: %w", op, err)
	}
	return ra, nil
}

func (s *Storage) GetForDay(event_date time.Time) ([]strct.Event, error) {
	const op = "storage.sqlite.GetForDay"
	var result = make([]strct.Event, 0)
	row, err := s.db.Query("SELECT * FROM events where date=?", event_date.Format(time.DateOnly))
	if err != nil {
		return []strct.Event{}, fmt.Errorf("%s: prepare statement %w", op, err)
	}
	for row.Next() { // Iterate and fetch the records from result cursor
		item := strct.Event{}
		err := row.Scan(&item.Date, &item.Name, &item.Description)
		if err != nil {
			return []strct.Event{}, fmt.Errorf("%s: scanning row %w", op, err)
		}
		result = append(result, item)
	}
	return result, nil
}
func (s *Storage) GetForWeek(event_date time.Time) ([]strct.Event, error) {
	const op = "storage.sqlite.GetForWeek"
	var result = make([]strct.Event, 0)
	for today := 0; today < 7; today++ {
		event_date = event_date.Add(time.Hour * 24)
		row, err := s.db.Query("SELECT * FROM events where date=?", event_date.Format(time.DateOnly))
		if err != nil {
			return []strct.Event{}, fmt.Errorf("%s: prepare statement %w", op, err)
		}
		for row.Next() {
			item := strct.Event{}
			err := row.Scan(&item.Date, &item.Name, &item.Description)
			if err != nil {
				return []strct.Event{}, fmt.Errorf("%s: scanning row %w", op, err)
			}
			result = append(result, item)
		}
	}
	return result, nil
}
func (s *Storage) GetForMonth(event_date time.Time) ([]strct.Event, error) {
	const op = "storage.sqlite.GetForMonth"
	var result = make([]strct.Event, 0)
	for today := 0; today < 31; today++ {
		event_date = event_date.Add(time.Hour * 24)
		row, err := s.db.Query("SELECT * FROM events where date=?", event_date.Format(time.DateOnly))
		if err != nil {
			return []strct.Event{}, fmt.Errorf("%s: prepare statement %w", op, err)
		}
		for row.Next() {
			item := strct.Event{}
			err := row.Scan(&item.Date, &item.Name, &item.Description)
			if err != nil {
				return []strct.Event{}, fmt.Errorf("%s: scanning row %w", op, err)
			}
			result = append(result, item)
		}
	}
	return result, nil
}
