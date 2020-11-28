package entity

type Product struct {
	ProductId int     `json:"id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderResponse struct {
	UserId      string  `json:"user_id"`
	ProductId int     `json:"id"`
	BuyingStock int  `json:"buying_stock"`
	Success bool `json:"success"`
}
