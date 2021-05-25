.PHONY: duild run

include .env
export

build:
	env GOOS=$(PLATFORM) go build -ldflags="-s -w" -o bin/main cmd/app/main.go

dbup:
	docker run --name=golang-web-app-db -e POSTGRES_PASSWORD=$(POSTGRESQL_PASSWORD) -p $(POSTGRESQL_PORT):$(POSTGRESQL_PORT) -d --rm postgres:13.3-alpine

migrateup:
	migrate -path ./schema -database '$(POSTGRESQL_CONNECTION)' up

migratedown:
	migrate -path ./schema -database '$(POSTGRESQL_CONNECTION)' down

dockerbuild:
	docker build -t golang-web-app .

dockerrun:
	docker run --name=golang-web-app -p $(PORT):$(PORT) golang-web-app

run:
	bin/./main

dockerup:
	docker compose -f docker-compose.yaml up

dockerdown:
	docker compose stop

lint:
	golangci-lint run

deploy: build lint dockerup migrateup
