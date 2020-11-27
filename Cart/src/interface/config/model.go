package config

type Server struct {
	Address string `json:"address"`
}

type Config struct {
	Server   Server    `json:"server"`
}
