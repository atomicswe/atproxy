package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atomicswe/atproxy/internal/proxy"
)

type ServerFlags struct {
	Port    int
	Address string
}

func StartServer(p *proxy.Proxy, s ServerFlags) {
	addr := fmt.Sprintf("%s:%d", s.Address, s.Port)
	log.Println("starting proxy server on adress:", addr)
	if err := http.ListenAndServe(addr, p); err != nil {
		log.Fatal("failed to listen and serve with error: ", err)
	}
}
