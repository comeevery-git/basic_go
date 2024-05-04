.PHONY: proto build run test

proto:
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/user.proto

build:
	go build -o basicgo ./cmd/basicgo

run:
    ./basicgo

test:
    go tool cover -html=coverage.out -o coverage.html