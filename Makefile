# Makefile for Screenshot-CLI

.PHONY: build build-release clean test install help

# Variables
VERSION ?= 1.0.0
BINARY_NAME = sshot
MAIN_PATH = ./cmd/main.go

# Default target
all: build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME)..."
	go build -ldflags "-s -w" -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: bin/$(BINARY_NAME)"

# Build release version for Windows
build-windows:
	@echo "Building Windows release..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
		-ldflags "-s -w -X main.version=$(VERSION)" \
		-o release/$(BINARY_NAME).exe $(MAIN_PATH)
	@echo "Windows build complete: release/$(BINARY_NAME).exe"

# Build release version for macOS
build-macos:
	@echo "Building macOS release..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-s -w -X main.version=$(VERSION)" \
		-o release/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	@echo "macOS build complete: release/$(BINARY_NAME)-darwin-amd64"

# Build release version for Linux
build-linux:
	@echo "Building Linux release..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags "-s -w -X main.version=$(VERSION)" \
		-o release/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	@echo "Linux build complete: release/$(BINARY_NAME)-linux-amd64"

# Build all platforms
build-all: build-windows build-macos build-linux
	@echo "All platform builds complete"

# Install globally
install:
	@echo "Installing $(BINARY_NAME)..."
	go install $(MAIN_PATH)
	@echo "Installation complete"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -cover ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/ release/
	@echo "Clean complete"

# Create directories
setup:
	@echo "Creating directories..."
	mkdir -p bin release
	@echo "Setup complete"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	@echo "Dependencies downloaded"

# Update dependencies
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy
	@echo "Dependencies updated"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build for current platform"
	@echo "  build-windows  - Build Windows release"
	@echo "  build-macos    - Build macOS release"
	@echo "  build-linux    - Build Linux release"
	@echo "  build-all      - Build for all platforms"
	@echo "  install        - Install globally"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  clean          - Clean build artifacts"
	@echo "  setup          - Create build directories"
	@echo "  deps           - Download dependencies"
	@echo "  deps-update    - Update dependencies"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  help           - Show this help" 