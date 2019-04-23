RELEASE?=0.1.0

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

proto_gen:
	protoc -I. --go_out=plugins=grpc:. \
		proto/quiz.proto

build_all: clean_all build_service build_cli

build_cli:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME} -X main.port=${PORT}" \
		-o bin/cli/cli \
		./cli

build_service:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME} -X main.port=${PORT}" \
		-o bin/service/service \
		./server

clean_all: clean_cli clean_service

clean_cli:
	@rm -f bin/cli/cli

clean_service:
	@rm -f bin/service/service

vendor: prepare_dep
	dep ensure

HAS_DEP := $(shell command -v dep;)

prepare_dep:
ifndef HAS_DEP
	go get -u -v -d github.com/golang/dep/cmd/dep && \
	go install -v github.com/golang/dep/cmd/dep
endif