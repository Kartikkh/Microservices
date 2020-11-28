package repository

import "github.com/Orders/src/entity"

type OrderRepository interface {
	AllocateProductToUser(request *entity.Product, userId string) error
}
