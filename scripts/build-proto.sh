#!/usr/bin/env bash

export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=./internal/delivery/grpc/proto --go_opt=paths=source_relative \
    --go-grpc_out=./internal/delivery/grpc/proto --go-grpc_opt=paths=source_relative \
    ./internal/delivery/grpc/proto/**.proto