# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -o $(BINARY_NAME) -v
BINARY_NAME=esctl
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) get

all: build

build:
	@echo "Building with no extensions or optional files"
	$(GOBUILD) 
	@echo "Done!"
build-opt: 
	@echo "Building with all optional files..."
	$(GOBUILD) -tags optional
	@echo "Done!"
build-ext:
	@echo "Building with extensions..."
	$(GOBUILD) -tags ext
	@echo "Done!"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
