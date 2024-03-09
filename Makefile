EXECUTABLE=frequencify
# we could get the version from git describe
VERSION=0.0.1

# arch specific names
DARWIN=$(EXECUTABLE)_darwin_amd64
LINUX=$(EXECUTABLE)_linux_amd64
WINDOWS=$(EXECUTABLE)_windows_amd64.exe

.PHONY: all test clean

all: test build ## Build and run tests

## Run unit tests in verbose mode
test:
	go test ./... -v

# cross compile for 3 architectures
build: darwin linux windows
	@echo version: $(VERSION)

darwin: $(DARWIN)

linux: $(LINUX)

windows: $(WINDOWS)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  main.go

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

docker: ## Containerise the application
	docker build -t frequencify .

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'