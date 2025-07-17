package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func ClientSideTunneler(localPort string) (url string, err error) {
	localPort = fmt.Sprintf(":%v", localPort)
	res, err := http.Get("http://localhost:4000/expose")
	if err != nil {
		log.Fatalf("Error while calling server")
		return	"", err
	}
	defer res.Body.Close()

	var parsed ExposeResponse
	err = json.NewDecoder(res.Body).Decode(&parsed)
	if err != nil || !parsed.Success {
		fmt.Print("Error while decoding resposne")
		return "",err
	}

	go func() {
		serverConn, err := net.Dial("tcp", "localhost"+parsed.AssignedPort)
		if err != nil {
			return 
		}
		localConn, err := net.Dial("tcp", "localhost"+localPort)
		if err != nil {
			return
		}
		go io.Copy(serverConn, localConn)
		io.Copy(localConn, serverConn)
	}()

	return parsed.PublicURL, nil
}


type ExposeResponse struct {
	Success      bool   `json:"success"`
	Error        string `json:"error,omitempty"`
	Slug         string `json:"slug,omitempty"`
	PublicURL    string `json:"publicUrl,omitempty"`
	AssignedPort string `json:"assignedPort,omitempty"`
}