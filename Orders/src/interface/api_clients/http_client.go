package api_clients

import (
	"net/http"
	"time"
)

var (
	ProductAPI *http.Client
	CartClient *http.Client
)

func CreateClient(clientTimeout, maxIdleConnection, idleConnTimeout int) *http.Client {

	if clientTimeout == 0 { //default value
		clientTimeout = 250
	}
	if maxIdleConnection == 0 { //default value
		maxIdleConnection = 10
	}
	if idleConnTimeout == 0 { //default value
		idleConnTimeout = 30
	}

	return &http.Client{
		Timeout: time.Duration(clientTimeout) * time.Millisecond,
		Transport: &http.Transport{
			MaxIdleConns:        maxIdleConnection,
			MaxIdleConnsPerHost: maxIdleConnection,
			IdleConnTimeout:     time.Duration(idleConnTimeout) * time.Second,
		}} //creating http client
}
