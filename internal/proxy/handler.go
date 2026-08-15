package proxy

import (
	"log"
	"net/http"

	"github.com/atomicswe/atproxy/internal/request"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received '%s' request to '%v' (from '%v')", r.Method, r.Host, r.RemoteAddr)
	if !request.Allowed(r) {
		log.Printf("request to '%v' is not allowed.", r.Host)
		return
	}
	if r.Method == http.MethodConnect {
		tunnel(w, r)
		return
	}
	forward(w, r)
}
