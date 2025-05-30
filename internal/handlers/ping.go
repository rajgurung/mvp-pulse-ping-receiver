package handlers

import (
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "pingreceiver/internal/client"
)

type PingHandler struct {
    Pinger client.Pinger
}

func (ph *PingHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    token := vars["token"]

    if token == "" {
        http.Error(w, "Missing token", http.StatusBadRequest)
        return
    }

    payload := map[string]interface{}{
        "token":     token,
        "timestamp": time.Now().UTC(),
    }

    err := ph.Pinger.ForwardPing(payload)
    if err != nil {
        http.Error(w, "Failed to forward ping", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}
