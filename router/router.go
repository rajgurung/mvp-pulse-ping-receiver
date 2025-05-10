package router

import (
	"github.com/gorilla/mux"
	"pingreceiver/internal/client"
	"pingreceiver/internal/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	pinger := &client.HttpPinger{}
	handler := &handlers.PingHandler{Pinger: pinger}

	r.HandleFunc("/ping/{token}", handler.HandlePing).Methods("POST")
	return r
}
