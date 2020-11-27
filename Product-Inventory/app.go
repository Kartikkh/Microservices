package main

import (
	"github.com/Product-Inventory/src/interface/config"
	"github.com/Product-Inventory/src/interface/dao"
	"github.com/Product-Inventory/src/interface/web"
	"github.com/Product-Inventory/src/usecases"
	"log"
)

func main() {

	// Read Config
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
	var interactor = &web.Interactor{
		ProductInteractor: usecases.Init(dao.New(cfg)),
	}
	return interactor
}
