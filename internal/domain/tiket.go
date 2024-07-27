package domain

type Tiket struct {
	ID    int     `json:"id"`
	Stock int     `json:"stock"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
