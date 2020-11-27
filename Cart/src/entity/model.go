package entity

type Product struct {
	ProductId int     `json:"id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	Products []*Product `json:"products"`
}
