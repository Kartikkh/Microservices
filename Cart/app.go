package main

import (
	"github.com/Cart/src/interface/config"
	"github.com/Cart/src/interface/dao"
	"github.com/Cart/src/interface/web"
	"github.com/Cart/src/usecases"
	"log"
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
	var interactor = &web.Interactor{
		CartInteractor: usecases.Init(dao.New()),
	}
	return interactor
}
