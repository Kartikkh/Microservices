package repository

type PaymentRepository interface {
	PaymentTransaction(amount float64, method string) (bool, error)
}
