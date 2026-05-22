GO ?= go

.PHONY: build test install

build:
	mkdir -p bin
	$(GO) build -o bin/aenv .

test:
	$(GO) test ./...

install:
	$(GO) install .
