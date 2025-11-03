# build:
# 	@go build -o bin/go-backend-service cmd/main.go

# test:
# 	@go test -v ./...

# run: build
# 	@./bin/go-backend-service
.PHONY: up down logs build

up:
	docker compose up -d --build
logs:
	docker compose logs -f --tail=200
down:
	docker compose down -v
build:
	docker compose build --no-cache
