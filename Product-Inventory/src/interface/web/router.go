package web

import (
	"github.com/Product-Inventory/lib/response"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) *response.JSONResponse

var HttpRouter *httprouter.Router

func init() {
	HttpRouter = httprouter.New()
}

func getRouter() *httprouter.Router {
	return HttpRouter
}

func handleNow(handlerFunc Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		res := handlerFunc(writer, request, params)
		if res.Error != nil {
			res.SetStatusCode(http.StatusInternalServerError)
			log.Println("error while making request", res.Error)
		}
		res.Send(writer)
		return
	}
}
