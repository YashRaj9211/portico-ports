package handlers

import (
	"fmt"
	"net/http"
)

func HelloFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello server is up and running! 😁")
}