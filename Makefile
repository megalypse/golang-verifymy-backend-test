build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/server ./cmd/server/main.go

IMAGE_ID := ${shell docker images 'verifymy_backend_golang_server' -a -q}
clean:
	docker compose stop server
	docker rmi ${IMAGE_ID} -f

DOC_DEPS_PATH=./internal/domain/models
SWAGGER_ENTRYPOINT=./internal/main/factory/router_factory.go
generate-docs-silent:
	swag init -g ${SWAGGER_ENTRYPOINT} --pd --quiet

run-compose: generate-docs-silent build
	docker compose up -d

run-compose-clean: generate-docs-silent clean build
	docker compose up -d

run-compose-clean-all: generate-docs-silent build
	docker compose down
	docker compose up -d


