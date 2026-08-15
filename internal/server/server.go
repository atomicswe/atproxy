package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atomicswe/atproxy/internal/proxy"
)

const (
	defaultPort   = 11111
	defaultAdress = "" // intentionally empty "" => localhost
)

type ServerConfig struct {
	Port    int
	Address string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Port:    defaultPort,
		Address: defaultAdress,
	}
}

func StartServer(p *proxy.Proxy, s ServerConfig) {
	addr := fmt.Sprintf("%s:%d", s.Address, s.Port)
	log.Println("starting proxy server on adress:", addr)
	if err := http.ListenAndServe(addr, p); err != nil {
		log.Fatal("failed to listen and serve with error: ", err)
	}
}
