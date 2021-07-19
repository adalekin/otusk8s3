
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_DIRECTORY=./bin
BINARY_NAME=otusk8s3
BINARY_PATH=$(BINARY_DIRECTORY)/$(BINARY_NAME)

all: test build

.PHONY: build

build:
	$(GOBUILD) -o $(BINARY_DIRECTORY) -v ./...

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_DIRECTORY)/*

run:
	$(GOBUILD) -o $(BINARY_DIRECTORY) -v ./...
	$(BINARY_PATH)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_DIRECTORY) -v ./...

build-docker:
	docker build -t adalekin/otusk8s3:latest -f build/docker/Dockerfile .
