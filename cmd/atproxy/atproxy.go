package main

import (
	"flag"
	"log"

	"github.com/atomicswe/atproxy/internal/config"
	"github.com/atomicswe/atproxy/internal/proxy"
	"github.com/atomicswe/atproxy/internal/request"
	"github.com/atomicswe/atproxy/internal/server"
)

func main() {
	address := flag.String("addr", "", "proxy address")
	port := flag.Int("port", 11111, "proxy port")
	flag.Parse()
	serverFlags := server.ServerFlags{
		Address: *address,
		Port:    *port,
	}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load configurations", err)
	}

	v := request.NewValidator(config.Validator)
	p := proxy.NewProxy(v)
	server.StartServer(p, serverFlags)
}
