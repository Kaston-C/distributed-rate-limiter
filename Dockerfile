# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o rate-limiter main.go

# Final stage
FROM golang:1.24

LABEL authors="kaston_carr"

WORKDIR /app
COPY --from=builder /app/rate-limiter /app/rate-limiter

EXPOSE 8080

ENTRYPOINT ["/app/rate-limiter"]