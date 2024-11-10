# Variables
PROJECT_NAME := honeybadgerbft
BUILD_DIR := bin
SRC_DIR := ./cmd/$(PROJECT_NAME)
BINARY := $(BUILD_DIR)/$(PROJECT_NAME)
TEST_PATTERN := ./...

# Commands
GO := go
GOTEST := go test
GOLINT := golint

# Targets
.PHONY: all build run test lint clean

# Build the project
all: build

# Build binary
build:
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BINARY) $(SRC_DIR)

# Run the project
run: build
	@echo "Running $(PROJECT_NAME)..."
	$(BINARY)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) $(TEST_PATTERN)

# Lint code
lint:
	@echo "Linting code..."
	$(GOLINT) ./...

# Clean up generated files
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Help
help:
	@echo "Makefile for $(PROJECT_NAME)"
	@echo "Usage:"
	@echo "  make          - Build the project"
	@echo "  make run      - Run the project"
	@echo "  make test     - Run tests"
	@echo "  make lint     - Run linter"
	@echo "  make clean    - Clean up build artifacts"
