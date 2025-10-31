-include .env

.PHONY: all build

all: build

build:
	go build -o bin/api cmd/api.go
	go build -o bin/handler cmd/handler.go

start:
	./bin/api

start-handler:
	./bin/handler
