package repository

import "github.com/Cart/src/entity"

type CartRepository interface {
	AddToCart(request *entity.Product, userId string) error
	CheckProductInCart(productId int, userId string) (bool, error)
}
