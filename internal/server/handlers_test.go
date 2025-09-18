package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	h := Handlers{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", h.PingHandler)
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

func TestBarsHandler(t *testing.T) {
	assert := assert.New(t)

	mockRenderer := NewMockRenderer(t)
	h := Handlers{renderer: mockRenderer}
	mockRenderer.EXPECT().RenderHTML("volkswagen", "passat", "2019").Return([]byte("ok"), nil)

	tests := []struct {
		name      string
		make      string
		model     string
		year      string
		status    int
		body      []byte
		wantError bool
	}{
		{
			name:   "successful request",
			make:   "volkswagen",
			model:  "passat",
			year:   "2019",
			status: http.StatusOK,
			body:   []byte("ok"),
		},
		{
			name:   "unsuccessful request",
			make:   "",
			model:  "passat",
			year:   "2019",
			status: http.StatusBadRequest,
			body:   []byte("missing 'make' query param\n"),
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /bars", h.BarsHandler)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rR := httptest.NewRecorder()
			baseURL := "/bars"
			queryParams := url.Values{}
			if test.make != "" {
				queryParams.Set("make", test.make)
			}
			if test.model != "" {
				queryParams.Set("model", test.model)
			}
			if test.year != "" {
				queryParams.Set("year", test.year)
			}
			req, err := http.NewRequest(http.MethodGet, baseURL, nil)
			if err != nil {
				t.Errorf("couldn't create http request to /bars, error: %v", err)
			}
			req.URL.RawQuery = queryParams.Encode()

			mux.ServeHTTP(rR, req)

			assert.Equal(test.status, rR.Code, "http response status correct")
			assert.Equal(test.body, rR.Body.Bytes(), "response body correct")

		})
	}
}
