package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func GetPrices(make string, model string, year string) ([]int, error) {
	baseUrl := os.Getenv("BASE_URL")
	urlPath := os.Getenv("URL_PATH")
	filters := "&from=0&include=extra_images%2Cbody&limit=40&make_id=28&mfg_year=2020-2020&mileage=-100000&model_id=1351&price=-100000&type=sell"
	url := baseUrl + urlPath + filters
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("couldn't perform GET request", "func", "GetPrices", "url", url, "error", err)
		return []int{}, fmt.Errorf("couldn't perform GET request, error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("couldn't read body of the response")
		return []int{}, fmt.Errorf("couldn't read body of the response")
	}

	var root Root
	if err := json.Unmarshal(body, &root); err != nil {
		slog.Error("couldn't unmarshall response body to a struct", "error", err)
	}

	var prices []int
	for _, ad := range root.Data {
		prices = append(prices, ad.Attributes.Price)
	}

	return prices, nil
}
