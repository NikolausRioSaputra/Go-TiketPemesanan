package domain

type Order struct {
	ID     int     `json:"id"`
	Date   string  `json:"date" valo:"notblank"`
	Status string  `json:"status" valo:"notblank"`
	User   User    `json:"user" valo:"valid"`
	Event  Event   `json:"event" valo:"valid"`
	Tiket  []Tiket `json:"tiket" valo:"notnil"`
	Total  float64 `json:"total"`
}
