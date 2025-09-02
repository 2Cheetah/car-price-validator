package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", PingHandler)
	respRecorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("couldn't create a GET request to /ping endpoint, error: %v+", err)
	}
	mux.ServeHTTP(respRecorder, req)

	if respRecorder.Code != http.StatusOK {
		t.Error("response http code is not 200")
	}

	var pingResponse PingResponse
	if err := json.Unmarshal(respRecorder.Body.Bytes(), &pingResponse); err != nil {
		t.Errorf("couldn't unmarshal response body to expected ping response, error: %v+", err)
	}

	if expected, got := 1, pingResponse.Pong; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}
}
