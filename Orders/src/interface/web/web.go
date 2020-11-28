package web

import (
	"github.com/Orders/src/interface/config"
	"log"
	"net/http"
)

type Handler struct {
	Cfg         *config.Config
	Interactor  *Interactor
	listenErrCh chan error
}

//New is the web handler initializer
func New(this *Handler) *Handler {
	a := &API{Cfg: this.Cfg, Interactor: this.Interactor}
	routes(a).Register()
	return this
}

//Run is to run the web apis
func (h *Handler) Run() {
	log.Printf("Listening on %s", h.Cfg.Server.Address)
	server := &http.Server{
		Handler: getRouter(),
		Addr:    h.Cfg.Server.Address,
	}
	h.listenErrCh <- server.ListenAndServe()
}

//ListenError will lister the error
func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
