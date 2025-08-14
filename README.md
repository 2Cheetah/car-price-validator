# Car Price Validator

This project creates a visual representation of price distribution for a selected car with various filters. The intention is to help car buyers and car sellers evaluate if the price is adequate.

## Features
- Scrapes car price data from sources (implemented in `internal/scraper`)
- Applies filters for car selection (model, year, mileage)
- Computes price distributions and statistics (`internal/statistics`)
- Generates visualizations of price distribution (`internal/vizualiser`)
- Exposes an HTTP server and handlers for requests (`internal/server`)

## Project layout
- `cmd/` — application entrypoint
- `internal/domain` — domain models
- `internal/scraper` — scraping and data ingestion
- `internal/statistics` — distribution and analysis logic
- `internal/visualiser` — visualisation code
- `internal/server` — HTTP server, handlers and middlewares

## Usage
Build (requires Go 1.24+):

```bash
make build
```

Run:

```bash
make run
```
