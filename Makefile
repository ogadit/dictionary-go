# Variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=dictionary

# Targets
all: test build

build:
	$(GOBUILD) -o build/$(BINARY_NAME).exe -v ./cmd/cli

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/cli
	./$(BINARY_NAME)

.PHONY: all build test clean run