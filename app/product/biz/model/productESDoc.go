package model

type ProductESDoc struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Stock    int32   `json:"stock"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
}