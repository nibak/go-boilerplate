FROM golang:1.24.0-alpine3.21 as builder

RUN mkdir /app

RUN apt update && apt-y upgrade

RUN go mod tidy
RUN go build -o /app/server cmd/server/main.go

COPY --from=builder /app/server /app/server
COPY --from=builder /migrations /app/migrations
WORKDIR /app
ENTRYPOINT [ "/app/server" ]