GROUP=github.com/dmtrk

# Go command parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Go env parameters
PROJECT_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
export GOPATH:=$(PROJECT_DIR)/deps:$(PROJECT_DIR)

BINARY_NAME:=$(shell basename $(GOPATH))
PACKAGE_NAME:=$(value GROUP)/$(value BINARY_NAME)

all: clean build

clean:
	$(GOCLEAN) && rm -rf build out

build: deps
	$(GOBUILD) -o build/$(BINARY_NAME) $(PACKAGE_NAME)

deps:
	$(GOGET) github.com/quickfixgo/quickfix

test:
	$(GOTEST) -v ./...
