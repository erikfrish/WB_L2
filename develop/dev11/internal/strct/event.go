package strct

import "time"

type Event struct {
	Date        time.Time `json:"date" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description,omitempty" validate:"datetime"`
}
