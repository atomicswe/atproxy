package proxy

import (
	"log"
	"net/http"
	"time"

	"github.com/atomicswe/atproxy/internal/request"
)

type requestsCounter struct {
	allowedRequests int32
	deniedRequests  int32
}

type Proxy struct {
	StartTime time.Time
	client    *http.Client
	validator *request.Validator
	counter   requestsCounter
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
		p.counter.deniedRequests++
		return
	}
	p.counter.allowedRequests++
	if r.Method == http.MethodConnect {
		tunnel(w, r)
		return
	}
	forward(w, r)
}

func (p *Proxy) Finish() {
	since := time.Since(p.StartTime).Round(time.Second)
	days := since / (24 * time.Hour)
	since -= days * 24 * time.Hour

	hours := since / time.Hour
	since -= hours * time.Hour

	minutes := since / time.Minute
	since -= minutes * time.Minute

	seconds := since / time.Second
	log.Printf("closing proxy. Allowed %d requests, denied %d requests. Total uptime: %d days, %d hours, %d minutes and %d seconds",
		p.counter.allowedRequests, p.counter.deniedRequests, days, hours, minutes, seconds)
}
