package usecases

import (
	"errors"
	"github.com/Orders/src/entity"
	"github.com/Orders/src/entity/repository"
	"log"
)

type OrdersManagement struct {
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
	paymentRepo repository.PaymentRepository
}

func Init(cartRepo repository.CartRepository, productRepo repository.ProductRepository, orderRepo repository.OrderRepository, paymentRepo repository.PaymentRepository) OrdersDefinition {
	return &OrdersManagement{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
		paymentRepo: paymentRepo,
	}
}

func (o *OrdersManagement) updateProductQuantity(productId, stockRemaining int) error {
	err := o.productRepo.UpdateQuantity(productId, stockRemaining)
	if err != nil {
		log.Println("error while updating quantity", err)
		return err
	}
	return nil
}

func (o *OrdersManagement) OrderProduct(userId string, product entity.Product) error {
	isPresent, err := o.cartRepo.IsProductInCart(userId, product.ProductId)
	if err != nil {
		log.Println("error while getting product from cart", err)
		return err
	}
	if !isPresent {
		return errors.New("the product you are trying to order is not present in the cart")
	}
	stockAvailable, err := o.productRepo.GetQuantity(product.ProductId)
	if err != nil {
		log.Println("error while getting stock", err)
		return err
	}
	if stockAvailable >= product.Quantity && product.Quantity > 0 {
		stockRemaining := stockAvailable - product.Quantity
		err = o.updateProductQuantity(product.ProductId, stockRemaining)
		if err != nil {
			return errors.New("try again later")
		}
		totalAmount := float64(product.Quantity) * product.Price
		success, err := o.paymentRepo.PaymentTransaction(totalAmount, "card")
		if err != nil || !success {
			// add a retry here -
			stockRemaining := stockAvailable - product.Quantity
			err = o.updateProductQuantity(product.ProductId, stockRemaining)
			if err != nil {
				// add slack alert 
				return errors.New("try again later")
			}
			log.Printf("error %T while doing payments for productId %d\n", err, product.ProductId)
			return err
		}
		return nil
	}
	return errors.New("stock not available")
}
