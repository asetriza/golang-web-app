.PHONY: duild run

include .env
export

build:
	env GOOS=$(PLATFORM) go build -ldflags="-s -w" -o bin/main cmd/app/main.go

dbup:
	docker run --name=golang-web-app-db -e POSTGRES_PASSWORD=$(POSTGRESQL_PASSWORD) -p $(POSTGRESQL_PORT):$(POSTGRESQL_PORT) -d --rm postgres:13.3-alpine

mup:
	migrate -path ./schema -database '$(POSTGRESQL_CONNECTION)' up

mdown:
	migrate -path ./schema -database '$(POSTGRESQL_CONNECTION)' down

dbuild:
	docker build -t golang-web-app .

drun:
	docker run --name=golang-web-app -p $(PORT):$(PORT) golang-web-app

run: build
	bin/./main

dup:
	docker compose -f docker-compose.yaml up

ddown:
	docker compose stop

di:
	docker images

dp:
	docker ps

lint:
	golangci-lint run

mockrepository:
	~/go/bin/mockgen --destination internal/repository/mock/repository.go github.com/asetriza/golang-web-app/internal/repository Authorization,Todo

mockservice:
	~/go/bin/mockgen --destination internal/service/mock/service.go github.com/asetriza/golang-web-app/internal/service Authorization,Todo

test:
	go test ./... -coverprofile=cover.txt

coverhtml:
	go tool cover -html=cover.txt

coverfunc:
	go tool cover -func=cover.txt

deploy: build lint dockerup
