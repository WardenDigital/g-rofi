BINARY_NAME=g-rofi
SOURCE_DIR=./
BUILD_DIR=bin
INSTALL_DIR=/usr/local/bin

# Go settings
GO=go

# Version and build info
VERSION=$(shell git describe --tags --always 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

.PHONY: all build clean install uninstall test fmt vet deps help

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 $(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)*.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Install the binary to system path
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@chmod +x $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Installation complete"

# Uninstall the binary from system path
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)..."
	@rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstall complete"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"
