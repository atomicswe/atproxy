package proxy

import (
	"io"
	"log"
	"maps"
	"net/http"
	"time"
)

func forward(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 30 * time.Second,
	}

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
		log.Println("ERROR: failed to forward request with error:", err.Error())
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response headers and status
	maps.Copy(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	log.Printf("request to '%v' finished, took %dms", r.Host, time.Since(start).Milliseconds())
}
