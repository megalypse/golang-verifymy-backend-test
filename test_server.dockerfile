FROM golang:1.20.2-alpine3.17 as build

RUN apk add --no-cache git

WORKDIR /app
COPY . .

RUN go get ./...
CMD go test ./...
