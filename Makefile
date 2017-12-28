GROUP=github.com/dmtrk
# Go command parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
# Go env parameters
export GOPATH:=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
BINARY_NAME:=$(shell basename $(GOPATH))
PACKAGE_NAME:=$(value GROUP)/$(value BINARY_NAME)

all: clean build

clean:
	$(GOCLEAN) && rm -rf build out

build:
	$(GOBUILD) -o build/$(BINARY_NAME) $(PACKAGE_NAME)

test:
	$(GOTEST) -v ./...
