BUILD_USER       ?= $(shell whoami)@$(shell hostname)
BUILD_DATE       ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILD_DATE_SHORT ?= $(shell date -u +"%Y%m%d")

GIT_VERSION ?= $(shell git describe --tags --always --dirty)

.PHONY: clean
clean:
	rm -rf dist/

.PHONY: build
build: clean
	go build -o dist/FIXME ./cmd/FIXME

.PHONY: test
test:
	go test -v ./pkg/...

.PHONY: lint
lint:
	golangci-lint run -v --timeout=10m

.PHONY: run
run:
	go run cmd/FIXME/main.go --log-level debug
