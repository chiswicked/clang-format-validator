BASE_PATH			:= $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
BASE_DIR			:= $(notdir $(BASE_PATH))

ORG					?= chiswicked
SERVICE				?= $(BASE_DIR)
DOCKER_IMAGE		:= $(ORG)/$(SERVICE)
VERSION				?= $(shell cat $(BASE_PATH)/VERSION 2> /dev/null || echo 0.0.1)

BUILD_PATH			:= $(BASE_PATH)/build

# Local commands

.PHONY: all clean install build test cover cover-clean run

all: clean install test build

clean: cover-clean
	@echo [clean] removing binary and other object files
	@go clean
	@rm -f $(BUILD_PATH)/$(SERVICE)

install:
	@echo [install] installing dependencies
	@go get -v -t -d ./...

build: clean
	@echo [build] building binary
	@go build -o $(BUILD_PATH)/$(SERVICE) -a .

test:
	@echo [test] running unit tests
	@go test -v -cover ./...

cover: cover-clean
	@echo [cover] generating test coverage report
	@go test -coverprofile cover.out ./...
	@go tool cover -html=cover.out -o cover.html
	@open cover.html

cover-clean:
	@echo [cover-clean] removing cover.out cover.html
	@rm -f cover.out cover.html

run:
	@echo [run] executing binary
	@$(BUILD_PATH)/$(SERVICE)

# CI commands

.PHONY: ci-docker-build

ci-docker-build:
	@echo "[ci-docker-build] building docker image $(DOCKER_IMAGE)"
	@docker build -t $(DOCKER_IMAGE) .
