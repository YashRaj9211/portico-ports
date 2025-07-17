package main

import (
	"log"
	"net/http"
	"fmt"
	restapi "github.com/YashRaj9211/portico-ports/internal/api/rest-api"
)


const PORT = ":4000" 

func main(){
	mux := http.NewServeMux()
	restapi.RegisterRoutes(mux);
	fmt.Printf("🚀 Server started at http://localhost%v \n", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
