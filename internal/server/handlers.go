package server

import (
	"encoding/json"
	"net/http"

	"github.com/2Cheetah/car-price-validator/internal/visualiser"
)

type PingResponse struct {
	Pong int `json:"pong"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if len(queryParams) != 0 {
		http.Error(w, "no query params allowed", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := PingResponse{Pong: 1}
	json.NewEncoder(w).Encode(response)
}

func BarsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if !queryParams.Has("make") {
		http.Error(w, "missing 'make' query param", http.StatusBadRequest)
		return
	}
	make := queryParams.Get("make")

	if !queryParams.Has("model") {
		http.Error(w, "missing 'model' query param", http.StatusBadRequest)
		return
	}
	model := queryParams.Get("model")

	if !queryParams.Has("year") {
		http.Error(w, "missing 'year' query param", http.StatusBadRequest)
		return
	}
	year := queryParams.Get("year")

	content, err := visualiser.RenderHTML(make, model, year)
	if err != nil {
		http.Error(w, "couldn't render bars", http.StatusInternalServerError)
	}
	w.Write(content)
}
