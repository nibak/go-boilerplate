FROM golang:1.21.3-alpine3.18 AS builder

WORKDIR /app

RUN apk add --no-cache curl
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]