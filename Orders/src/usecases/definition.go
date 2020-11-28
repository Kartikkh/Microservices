package usecases

import "github.com/Orders/src/entity"

type OrdersDefinition interface {
	OrderProduct(userId string, product entity.Product) error
}
