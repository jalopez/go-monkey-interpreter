BUILD_USER       ?= $(shell whoami)@$(shell hostname)
BUILD_DATE       ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILD_DATE_SHORT ?= $(shell date -u +"%Y%m%d")

GIT_VERSION ?= $(shell git describe --tags --always --dirty)

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: clean
clean:
	rm -rf dist/

.PHONY: build
build: clean
	go build -o dist/monkey ./cmd/monkey

.PHONY: test
test:
	go test -v ./pkg/...

.PHONY: lint
lint:
	golangci-lint run -v --timeout=10m

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check: lint
	@test -z $(shell gofmt -l $(SRC) | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@go vet ./...

.PHONY: run
run:
	go run cmd/monkey/main.go --log-level debug
