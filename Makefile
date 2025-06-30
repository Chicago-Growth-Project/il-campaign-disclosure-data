GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

SUPPORTED_OS= linux darwin windows
SUPPORTED_ARCH=amd64 arm64
BINARY_NAME=il-data-collector

.PHONY: all build-all build-local linux-amd64 darwin-arm64 windows-amd64 test clean
all: test build-all
build-all: build-local
linux-amd64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME)-linux-amd64 -v .
# Windows build currently does not work as we need a cross-compiler for Windows on Linux.
windows-amd64:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME).exe -v .
# Darwin ARM64 build currently does not work as we need a cross-compiler for ARM64 on macOS.
darwin-arm64:
#	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o dist/$(BINARY_NAME)-macos -v .
build-local:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v .
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f dist/*
