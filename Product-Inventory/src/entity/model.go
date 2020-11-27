package entity

type GetProductQuantityResponse struct {
	ProductId int `json:"id"`
	Stock     int `json:"stock"`
}

type SetProductQuantityResponse struct {
	ProductId int  `json:"id"`
	Success   bool `json:"success"`
}
