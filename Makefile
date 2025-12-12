SHELL := /usr/bin/env bash

GO ?= go
API_CMD := ./cmd/api
BIN_DIR := bin

DOCKER_IMAGE := spotify-etl-api 

.PHONY: docker_build docker-run compose-up compose-down build-api lint-go tidy

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run --rm -p 8080:8080 $(DOCKER_IMAGE)

compose-up:
	docker compose up --build

compose-down:
	docker compose down

build-api:
	$(GO) build -o $(BIN_DIR)/api $(API_CMD)

lint-go:
	golangci-lint run ./...

tidy:
	$(GO) mod tidy
