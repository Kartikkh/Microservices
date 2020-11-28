package config

type Server struct {
	Address string `json:"address"`
}

type ProductService struct {
	APIHost string `json:"api_host"`
}

type CartService struct {
	APIHost string `json:"api_host"`
}
type CircuitBreaker struct {
	ErrorThreshold   int `json:"error_threshold"`
	SuccessThreshold int `json:"success_threshold"`
	Timeout          int `json:"timeout"`
}
type Config struct {
	Server         Server         `json:"server"`
	Product        ProductService `json:"product"`
	Cart           CartService    `json:"cart"`
	CircuitBreaker CircuitBreaker `json:"circuit_breaker"`
}
