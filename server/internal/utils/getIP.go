package utils

import (
	"net/http"
	"strings"
	"net"
)
func GetIP(r *http.Request) string {
	// Check X-Forwarded-For header (set by reverse proxies)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}

	// Fallback: use remote address
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return strings.Trim(host, "[]")
}
