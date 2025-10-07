package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/2Cheetah/car-price-validator/internal/visualiser"
)

type Handlers struct {
	renderer Renderer
}

type Renderer interface {
	RenderHTML(make string, model string, year string) ([]byte, error)
}

type PingResponse struct {
	Pong int `json:"pong"`
}

func NewHandlers() *Handlers {
	return &Handlers{
		renderer: &visualiser.Visualiser{},
	}
}

func (h *Handlers) PingHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if len(queryParams) != 0 {
		http.Error(w, "no query params allowed", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := PingResponse{Pong: 1}
	json.NewEncoder(w).Encode(response)
}

func (h *Handlers) BarsHandler(w http.ResponseWriter, r *http.Request) {
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

	content, err := h.renderer.RenderHTML(make, model, year)
	if err != nil {
		slog.Error("error while trying to render bars", "error", err)
		http.Error(w, "couldn't render bars", http.StatusInternalServerError)
		return
	}
	w.Write(content)
}
