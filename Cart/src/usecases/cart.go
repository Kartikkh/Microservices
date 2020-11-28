package usecases

import (
	"github.com/Cart/src/entity"
	"github.com/Cart/src/entity/repository"
)

type Cart struct {
	CartRepo repository.CartRepository
}

func (c *Cart) AddProductToCart(product *entity.Product, userId string) error {
	return c.CartRepo.AddToCart(product, userId)
}

func (c *Cart) IsProductInCart(productId int, userId string) (bool, error) {
	isPresent, err := c.CartRepo.CheckProductInCart(productId, userId)
	return isPresent, err
}

func Init(repo repository.CartRepository) *Cart {
	return &Cart{
		CartRepo: repo,
	}
}
