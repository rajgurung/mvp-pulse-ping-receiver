package main

import (
    "log"
    "net/http"
    "pingreceiver/router"     // ← Import your router package
)

func main() {
    r := router.NewRouter()
    log.Println("Ping Receiver listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
