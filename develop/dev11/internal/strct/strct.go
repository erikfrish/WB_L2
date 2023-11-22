package strct

type Event struct {
	Date string `json:"date" validate:"required"`
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc,omitempty"`
}
