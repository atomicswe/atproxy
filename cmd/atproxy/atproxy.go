package main

import (
	"log"
	"net/http"

	"github.com/atomicswe/atproxy/internal/proxy"
)

func main() {
	log.Println("starting proxy server on port 11111")
	if err := http.ListenAndServe(":11111", http.HandlerFunc(proxy.Handler)); err != nil {
		log.Fatal("failed to listen and serve with error: ", err)
	}
}
