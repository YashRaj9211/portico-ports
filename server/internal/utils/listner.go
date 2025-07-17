package utils

import (
	"log"
	"net"
)


func TCPListner(port string, slugI string) (tunnelConn net.Conn, slugO string) {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal("Tunnel listener error:", err)
		}
		log.Printf("📡 Waiting for tunnel on %v \n", port)
		tunnelConn, err = listener.Accept()

		if err != nil {
			log.Fatal("Failed to accept tunnel:", err)
		}
		log.Println("🔌 Tunnel connected")
		return tunnelConn, slugI
}