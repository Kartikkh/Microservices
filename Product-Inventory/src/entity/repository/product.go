package repository

import "github.com/Product-Inventory/src/interface/config"

type ProductRepository interface {
	InitializeProducts(config *config.Config)
	GetProductQuantity(productID int) (int, error)
	SetProductQuantity(productID int, buyingStock int) error
}
