FROM golang:1.20.2-alpine3.17 as build

RUN apk add --no-cache git

WORKDIR /app
COPY . .

RUN go get ./...
RUN go get github.com/swaggo/swag/gen
RUN go get github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag
RUN swag init -g /internal/factory/router/routes.go --pd
RUN GOARCH=amd64 GOOS=linux go build -o ./bin/server ./cmd/server/main.go

FROM alpine:latest

RUN apk update
RUN apk upgrade
RUN apk add --no-cache bash
RUN apk add --upgrade --no-cache coreutils

WORKDIR /

COPY --from=build /app/bin/server /bin/server

ENTRYPOINT ["/bin/server"]
