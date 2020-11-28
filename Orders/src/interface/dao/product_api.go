package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Orders/src/interface/api_clients"
	"github.com/Orders/src/interface/config"
	"github.com/eapache/go-resiliency/breaker"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ProductApI struct {
	cfg            *config.Config
	circuitBreaker *breaker.Breaker
}

func NewProductAPI(cfg *config.Config, circuitBreaker *breaker.Breaker) *ProductApI {
	return &ProductApI{cfg: cfg, circuitBreaker: circuitBreaker}
}

type GetQuantityResponse struct {
	ProductID int `json:"id"`
	Stock     int `json:"stock"`
}

type SetQuantityResponse struct {
	ProductID int  `json:"id"`
	Success   bool `json:"success"`
}

func (p *ProductApI) GetQuantity(productId int) (int, error) {
	var productResponse GetQuantityResponse

	reqUrl := fmt.Sprintf("%v/%v/%v", p.cfg.Product.APIHost, "/api/quantity/", productId)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("[GetQuantity] error : ", err)
		return productResponse.Stock, errors.New("error creating request")
	}
	if api_clients.ProductAPI == nil {
		api_clients.ProductAPI = api_clients.CreateClient(10, 10, 10)
	}
	var resp *http.Response
	var errReq error
	req.Header.Set("Content-Type", "application/json")
	breakerResponse := p.circuitBreaker.Run(func() error {
		resp, errReq = api_clients.ProductAPI.Do(req) //Request to api
		if errReq != nil {
			log.Println("[GetQuantity] Error doing request to product microservice", errReq)
			return errReq
		}
		return nil
	})
	if breakerResponse != nil {
		if breakerResponse == breaker.ErrBreakerOpen {
			log.Println("[GetQuantity] error in circuit breaker[best match score]", breakerResponse)
			return productResponse.Stock, errors.New("error getting product quantity")
		}
		return productResponse.Stock, errors.New("error getting product quantity")
	}

	body, err := ioutil.ReadAll(resp.Body) //Reading response data body
	resp.Body.Close()
	if err != nil {
		log.Println("[GetQuantity] error : ", err)
		return productResponse.Stock, errors.New("error getting product quantity")
	}

	err = json.Unmarshal(body, &productResponse)
	if err != nil {
		log.Println("[GetQuantity] error : ", err)
		return productResponse.Stock, errors.New("error getting product quantity")
	}
	return productResponse.Stock, nil
}

func (p *ProductApI) UpdateQuantity(productId int, stockRemaining int) error {
	var res SetQuantityResponse
	reqUrl := fmt.Sprintf("%v/%v", p.cfg.Product.APIHost, productId)
	data := url.Values{}
	data.Set("stock", fmt.Sprintf("%v", stockRemaining))

	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println("[UpdateQuantity] error : ", err)
		return errors.New("error creating set quantity request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if api_clients.ProductAPI == nil {
		api_clients.ProductAPI = api_clients.CreateClient(10, 10, 10)
	}
	var resp *http.Response
	var errReq error
	breakerResponse := p.circuitBreaker.Run(func() error {
		resp, errReq = api_clients.ProductAPI.Do(req) //Request to api
		if errReq != nil {
			log.Println("[UpdateQuantity] Error doing request to product microservice", errReq)
			return errReq
		}
		return nil
	})
	if breakerResponse != nil {
		if breakerResponse == breaker.ErrBreakerOpen {
			log.Println("[UpdateQuantity] error in circuit breaker[best match score]", breakerResponse)
		}
		return errors.New("error getting product quantity")
	}

	body, err := ioutil.ReadAll(resp.Body) //Reading response data body
	resp.Body.Close()
	if err != nil {
		log.Println("[UpdateQuantity] error : ", err)
		return errors.New("error getting product quantity")
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("[UpdateQuantity] error : ", err)
		return errors.New("error getting product quantity")
	}
	if !res.Success {
		return errors.New("product quantity update failed")
	}
	return nil
}
