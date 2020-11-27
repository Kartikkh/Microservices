package web

import (
	"encoding/json"
	"github.com/Product-Inventory/lib/response"
	"github.com/Product-Inventory/src/entity"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *API) SetQuantity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) *response.JSONResponse {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "check the data in the request body").SetStatusCode(http.StatusBadRequest)
	}

	var buyingStock int
	err = json.Unmarshal(body, &buyingStock)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "check the data in the request body").SetStatusCode(http.StatusBadRequest)
	}

	productID := ps.ByName("product_id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "").SetStatusCode(http.StatusBadRequest)
	}

	err = a.Interactor.ProductInteractor.SetQuantity(id, buyingStock)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "").SetStatusCode(http.StatusBadRequest)
	}

	resMsg := entity.SetProductQuantityResponse{
		ProductId: id,
		Success:   true,
	}
	return response.NewJSONResponse().SetData(resMsg).SetStatusCode(http.StatusOK)
}

func (a *API) GetQuantity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) *response.JSONResponse {
	w.Header().Set("Content-Type", "application/json")

	productID := ps.ByName("product_id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "").SetStatusCode(http.StatusBadRequest)
	}

	quantity, err := a.Interactor.ProductInteractor.GetQuantity(id)
	if err != nil {
		return response.NewJSONResponse().SetError(err, "").SetStatusCode(http.StatusBadRequest)
	}
	resMsg := entity.GetProductQuantityResponse{
		ProductId: id,
		Stock:     quantity,
	}
	return response.NewJSONResponse().SetData(resMsg).SetStatusCode(http.StatusOK)

}
