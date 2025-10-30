-include .env

.PHONY: all build

all:
	go run cmd/*.go

build:
	go build -o bin/api cmd/api.go

build-handler:
	go build -o bin/handler cmd/handler.go

start:
	./bin/api

start-handler:
	./bin/handler
