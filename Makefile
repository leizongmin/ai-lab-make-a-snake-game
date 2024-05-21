# Makefile for building, cleaning, running the Go project, and executing unit tests

# Build the project
.PHONY: dist
dist:
	@echo "Building the project..."
	@mkdir -p dist
	@go build -ldflags="-s -w" -o dist/snake-game cmd/main.go

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf dist

# Run the game for local development and debugging
.PHONY: run
run:
	@echo "Running the game..."
	@go run cmd/main.go

# Execute unit tests
.PHONY: test
test:
	@echo "Executing unit tests..."
	@go test ./...
