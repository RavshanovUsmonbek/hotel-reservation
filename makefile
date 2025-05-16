.PHONY: build run test clean

# Build the application
build:
	go build -o bin/api ./main.go

# Run the application
run: build
	./bin/api

# Run tests with verbose output
test:
	go test -v ./...

# Clean up build artifacts
clean:
	rm -rf bin/