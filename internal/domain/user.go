package domain

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name" valo:"notblank"`
	Address string  `json:"address" valo:"notblank"`
	Balance float64 `json:"balance,omitempty"`
}
