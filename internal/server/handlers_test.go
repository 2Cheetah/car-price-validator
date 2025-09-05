package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", PingHandler)
	respRecorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("couldn't create a GET request to /ping endpoint, error: %v+", err)
	}

	// Act
	mux.ServeHTTP(respRecorder, req)

	// Assert
	assert.Equal(http.StatusOK, respRecorder.Code, "response http status is 200")

	var pingResponse PingResponse
	if err := json.Unmarshal(respRecorder.Body.Bytes(), &pingResponse); err != nil {
		t.Errorf("couldn't unmarshal response body to expected ping response, error: %v+", err)
	}

	if assert.NotEmpty(pingResponse) {
		assert.Equal(1, pingResponse.Pong, "pong response received")
	}
}

// func TestBarsHandler(t *testing.T) {
// 	h := NewHandlers(nil)
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET /bars", h.BarsHandler)

// 	tests := []struct {
// 		name           string
// 		make           string
// 		model          string
// 		year           string
// 		expectedStatus int
// 	}{
// 		{
// 			name:           "success /bars?make=volkswagen&model=passat&year=2019",
// 			make:           "volkswagen",
// 			model:          "passat",
// 			year:           "2019",
// 			expectedStatus: http.StatusOK,
// 		},
// 		{
// 			name:           "empty make /bars?make=&model=passat&year=2019",
// 			make:           "",
// 			model:          "passat",
// 			year:           "2019",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "missing make /bars?model=passat&year=2019",
// 			model:          "passat",
// 			year:           "2019",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			respRec := httptest.NewRecorder()

// 			queryParams := url.Values{}
// 			if test.make != "" {
// 				queryParams.Set("make", test.make)
// 			}
// 			if test.model != "" {
// 				queryParams.Set("model", test.model)
// 			}
// 			if test.year != "" {
// 				queryParams.Set("year", test.year)
// 			}

// 			req, err := http.NewRequest(http.MethodGet, "/bars", nil)
// 			if err != nil {
// 				t.Fatalf("couldn't create a GET request to /bars endpoint, error: %v+", err)
// 			}
// 			req.URL.RawQuery = queryParams.Encode()

// 			mux.ServeHTTP(respRec, req)

// 			if respRec.Code != test.expectedStatus {
// 				t.Errorf("expected status: %d, got status: %d", test.expectedStatus, respRec.Code)
// 			}
// 		})
// 	}
// }
