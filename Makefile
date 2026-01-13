# Reglet Plugins Makefile
# Top-level commands for building, testing, and managing plugins

PLUGINS := $(wildcard plugins/*)

.PHONY: all build-all test-all clean-all lint-all help

all: build-all ## Default target: build all plugins

build-all: ## Build all plugins to WASM
	@echo "Building all plugins..."
	@for dir in $(PLUGINS); do \
		echo "  Building $$dir..."; \
		$(MAKE) -C $$dir build || exit 1; \
	done
	@echo "All plugins built successfully!"

test-all: ## Run tests for all plugins
	@echo "Testing all plugins..."
	@for dir in $(PLUGINS); do \
		echo "  Testing $$dir..."; \
		$(MAKE) -C $$dir test || exit 1; \
	done
	@echo "All tests passed!"

clean-all: ## Remove all build artifacts
	@echo "Cleaning all plugins..."
	@for dir in $(PLUGINS); do \
		$(MAKE) -C $$dir clean || true; \
	done
	@echo "All plugins cleaned."

lint-all: ## Run linter on all plugins
	@echo "Linting all plugins..."
	@for dir in $(PLUGINS); do \
		echo "  Linting $$dir..."; \
		cd $$dir && go vet ./... && cd ../..; \
	done
	@echo "Linting complete."

# Individual plugin shortcuts
.PHONY: command dns file http smtp tcp

command: ## Build the command plugin
	$(MAKE) -C plugins/command build

dns: ## Build the dns plugin
	$(MAKE) -C plugins/dns build

file: ## Build the file plugin
	$(MAKE) -C plugins/file build

http: ## Build the http plugin
	$(MAKE) -C plugins/http build

smtp: ## Build the smtp plugin
	$(MAKE) -C plugins/smtp build

tcp: ## Build the tcp plugin
	$(MAKE) -C plugins/tcp build

help: ## Display this help message
	@echo "Reglet Plugins - Available targets:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
