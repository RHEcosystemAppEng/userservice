## Adapted from https://gist.github.com/thomaspoignant/5b72d579bd5f311904d973652180c705
BINARY_NAME=userservice

.PHONY: all

all: test clean build

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	go build -o out/bin/$(BINARY_NAME) .

clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

## Test:
tests: ## Run the tests of the project
	go test ./test/...

## Docker:
docker-build: ## Use the dockerfile to build the container
	docker build --rm --tag $(BINARY_NAME) .
