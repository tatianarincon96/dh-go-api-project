package domain

type Compra struct {
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
}
