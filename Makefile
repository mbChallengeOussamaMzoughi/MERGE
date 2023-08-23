# Variables
BINARY_NAME=mergeapp

all: build

build:
	@echo "Building the application..."
	go build -o $(BINARY_NAME)

run: build
	@echo "Running the application..."
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

test:
	@echo "Running tests..."
	go test -v

.PHONY: all build run clean test
