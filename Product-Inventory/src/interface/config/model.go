package config

type Server struct {
	Address string `json:"address"`
}

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"string"`
	Stock int    `json:"stock"`
}

type Config struct {
	Server   Server    `json:"server"`
	Products []Product `json:"products"`
}
