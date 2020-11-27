package repository

import "github.com/Cart/src/entity"

type CartRepository interface {
	AddToCart(request *entity.Product, userId string) error
}

