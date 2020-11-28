package dao

import "github.com/eapache/go-resiliency/breaker"

type Payment struct {
	circuitBreaker *breaker.Breaker
}

func NewPayment(circuitBreaker *breaker.Breaker) *Payment {
	return &Payment{circuitBreaker: circuitBreaker}
}

func (p *Payment) PaymentTransaction(amount float64, method string) (bool, error) {
	// Make a API call to payment microservice
	return true, nil
}
