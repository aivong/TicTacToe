.PHONY: build test lint clean build-all coverage fmt install help

# Binary name
BINARY_NAME=tictactoe

# Build output directory
BUILD_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOINSTALL=$(GOCMD) install

# Build flags
LDFLAGS=-ldflags="-s -w"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build binary for current platform
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

test: ## Run all tests with coverage
	@echo "Running tests..."
	$(GOTEST) -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo "Tests complete"

lint: ## Run golangci-lint
	@echo "Running linters..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin"; \
	fi

fmt: ## Format code with gofmt
	@echo "Formatting code..."
	$(GOFMT) ./...
	@echo "Format complete"

vet: ## Run go vet
	@echo "Running go vet..."
	$(GOVET) ./...
	@echo "Vet complete"

clean: ## Remove build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.txt coverage.html
	@echo "Clean complete"

install: build ## Build and install binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	$(GOINSTALL) .
	@echo "Install complete"

coverage: test ## Generate HTML coverage report
	@echo "Generating coverage report..."
	$(GOCMD) tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report: coverage.html"

# Cross-platform builds
build-all: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 ## Build for all platforms

build-linux-amd64: ## Build for Linux AMD64
	@echo "Building for Linux AMD64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .

build-linux-arm64: ## Build for Linux ARM64
	@echo "Building for Linux ARM64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .

build-darwin-amd64: ## Build for macOS AMD64
	@echo "Building for macOS AMD64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .

build-darwin-arm64: ## Build for macOS ARM64
	@echo "Building for macOS ARM64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

build-windows-amd64: ## Build for Windows AMD64
	@echo "Building for Windows AMD64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

tidy: ## Tidy go.mod
	$(GOMOD) tidy
