package handlers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ClientRequest(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) < 1 || segments[0] == "" {
		http.Error(w, "Invalid slug path", http.StatusBadRequest)
		return
	}
	slug := segments[0]
	restPath := "/" + strings.Join(segments[1:], "/")
	TunnelMutex.Lock()
	tunnel := TunnelMap[slug]
	TunnelMutex.Unlock()
	fmt.Printf("🌐 Received request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

	if tunnel == nil {
		http.Error(w, "Tunnel not connected", http.StatusBadGateway)
		return
	}

	// Clone the request and modify the path
	reqClone := r.Clone(r.Context())
	reqClone.URL.Path = restPath
	reqClone.RequestURI = restPath // important for writing raw HTTP
	reqClone.Host = r.Host

	// Send request to tunnel
	err := reqClone.Write(tunnel) // write raw HTTP request to tunnel
	if err != nil {
		TunnelMutex.Lock()
		tunnel.Close()
		TunnelMutex.Unlock()
		delete(TunnelMap, slug)
		log.Printf("\n Closed tunnel: %v \n", tunnel)
		log.Println("Failed to write to tunnel:", err)
		http.Error(w, "Tunnel write failed", 500)
		return
	}

	// Read response from tunnel
	resp, err := http.ReadResponse(bufio.NewReader(tunnel), reqClone)
	if err != nil {
		log.Println("Failed to read from tunnel:", err)
		http.Error(w, "Tunnel read failed", 500)
		return
	}

	// Copy response back to original client
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
