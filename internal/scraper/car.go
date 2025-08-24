package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	PerPage int = 40
)

func GetPageData(page int, make string, model string, year string) (Root, error) {
	baseUrl := os.Getenv("BASE_URL")
	urlPath := os.Getenv("URL_PATH")
	base, err := url.Parse(baseUrl + urlPath)
	if err != nil {
		return Root{}, fmt.Errorf("couldn't parse base url, error: %w", err)
	}

	from := PerPage * page
	params := url.Values{}
	params.Set("category", "1020")
	params.Set("from", strconv.Itoa(from))
	params.Set("include", "extra_images,body")
	params.Set("limit", strconv.Itoa(PerPage))
	params.Set("make_id", "36")
	params.Set("mfg_year", fmt.Sprintf("%s-%s", year, year))
	params.Set("mileage", "-100000")
	params.Set("model_id", "1756")
	params.Set("price", "-100000")
	params.Set("type", "sell")
	base.RawQuery = params.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		slog.Error("couldn't perform GET request", "func", "GetPrices", "url", base, "error", err)
		return Root{}, fmt.Errorf("couldn't perform GET request, error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("couldn't read body of the response")
		return Root{}, fmt.Errorf("couldn't read body of the response")
	}

	var root Root
	if err := json.Unmarshal(body, &root); err != nil {
		slog.Error("couldn't unmarshall response body to a struct", "error", err)
	}

	slog.Debug("Meta data", "total-results", root.Meta.TotalResults, "total-showing", root.Meta.TotalShowing)

	return root, nil
}

func GetAllPrices(make string, model string, year string) ([]int, error) {
	// get prices from the first page + metadata to understand if need to fetch more pages
	pageData, err := GetPageData(0, make, model, year)
	if err != nil {
		return []int{}, fmt.Errorf("couldn't get data from the first page. Error: %w", err)
	}

	pages := (pageData.Meta.TotalResults - 1) / PerPage

	var prices []int
	if pages == 0 {
		prices = GetPricesFromData(pageData)
		return prices, nil
	}

	for i := range pages {
		pageNumber := i + 1
		pageData, err := GetPageData(pageNumber, make, model, year)
		if err != nil {
			return []int{}, fmt.Errorf("couldn't get data from the %d page. Error: %w", pageNumber, err)
		}
		prices = append(prices, GetPricesFromData(pageData)...)
	}
	return prices, nil
}

func GetPricesFromData(r Root) []int {
	var prices []int
	for _, ad := range r.Data {
		prices = append(prices, ad.Attributes.Price)
	}
	return prices
}
