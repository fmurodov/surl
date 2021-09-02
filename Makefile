.DEFAULT_GOAL := help

.PHONY: proto
proto: 
	protoc -I api/proto --go-grpc_out=. --go_out=. api/proto/surl.proto

.PHONY: build
build: proto
	go build -o bin/surl ./cmd/server

.PHONY: test
test: build
	go test ./...

.PHONY: lint
lint: build
	golangci-lint run ./...

.PHONY: build-docker
build-docker:
	docker build . -t surl

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  proto			generate proto files"
	@echo "  build			build the application"
	@echo "  build-docker	build docker image"
	@echo "  test			run the tests"
	@echo "  help			display this help"
	@echo ""
