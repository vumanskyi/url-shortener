FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o url-shortener ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/url-shortener .

ENV APP_PORT=8080

EXPOSE ${APP_PORT}

CMD ["./url-shortener"]
