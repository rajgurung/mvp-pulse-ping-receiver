package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockPinger struct {
	Called bool
}

func (m *mockPinger) ForwardPing(payload map[string]interface{}) error {
	m.Called = true
	return nil
}

func TestHandlePing(t *testing.T) {
	mock := &mockPinger{}
	handler := &PingHandler{Pinger: mock}

	req := httptest.NewRequest(http.MethodPost, "/ping/test-token", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/ping/{token}", handler.HandlePing).Methods("POST")
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	if rr.Body.String() != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", rr.Body.String())
	}

	if !mock.Called {
		t.Error("Expected ForwardPing to be called")
	}
}
