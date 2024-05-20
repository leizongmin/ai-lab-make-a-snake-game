# Makefile for building and cleaning the Go project

# Build the project
dist:
	@echo "Building the project..."
	@mkdir -p dist
	@go build -o dist/snake-game cmd/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf dist
