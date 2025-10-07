.DEFAULT_GOAL := run
.PHONY: run test fmt vet

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

run: test
	go run cmd/main.go

test: fmt vet
	go test -v ./...

build: test
	go build -o ./car-price-validator ./cmd/main.go

dev-build:
	docker compose -f docker-compose.yaml -f docker-compose.override.yaml build

dev-up:
	docker compose -f docker-compose.yaml -f docker-compose.override.yaml up

dev-down:
	docker compose -f docker-compose.yaml -f docker-compose.override.yaml down

prod-build:
	docker compose -f docker-compose.yaml build

prod-up:
	docker compose -f docker-compose.yaml up

prod-down:
	docker compose -f docker-compose.yaml down
