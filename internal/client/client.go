package client

import (
    "bytes"
    "encoding/json"
    "net/http"
    "os"
)

func ForwardPing(payload map[string]interface{}) error {
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return err
    }

    endpoint := os.Getenv("RAILS_API_URL")
    if endpoint == "" {
        endpoint = "http://localhost:3000/internal/ping-events"
    }

    resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return nil
}
