# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
LINT_IMAGE := golangci/golangci-lint:v1.54.1

# Main package path
MAIN_PKG = ./cmd

# Output directory
OUT_DIR = bin

# Build target
BUILD_TARGET = $(OUT_DIR)/cricket-fever

# Default target
default: test

# Run the application using the latest built Docker image
run:
	@DOCKER_IMAGE_NAME=$$(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'app-sov-common-alpine:latest'); \
	if [ -z "$$DOCKER_IMAGE_NAME" ]; then \
		echo "Docker image not found. Please build the image using 'make build-docker'."; \
	else \
		echo "Running the application using Docker image: $$DOCKER_IMAGE_NAME"; \
		(docker run --rm -p 8980:8980 $$DOCKER_IMAGE_NAME || true); \
	fi

# Build the Docker image
build:
	@echo "Building Docker image..."
	@cd kubernetes/ && ./app-common-container-build-for-release.sh
	@cd ../../

# Run the golangci-lint command inside a Docker container
lint:
	@echo "Running golangci-lint..."
	docker run --rm -v $(PWD):/app -w /app $(LINT_IMAGE) golangci-lint run -v

# Run tests with code coverage and generate coverage report
test:
	@echo "Running tests with coverage..."
	@$(GOTEST) -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -func=coverage.out
	@$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean up
clean:
	@echo "Cleaning up..."
	@echo "Removing build dir: [$(OUT_DIR)]"
	@rm -rf $(OUT_DIR)
	@echo "Remove coverage reports"
	@rm -f coverage.out coverage.html sov.json
	@echo "Removing kubernetes/buildArtifacts"
	@rm -rf kubernetes/buildArtifacts

# Install project dependencies
dependencies:
	@echo "Downloading Go module dependencies..."
	@$(GOCMD) mod download


generate:
	sqlboiler psql

.PHONY: default build run build-docker test clean deps