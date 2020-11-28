package web

import (
	"encoding/json"
	"github.com/Orders/lib/response"
	"github.com/Orders/src/entity"
	"github.com/Orders/src/interface/config"
	"github.com/Orders/src/usecases"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Interactor struct {
	OrdersInteractor usecases.OrdersDefinition
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

	router.POST("/api/order/:product_id", handleNow(a.OrderProduct))

}

func (a *API) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *response.JSONResponse {
	return response.NewJSONResponse().SetData("pong !!").SetStatusCode(200).SetMessage("success")
}

func (a *API) OrderProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) *response.JSONResponse {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "check the data in the request body").SetStatusCode(http.StatusBadRequest)
	}
	var order entity.Product
	err = json.Unmarshal(body, &order)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "check the data in the request body").SetStatusCode(http.StatusBadRequest)
	}
	userId := r.Header.Get("UserId")
	order.ProductId, err = strconv.Atoi(ps.ByName("product_id"))
	if err != nil {
		return response.NewJSONResponse().SetError(err, "error while getting product-Id").SetStatusCode(http.StatusBadRequest)
	}
	err = a.Interactor.OrdersInteractor.OrderProduct(userId, order)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "error while order").SetStatusCode(http.StatusBadRequest)
	}
	resMsg := entity.OrderResponse{
		UserId:      userId,
		ProductId:   order.ProductId,
		BuyingStock: order.Quantity,
		Success:     true,
	}
	return response.NewJSONResponse().SetData(resMsg).SetMessage("order successfully placed").SetStatusCode(http.StatusOK)

}
