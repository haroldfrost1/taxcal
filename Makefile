.PHONY: build docker-build docker-run clean test

# Build the Go binary
build:
	go build -o taxcal

# Run the binary
run:
	./taxcal

# Build Docker image
docker-build:
	docker build -t taxcal .

# Run with Docker interactively
docker-run:
	docker run -it taxcal

# Build both binary and Docker image
all: build docker-build

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f taxcal
	docker rmi taxcal 2>/dev/null || true

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the Go binary"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker interactively"
	@echo "  all          - Build both binary and Docker image"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  help         - Show this help"
