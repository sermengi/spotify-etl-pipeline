SHELL := /usr/bin/env bash

GO ?= go
API_CMD := ./cmd/api
BIN_DIR := bin

.PHONY: build-api lint-go tidy

build-api:
	$(GO) build -o $(BIN_DIR)/api $(API_CMD)

lint-go:
	golangci-lint run ./...

tidy:
	$(GO) mod tidy
