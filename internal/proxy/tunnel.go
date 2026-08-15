package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func tunnelConn(dst, src net.Conn) {
	io.Copy(dst, src)

	if tcp, ok := dst.(*net.TCPConn); ok {
		tcp.CloseWrite()
		return
	}
	dst.Close()
}

func tunnel(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	target, err := net.Dial("tcp", r.Host)
	if err != nil {
		log.Println("ERROR: failed to dial to target", r.Host)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	hj, ok := w.(http.Hijacker)
	if !ok {
		log.Println("ERROR: http server doesn't support hijacking connection")
	}

	client, _, err := hj.Hijack()
	if err != nil {
		log.Println("ERROR: http hijacking failed with error:", err)
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
	log.Printf("request to '%v' finished, took %dms", r.Host, time.Since(start).Milliseconds())
}
