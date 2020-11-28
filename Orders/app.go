package main

import (
	"github.com/Orders/src/interface/config"
	"github.com/Orders/src/interface/dao"
	"github.com/Orders/src/interface/web"
	"github.com/Orders/src/usecases"
	"github.com/eapache/go-resiliency/breaker"
	"log"
	"time"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Println("couldn't Initialize module config", err)
		return
	}
	interactorObj := generateInteractor(cfg)
	h := web.Handler{Cfg: cfg, Interactor: interactorObj}
	server := web.New(&h)
	go server.Run()

	err = <-server.ListenError()
	if err != nil {
		log.Fatal("Error starting web server, exiting gracefully:", err)
	}
}

func generateInteractor(cfg *config.Config) *web.Interactor {
	cartRepo := dao.NewCartAPI(cfg, breaker.New(cfg.CircuitBreaker.ErrorThreshold, cfg.CircuitBreaker.SuccessThreshold, time.Duration(cfg.CircuitBreaker.Timeout)))
	productRepo := dao.NewProductAPI(cfg, breaker.New(cfg.CircuitBreaker.ErrorThreshold, cfg.CircuitBreaker.SuccessThreshold, time.Duration(cfg.CircuitBreaker.Timeout)))
	orderRepo := dao.NewUserOrder()
	paymentRepo := dao.NewPayment(breaker.New(cfg.CircuitBreaker.ErrorThreshold, cfg.CircuitBreaker.SuccessThreshold, time.Duration(cfg.CircuitBreaker.Timeout)))
	var interactor = &web.Interactor{
		OrdersInteractor: usecases.Init(cartRepo, productRepo, orderRepo, paymentRepo),
	}
	return interactor
}
