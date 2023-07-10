# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Binary name
BINARY_NAME = gopher

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/gopher

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run all tests
test:
	$(GOTEST) -v ./...

# Install project dependencies
deps:
	$(GOGET) -v ./...

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/gopher
	./$(BINARY_NAME)

.PHONY: build clean test deps run

