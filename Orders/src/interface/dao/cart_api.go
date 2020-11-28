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
)

type CartAPI struct {
	config         *config.Config
	circuitBreaker *breaker.Breaker
}

func NewCartAPI(config *config.Config, circuitBreaker *breaker.Breaker) *CartAPI {
	return &CartAPI{config: config, circuitBreaker: circuitBreaker}
}

type CartProductPresentResponse struct {
	ProductId int
	Present   bool
}



func (c *CartAPI) IsProductInCart(userId string, productId int) (bool, error) {
	var cartProductResponse CartProductPresentResponse
	reqUrl := fmt.Sprintf("%v/%v/%v", c.config.Cart.APIHost, "api/cart/status/", productId)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("[GetQuantity] error : ", err)
		return cartProductResponse.Present, errors.New("error creating request")
	}
	req.Header.Set("userId", userId)
	if api_clients.CartClient == nil {
		api_clients.CartClient = api_clients.CreateClient(10, 10, 10)
	}
	var resp *http.Response
	var errReq error
	req.Header.Set("Content-Type", "application/json")
	breakerResponse := c.circuitBreaker.Run(func() error {
		resp, errReq = api_clients.CartClient.Do(req) //Request to api
		if errReq != nil {
			log.Println("[GetQuantity] Error doing request to product microservice", errReq)
			return errReq
		}
		return nil
	})
	if breakerResponse != nil {
		if breakerResponse == breaker.ErrBreakerOpen {
			log.Println("[GetQuantity] error in circuit breaker[best match score]", breakerResponse)
			return cartProductResponse.Present, errors.New("error creating request")
		}
		return cartProductResponse.Present, errors.New("error creating request")
	}

	body, err := ioutil.ReadAll(resp.Body) //Reading response data body
	resp.Body.Close()
	if err != nil {
		log.Println("[GetQuantity] error : ", err)
		return cartProductResponse.Present, errors.New("error creating request")
	}

	err = json.Unmarshal(body, &cartProductResponse)
	if err != nil {
		log.Println("[GetQuantity] error : ", err)
		return cartProductResponse.Present, errors.New("error creating request")
	}
	return cartProductResponse.Present, nil
}
