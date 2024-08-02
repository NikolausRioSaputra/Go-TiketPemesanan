package domain

type Tiket struct {
	ID    int     `json:"id"`
	Stock int     `json:"stock" valo:"min=1"`
	Type  string  `json:"type" valo:"notblank"`
	Price float64 `json:"price" valo:"min=1"`
}
