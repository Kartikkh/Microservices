package web

import (
	"encoding/json"
	"github.com/Cart/lib/response"
	"github.com/Cart/src/entity"
	"github.com/Cart/src/interface/config"
	"github.com/Cart/src/usecases"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Interactor struct {
	CartInteractor usecases.CartDefinition
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

	router.POST("/api/cart/add/:user_id", handleNow(a.AddToCart))
	router.GET("/api/cart/status/:product_id", handleNow(a.IsProductInCart))

}

func (a *API) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *response.JSONResponse {
	return response.NewJSONResponse().SetData("pong !!").SetStatusCode(200).SetMessage("success")
}

func (a *API) IsProductInCart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) *response.JSONResponse {
	userId := r.Header.Get("UserId")
	productId := ps.ByName("product_id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "").SetStatusCode(http.StatusBadRequest)
	}
	isPresent, err := a.Interactor.CartInteractor.IsProductInCart(id, userId)
	resMsg := entity.IsProductPresentRes{
		IsPresent: isPresent,
		ProductId: id,
	}
	return response.NewJSONResponse().SetData(resMsg).SetStatusCode(http.StatusOK)
}

func (a *API) AddToCart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) *response.JSONResponse {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "check the data in the request body").SetStatusCode(http.StatusBadRequest)
	}
	var cartReq entity.Product
	err = json.Unmarshal(body, &cartReq)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "check the data in the request body").SetStatusCode(http.StatusBadRequest)
	}
	userId := r.Header.Get("UserId")
	err = a.Interactor.CartInteractor.AddProductToCart(&cartReq, userId)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "error adding to cart").SetStatusCode(http.StatusBadRequest)
	}
	return response.NewJSONResponse().SetMessage("success").SetStatusCode(http.StatusOK)
}
