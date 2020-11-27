package web

import (
	"github.com/Product-Inventory/lib/response"
	"github.com/Product-Inventory/src/interface/config"
	"github.com/Product-Inventory/src/usecases"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Interactor struct {
	ProductInteractor usecases.ProductDefinition
}

//API is the api struct
type API struct {
	Cfg        *config.Config
	Interactor *Interactor
}

//New is the api initializer
func routes(this *API) *API {
	return &API{Cfg: this.Cfg, Interactor: this.Interactor}
}

func (a *API) Register() {
	router := getRouter()
	router.GET("/ping", handleNow(a.Ping))

	// upload products
	router.POST("/api/quantity/:product_id", handleNow(a.SetQuantity))
	router.GET("/api/quantity/:product_id", handleNow(a.GetQuantity))

}

func (a *API) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *response.JSONResponse {
	return response.NewJSONResponse().SetData("pong !!").SetStatusCode(200).SetMessage("success")
}
