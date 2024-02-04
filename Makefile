.DEFAULT_GOAL := build
.PHONY: test

mq:
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management

build:
	go build -o bin/client cmd/client/*.go
	go build -o bin/server cmd/server/*.go

test:
	@go test -failfast model/ordmap/*.go

test-race:
	@go test -race ./...
