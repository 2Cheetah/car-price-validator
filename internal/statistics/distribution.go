package statistics

import (
	"fmt"
	"log/slog"
	"math"
	"sort"

	"github.com/2Cheetah/car-price-validator/internal/domain"
)

const (
	RangeWidth     int = 5000
	NumberOfRanges int = 10
)

func PricesToBarData(prices []int) []domain.BarData {
	if len(prices) == 0 {
		slog.Warn("prices slice is empty", "func", "PricesToBarData")
		return []domain.BarData{}
	}
	sort.Ints(prices)
	slog.Debug("prices", "values", prices)
	minPrice := prices[0]
	slog.Debug("minPrice", "value", minPrice)
	maxPrice := prices[len(prices)-1]
	slog.Debug("maxPrice", "value", maxPrice)
	lowerLimit := math.Round((float64(minPrice)/5000)-0.51) * 5000
	slog.Debug("lowerLimit", "value", lowerLimit)
	upperLimit := math.Round((float64(maxPrice)/5000)+0.5) * 5000
	slog.Debug("upperLimit", "value", upperLimit)
	rangeWidth := int((upperLimit - lowerLimit) / float64(NumberOfRanges))
	numberOfRanges := int((upperLimit - lowerLimit) / 5000)
	slog.Debug("numberOfRanges", "value", numberOfRanges)

	var bars []domain.BarData
	for i := range NumberOfRanges {
		ll := int(lowerLimit) + rangeWidth*i
		ul := int(lowerLimit) + rangeWidth*(i+1) - 1
		barName := fmt.Sprintf("%d-%d", ll, ul)
		var itemsCount int
		for _, price := range prices {
			if price >= ll && price <= ul {
				itemsCount++
			}
		}
		bar := domain.BarData{
			Name:       barName,
			LowerLimit: ll,
			UpperLimit: ul,
			Items:      itemsCount,
		}
		bars = append(bars, bar)
	}
	slog.Debug("domain bar data", "values", bars)

	return bars
}
