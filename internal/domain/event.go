package domain

type Event struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" valo:"notblank,sizeMin=2,sizeMax=50"`
	Date     string  `json:"date" valo:"notblank"`
	Location string  `json:"location" valo:"notblank"`
	Tiket    []Tiket `json:"tiket" valo:"notnil"`
}
