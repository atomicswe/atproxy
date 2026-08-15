package main

import (
	"io"
	"log"
	"maps"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

func forward(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 30 * time.Second,
	}

	log.Printf("received request for: %s\n", r.URL)

	req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Copy request headers
	maps.Copy(req.Header, r.Header)

	// Forward request
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response headers and status
	maps.Copy(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func tunnelConn(dst, src net.Conn) {
	io.Copy(dst, src)

	if tcp, ok := dst.(*net.TCPConn); ok {
		tcp.CloseWrite()
		return
	}
	dst.Close()
}

func tunnel(w http.ResponseWriter, r *http.Request) {
	target, err := net.Dial("tcp", r.Host)
	if err != nil {
		log.Println("ERROR: failed to dial to target", r.Host)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	hj, ok := w.(http.Hijacker)
	if !ok {
		log.Fatal("ERROR: http server doesn't support hijacking connection")
	}

	client, _, err := hj.Hijack()
	if err != nil {
		log.Fatal("ERROR: http hijacking failed", err)
	}

	log.Printf("creating tunnel to '%v'", r.Host)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		tunnelConn(target, client)
		wg.Done()
	}()
	go func() {
		tunnelConn(client, target)
		wg.Done()
	}()
	wg.Wait()
	client.Close()
	target.Close()
}

func allowed(r *http.Request) bool {
	if strings.Contains(r.URL.Hostname(), "example.com") {
		return false
	}
	return true
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received '%s' request to '%v' (from '%v')", r.Method, r.Host, r.RemoteAddr)
		if !allowed(r) {
			log.Printf("request to '%v' is not allowed.", r.Host)
			return
		}
		if r.Method == http.MethodConnect {
			tunnel(w, r)
			return
		}
		forward(w, r)
	}

	log.Println("starting proxy server on port 11111")
	if err := http.ListenAndServe(":11111", http.HandlerFunc(handler)); err != nil {
		log.Fatal("failed to listen and serve with error: ", err)
	}
}
