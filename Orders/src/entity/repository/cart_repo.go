package repository

type CartRepository interface {
	IsProductInCart(userID string, productId int) (bool, error)
}
