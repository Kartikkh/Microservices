package usecases

import "github.com/Product-Inventory/src/entity/repository"

type Product struct {
	ProductRepo repository.ProductRepository
}

func Init(repo repository.ProductRepository) *Product {
	return &Product{
		ProductRepo: repo,
	}
}

func (p *Product) GetQuantity(productID int) (int, error) {
	return p.ProductRepo.GetProductQuantity(productID)
}

func (p *Product) SetQuantity(productID int, stock int) error {
	return p.ProductRepo.SetProductQuantity(productID, stock)
}
