package dao

import (
	"errors"
	"github.com/Cart/src/entity"
	"github.com/Cart/src/entity/repository"
)

type Cart struct {
	UserCartStore map[string]*entity.Cart
}

func New() repository.CartRepository {
	return &Cart{
		UserCartStore: make(map[string]*entity.Cart, 0),
	}
}

func (c *Cart) CheckProductInCart(productId int, userId string) (bool, error) {
	if c.UserCartStore == nil {
		return false, errors.New("error initialising cart")
	}
	cart, found := c.UserCartStore[userId]
	if found {
		for _, product := range cart.Products {
			if productId == product.ProductId {
				return true, nil
			}
		}
	}
	return false, nil
}

func (c *Cart) AddToCart(product *entity.Product, userId string) error {
	if c.UserCartStore == nil {
		return errors.New("error initialising cart")
	}
	cart, found := c.UserCartStore[userId]
	if found {
		productAlreadyInCart := false
		for _, productInCart := range cart.Products {
			if productInCart.ProductId == product.ProductId {
				productAlreadyInCart = true
				product.Quantity += productInCart.Quantity
				break
			}
		}
		if !productAlreadyInCart {
			cart.Products = append(cart.Products, product)
		}
		c.UserCartStore[userId] = cart
		return nil
	} else {
		products := make([]*entity.Product, 0)
		products = append(products, product)
		c.UserCartStore[userId] = &entity.Cart{Products: products}
		return nil
	}
}
