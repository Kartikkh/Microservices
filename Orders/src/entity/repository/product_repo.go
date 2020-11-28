package repository

type ProductRepository interface {
	GetQuantity(productId int) (int,error)
	UpdateQuantity(productId int, stockRemaining int) error
}

