.PHONY: help buf-gen

.DEFAULT_GOAL := help

# Colors for terminal output
CYAN   := \033[0;36m
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RESET  := \033[0m

GOBIN := $(shell go env GOPATH)/bin

help: ## Show this help message
	@printf "$(CYAN)Usage:$(RESET) make [target]\n\n"
	@printf "$(CYAN)Available Targets:$(RESET)\n"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)


buf-gen: remove-stale-buf-gen-files buf-generate ## Clean and generate protobuf/gRPC files

buf-generate:
	@printf "$(CYAN)==> Generating protobuf and gRPC code with buf...$(RESET)\n"
	PATH="$(GOBIN):$$PATH" buf generate
	@printf "$(GREEN)✓ Buf generation completed.$(RESET)\n"

remove-stale-buf-gen-files:
	@printf "$(YELLOW)==> Cleaning stale buf generated files...$(RESET)\n"
	@find ./pkg/gen -mindepth 1 -delete 2>/dev/null || true
	@printf "$(GREEN)✓ Stale buf files removed.$(RESET)\n"