EXECUTABLE=blackjack
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_arm64
VERSION=$(shell git describe --tags --always --long --dirty)
TAG=$(shell git describe --tags --abbrev=0)

.PHONY: all test release clean

all: test build ## Build and run tests

test: ## Run unit tests
	./scripts/test_unit.sh

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=arm64 go build -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/game/main.go

release: ## Create release
	./scripts/release.sh

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'