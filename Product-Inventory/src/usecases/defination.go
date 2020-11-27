package usecases

type ProductDefinition interface {
	GetQuantity(productId int) (int, error)
	SetQuantity(productId int, stock int) error
}
