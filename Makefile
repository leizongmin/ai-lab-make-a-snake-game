# Makefile for building, cleaning, and running the Go project

# Build the project
dist:
	@echo "Building the project..."
	@mkdir -p dist
	@go build -o dist/snake-game cmd/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf dist

# Run the game for local development and debugging
run:
	@echo "Running the game..."
	@go run cmd/main.go
