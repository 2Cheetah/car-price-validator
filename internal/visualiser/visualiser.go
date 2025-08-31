package visualiser

import (
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/2Cheetah/car-price-validator/internal/domain"
	"github.com/2Cheetah/car-price-validator/internal/repository"
	"github.com/2Cheetah/car-price-validator/internal/scraper"
	"github.com/2Cheetah/car-price-validator/internal/statistics"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	YearFirstCarProduced int = 1900
)

func RenderHTML(make string, model string, year string) ([]byte, error) {
	// validate input params
	if err := ValidateMake(make); err != nil {
		return []byte{}, err
	}

	if err := ValidateModel(make, model); err != nil {
		return []byte{}, err
	}

	if err := ValidateYear(year); err != nil {
		return []byte{}, err
	}

	// get prices
	prices, err := scraper.GetAllPrices(make, model, year)
	if err != nil {
		slog.Error("error while trying to get prices", "error", err)
		return []byte{}, fmt.Errorf("couldn't render HTML - no prices fetched, error: %w", err)
	}

	// convert to domain.barData
	domainBarData := statistics.PricesToBarData(prices)

	// convert domain.barData to opts.barData
	barData := GenerateBarData(domainBarData)

	// render the bars
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Car price distribution",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width: "1200px",
		}),
	)

	xAxisLabels := XAxisLabels(domainBarData)
	bar.SetXAxis(xAxisLabels).
		AddSeries("Car price distribution", barData)
	content := bar.RenderContent()

	return content, nil
}

func ValidateMake(make string) error {
	carMakes, err := repository.GetCarMakes()
	if err != nil {
		return fmt.Errorf("couldn't get car makes, error: %w", err)
	}

	carMake := strings.ToLower(make)
	if !slices.Contains(carMakes, carMake) {
		slog.Debug("make is not in carMakes", "make", carMake)
		return fmt.Errorf("unsupported car make")
	}
	return nil
}

func ValidateModel(make string, model string) error {
	carModels, err := repository.GetCarModelsByMake(make)
	if err != nil {
		return fmt.Errorf("couldn't get car models by make, error: %w", err)
	}

	model = strings.ToLower(model)
	if !slices.Contains(carModels, model) {
		slog.Debug("requested model is not in carModels", "model", model)
		return fmt.Errorf("unsupported car model")
	}
	return nil
}

func ValidateYear(year string) error {
	yInt, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("couldn't convert string year to a number")
	}

	if yInt < YearFirstCarProduced {
		return fmt.Errorf("are you sure cars where produced back then?")
	}

	currentYear := time.Now().Year()
	if yInt > currentYear {
		return fmt.Errorf("can't look into the future")
	}

	return nil
}

func XAxisLabels(barData []domain.BarData) []string {
	var labels []string
	for _, bar := range barData {
		labels = append(labels, bar.Name)
	}
	return labels
}

func GenerateBarData(barData []domain.BarData) []opts.BarData {
	b := make([]opts.BarData, 0)
	for _, bar := range barData {
		b = append(b, opts.BarData{Value: bar.Items})
	}
	return b
}

// func GenerateHTML(XAxisLabels []string, barData []opts.BarData) []byte {
// 	bar := charts.NewBar()
// 	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
// 		Title: "Prices distribution",
// 	}))
// 	bar.SetXAxis(XAxisLabels).
// 		AddSeries("Car prices distribution", barData)

// 	return bar.RenderContent()
// }
