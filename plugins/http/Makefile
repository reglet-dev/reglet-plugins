.PHONY: all build clean test help

PLUGIN_NAME = http
WASM_FILE = $(PLUGIN_NAME).wasm

all: build ## Default target: build

build: ## Build plugin to WASM
	@echo "Building $(PLUGIN_NAME) plugin to WASM..."
	GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o $(WASM_FILE) .
	@echo "Built: $(WASM_FILE)"
	@ls -lh $(WASM_FILE)

clean: ## Remove build artifacts
	@echo "Cleaning..."
	rm -f $(WASM_FILE)

test: ## Run plugin tests
	@echo "Running tests..."
	go test -v ./...

help: ## Display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
