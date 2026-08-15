package proxy

import (
	"log"
	"net/http"

	"github.com/atomicswe/atproxy/internal/request"
)

type Proxy struct {
	client    *http.Client
	validator *request.Validator
}

func NewProxy(validator *request.Validator) *Proxy {
	return &Proxy{
		client:    &http.Client{},
		validator: validator,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received '%s' request to '%v' (from '%v')", r.Method, r.Host, r.RemoteAddr)
	if !p.validator.Allowed(r) {
		log.Printf("request to '%v' is not allowed", r.Host)
		http.Error(w, "The request domain is not allowed to pass through this proxy.", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodConnect {
		tunnel(w, r)
		return
	}
	forward(w, r)
}
