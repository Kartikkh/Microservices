package entity

type Product struct {
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	Products []*Product `json:"products"`
}

type IsProductPresentRes struct {
	IsPresent bool `json:"is_present"`
	ProductId int  `json:"product_id"`
}
