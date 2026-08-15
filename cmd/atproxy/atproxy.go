package main

import (
	"log"

	"github.com/atomicswe/atproxy/internal/config"
	"github.com/atomicswe/atproxy/internal/proxy"
	"github.com/atomicswe/atproxy/internal/request"
	"github.com/atomicswe/atproxy/internal/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load configurations", err)
	}

	v := request.NewValidator(config.Validator)
	p := proxy.NewProxy(v)
	server.StartServer(p, config.Server)
}
