RELEASE?=0.1.0
PORT?=":50051"
URI?="localhost"

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

help: ## This help dialog.
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
		IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "%-30s %s\n" $$help_command $$help_info ; \
	done

proto_gen: ## Generates *pb.go files in the ./proto directory
	protoc -I. --go_out=plugins=grpc:. \
		proto/quiz.proto

build_all: ## Deletes binaries from ./bin/... and creates new ones
build_all: clean_all build_service build_cli

build_cli: ## Builds CLI client and puts it into ./bin/cli directory + copies conf.yaml to the same destination
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}" \
		-o bin/cli/quiz \
		./cli
		cp ./cli/conf.yaml ./bin/cli

build_service: ## Builds gRCP server and puts it into ./bin/service directory + copies cquestions.json to the same destination
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME} -X main.port=${PORT}" \
		-o bin/service/service \
		./server
		cp ./server/questions.json ./bin/service

clean_all: ## Deletes all built binaries and config files
clean_all: clean_cli clean_service

clean_cli: ## Deletes binaries and config files from bin/cli/
	@rm -f bin/cli/quiz
	@rm -f bin/cli/conf.yaml

clean_service: ## Deletes binaries and config files from bin/service/
	@rm -f bin/service/service
	@rm -f bin/service/questions.json

vendor: ## Runs prepare_dep and dep ensure
vendor: prepare_dep
	dep ensure

HAS_DEP := $(shell command -v dep;)

prepare_dep: ## Gets the need to use a dependency management utility
ifndef HAS_DEP
	go get -u -v -d github.com/golang/dep/cmd/dep && \
	go install -v github.com/golang/dep/cmd/dep
endif