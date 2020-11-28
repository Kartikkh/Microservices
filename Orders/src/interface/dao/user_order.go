package dao

import (
	"github.com/Orders/src/entity"
	"sync"
)

type UserOrder struct {
	UserOrderMapping map[string][]*entity.Product
	locker           *sync.RWMutex
}

func NewUserOrder() *UserOrder {
	return &UserOrder{
		UserOrderMapping: make(map[string][]*entity.Product, 0),
		locker:           &sync.RWMutex{},
	}
}

func (u *UserOrder) AllocateProductToUser(request *entity.Product, userId string) error {
	u.locker.Lock()
	productList, ok := u.UserOrderMapping[userId]
	if !ok {
		productList = make([]*entity.Product, 0)
	}
	productList = append(productList, request)
	u.UserOrderMapping[userId] = productList
	u.locker.Unlock()
	return nil
}
