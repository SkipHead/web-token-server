FROM golang:latest

COPY ./ /app

WORKDIR /app

RUN go build -o web-token-server ./cmd/main.go
ENRTYPOINT ["/app/web-token-server", "--config-path=/app/config/server.json"]

