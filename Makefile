# Project settings
BINARY_NAME=loadbalancer
CMD_PATH=./cmd/loadbalancer

.PHONY: all build test lint clean run

all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

# Run tests with race detector and coverage
test:
	go test -race -cover ./...

# Lint using golangci-lint (install if not present)
lint:
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

# Remove built files and cache
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run the binary with default config
run: build
	./$(BINARY_NAME) -config config.yaml

# Cross-compile for Linux AMD64
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux $(CMD_PATH)
