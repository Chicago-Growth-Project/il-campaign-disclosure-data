GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

SUPPORTED_OS= linux darwin windows
SUPPORTED_ARCH=amd64 arm64

COMMANDS := $(notdir $(patsubst %/,%,$(wildcard cmd/*/)))

.PHONY: all build-all build-local test fmt clean $(COMMANDS)

all: fmt test build-all

build-all: build-local

# Build all commands locally
build-local: $(COMMANDS)

$(COMMANDS):
	@mkdir -p dist
	$(GOBUILD) -o dist/$@ ./cmd/$@

fmt:
	gofmt -w .

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf dist/*
