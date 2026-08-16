package main

import (
	"log"
	"os"
	"os/signal"

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig != os.Interrupt && sig != os.Kill {
				continue
			}
			p.Finish()
			os.Exit(0)
		}
	}()

	server.StartServer(p, config.Server)
}
