# Makefile for building and managing the project.
#
# For help, run: make help
#
# Usage:
#   make [target]

BINARY_NAME=loadbalancer
CMD_PATH=./cmd/loadbalancer

PING_PONG_BIN_NAME=ping-pong
PING_PONG_CMD_PATH=./cmd/ping-pong

BIN_DIR=./bin

.PHONY: all help build-ci build build-linux build-deploy test test-ci lint vendor clean run run-pong install-deps

all: build

install-deps:
	go mod download

help:
	@echo "Common targets:"
	@echo "  build        Build all binaries"
	@echo "  build-ci     Build for CI (no output binary)"
	@echo "  test         Run tests with race and coverage"
	@echo "  test-ci      Run CI tests with coverage report"
	@echo "  lint         Run golangci-lint"
	@echo "  vendor       Vendor dependencies"
	@echo "  install-deps Download Go module dependencies"
	@echo "  clean        Clean binaries and cache"
	@echo "  run          Run main binary"
	@echo "  run-pong     Run ping-pong binary (use: make run-pong port=8080)"

# CI Build: Check that everything compiles, but don't produce a binary
build-ci: install-deps vendor
	go build -mod=vendor -v ./...

# Deployment Build: Build the main binaries for deployment
build: install-deps vendor
	mkdir -p $(BIN_DIR)
	go build -mod=vendor -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_PATH)
	go build -mod=vendor -o $(BIN_DIR)/$(PING_PONG_BIN_NAME) $(PING_PONG_CMD_PATH)

# Run tests with race detector and coverage
test: lint
	go test -race -cover ./...

test-ci: install-deps vendor
	@echo "Running tests with coverage..."
	go test ./... -mod=vendor -covermode=atomic -coverprofile=coverage.out 1>/dev/null 2>&1 || { echo "Tests failed"; exit 1; }
	go tool cover -func=coverage.out | tee coverage.txt

# Lint using golangci-lint (install if not present)
lint: install-deps vendor
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8; \
	fi
	golangci-lint run

# Vendor dependencies
vendor: install-deps
	go mod vendor

# Remove built files and cache
clean:
	go clean
	rm -f $(BIN_DIR)/$(BINARY_NAME) $(BIN_DIR)/$(PING_PONG_BIN_NAME)
	rm -rf vendor coverage.out coverage.txt bin
	rm -rf $(BIN_DIR)

# Run the binary with default config
run: build
	$(BIN_DIR)/$(BINARY_NAME) -config cmd/loadbalancer/config.yaml

# This target will launch the Pong service for testing.
# Usage:
#   make run-pong port=<port>, where <port> is the port number for the Pong service.
run-pong: build
	@if [ -z "$(port)" ]; then \
		echo "Usage: make run-pong port=8081"; \
	exit 1; \
	fi
	$(BIN_DIR)/$(PING_PONG_BIN_NAME) -port $(port)
