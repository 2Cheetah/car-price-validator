.DEFAULT_GOAL := run
.PHONY: run test fmt vet

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

run: vet
	go run cmd/main.go

test: fmt vet
	go test ./...

build: test
	go build -o ./car-price-validator ./cmd/main.go
