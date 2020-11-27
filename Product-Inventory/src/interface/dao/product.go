package dao

import (
	"errors"
	"github.com/Product-Inventory/src/interface/config"
	"sync"
)

type ProductDao struct {
	ProductStore map[int]config.Product
	lock         *sync.RWMutex
}

func New(cfg *config.Config) *ProductDao {
	product := &ProductDao{
		ProductStore: make(map[int]config.Product),
		lock:         &sync.RWMutex{},
	}
	product.InitializeProducts(cfg)
	return product
}

func (p *ProductDao) InitializeProducts(cfg *config.Config) {
	products := cfg.Products
	for _, product := range products {
		p.ProductStore[product.Id] = product
	}
	return
}

func (p *ProductDao) GetProductQuantity(productID int) (int, error) {
	if p.ProductStore == nil {
		return 0, errors.New("product not initialised, connection error")
	}
	p.lock.RLock()
	product, found := p.ProductStore[productID]
	if found {
		p.lock.RUnlock()
		return product.Stock, nil
	}
	p.lock.RUnlock()
	return 0, errors.New("invalid product id")
}

func (p *ProductDao) SetProductQuantity(productID int, buyingStock int) error {
	if p.ProductStore == nil {
		return errors.New("product not initialised, connection error")
	}
	p.lock.Lock()
	product, found := p.ProductStore[productID]
	if found {
		remainingStock := product.Stock - buyingStock
		if remainingStock < 0 {
			p.lock.Unlock()
			return errors.New("stock not available to buy")
		}
		product.Stock = remainingStock
		p.ProductStore[productID] = product
		p.lock.Unlock()
		return nil
	}
	p.lock.Unlock()
	return errors.New("invalid product id")
}
