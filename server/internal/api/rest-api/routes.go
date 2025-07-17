package restapi

import (
	"net/http"
	"github.com/YashRaj9211/portico-ports/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux){
	mux.HandleFunc("/hello", handlers.HelloFunc)
	mux.HandleFunc("/expose", handlers.ExposeHandler)
	mux.HandleFunc("/", handlers.ClientRequest)
}