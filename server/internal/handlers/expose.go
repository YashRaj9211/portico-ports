package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/YashRaj9211/portico-ports/internal/utils"
	"net"
	"net/http"
	"sync"
	"time"
)

type ExposeRequest struct {
	Protocol  string `json:"protocol"`
	LocalPort int    `json:"localPort"`
	UserID    string `json:"userId"`
}

type ExposeResponse struct {
	Success      bool   `json:"success"`
	Error        string `json:"error,omitempty"`
	Slug         string `json:"slug,omitempty"`
	PublicURL    string `json:"publicUrl,omitempty"`
	AssignedPort string `json:"assignedPort,omitempty"`
}

var (
	TunnelMap   = make(map[string]net.Conn)
	TunnelMutex = sync.RWMutex{}
)


func ExposeHandler(w http.ResponseWriter, r *http.Request) {
	// var req ExposeRequest

	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 	return
	// }
	ip := utils.GetIP(r)
	slug, err := utils.GenerateSlug(ip, "readable", 0, fmt.Sprint(time.Now().UnixNano()))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Faild to assign unique url", http.StatusInternalServerError)
		return
	}
	publicPort := utils.GetPort()
	fmt.Printf("Assigned port: %s\n", publicPort)
	publicURL := fmt.Sprintf("http://portico-ports.ddns.net/%s", slug)

	go func(publicPort string, slug string) {
		tunnelConn, _ := utils.TCPListner(publicPort, slug)
		TunnelMutex.Lock()
		TunnelMap[slug] = tunnelConn
		TunnelMutex.Unlock()
		fmt.Println("Tunnel listning")
	}(publicPort, slug)

	response := ExposeResponse{
		Success:      true,
		PublicURL:    publicURL,
		Slug:         slug,
		AssignedPort: publicPort,
		Error:        "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
