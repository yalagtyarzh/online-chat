export

include .env

.PHONY:

build:
	go build -o ./.bin/online-chat ./cmd/main.go

run: build
	./.bin/online-chat

linter-golangci:
	golangci-lint run
