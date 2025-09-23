FROM golang:1.25-alpine AS builder

WORKDIR /car-price-validator

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum .

RUN go mod download

COPY cmd/ ./cmd/

COPY internal/ ./internal/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o car-price-validator ./cmd

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder car-price-validator /

ENTRYPOINT ["/car-price-validator"]

