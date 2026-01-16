.PHONY: build install clean test release-dry-run

BINARY := kjbc-mcp
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

build:
	go build -ldflags "-s -w" -o $(BINARY) ./cmd/kjbc-mcp

install:
	go install ./cmd/kjbc-mcp

clean:
	rm -f $(BINARY)

test:
	go test -v ./...

# Test release locally without publishing
release-dry-run:
	goreleaser release --snapshot --clean
