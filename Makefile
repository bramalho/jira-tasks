# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=jira_tasks
BINARY_UNIX=$(BINARY_NAME)_unix

all: clean build build-linux
clean:
	$(GOCLEAN)
	rm -rf build/
build:
	$(GOBUILD) -o build/$(BINARY_NAME) -v
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_UNIX) -v
