package entity

type Transaction struct {
	ID    string  `json:"id"`
	To    string  `json:"to"`
	From  string  `json:"from"`
	Price float64 `json:"price"`
}
