package router

import (
    "github.com/gorilla/mux"
    "pingreceiver/internal/handlers"  // 👈 import the handlers package
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/ping/{token}", handlers.HandlePing).Methods("POST")
    return r
}
