package usecases

import "github.com/Cart/src/entity"

type CartDefinition interface {
	AddProductToCart(request *entity.Product, userId string) error
}
