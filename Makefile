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

# Install the application
install:
	$(GOBUILD) -o $(BINARY_NAME) github.com/theghostmac/gopher/cmd/gopher
	#mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	mv gopher /Users/macbobbychibuzor/go/bin/gopher

# Uninstall the application
uninstall:
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

.PHONY: build clean test deps run

