# Makefile for termhyo

.PHONY: test clean build fmt vet revive staticcheck golangci-lint check examples help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Test parameters
TEST_FLAGS=-v -race
COVERAGE_OUT=coverage.out

all: check test

# Run all tests
test:
	$(GOTEST) $(TEST_FLAGS) ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -coverprofile=$(COVERAGE_OUT) ./...
	$(GOCMD) tool cover -html=$(COVERAGE_OUT)

# Update golden files
test-update:
	$(GOTEST) -update ./...

# Format code
fmt:
	$(GOFMT) ./...

# Check formatting
check-fmt:
	@if [ -n "$$(gofmt -s -l .)" ]; then \
		echo "Code is not formatted:"; \
		gofmt -s -l .; \
		exit 1; \
	fi

# Vet code
vet:
	$(GOVET) ./...

# Install revive and run it (golint replacement)
revive:
	@if ! command -v revive >/dev/null 2>&1; then \
		echo "Installing revive..."; \
		$(GOCMD) install github.com/mgechev/revive@latest; \
	fi
	revive -set_exit_status ./...

# Install staticcheck and run it (most recommended)
staticcheck:
	@if ! command -v staticcheck >/dev/null 2>&1; then \
		echo "Installing staticcheck..."; \
		$(GOCMD) install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi
	staticcheck ./...

# Install golangci-lint and run it (comprehensive linter suite)
golangci-lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

# Check mod tidy
check-mod:
	$(GOMOD) tidy
	@if ! git diff --quiet go.mod go.sum; then \
		echo "go.mod or go.sum needs to be updated"; \
		git diff go.mod go.sum; \
		exit 1; \
	fi

# Run all checks
check: check-fmt vet staticcheck check-mod

# Build examples
examples:
	@echo "Building examples..."
	@for dir in examples/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			echo "Building $$dir"; \
			$(GOBUILD) -o "$${dir%/}" "$$dir"; \
		fi \
	done

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(COVERAGE_OUT)
	@for dir in examples/*/; do \
		if [ -f "$${dir%/}" ]; then \
			echo "Removing $${dir%/}"; \
			rm -f "$${dir%/}"; \
		fi \
	done

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) verify

# Release preparation
release-check: clean check test
	@echo "Release checks passed!"

# Help
help:
	@echo "Available targets:"
	@echo "  test          - Run all tests"
	@echo "  test-coverage - Run tests with coverage report" 
	@echo "  test-update   - Update golden test files"
	@echo "  fmt           - Format code"
	@echo "  check-fmt     - Check code formatting"
	@echo "  vet           - Run go vet"
	@echo "  revive        - Run revive (golint replacement)"
	@echo "  staticcheck   - Run staticcheck (most recommended)"
	@echo "  golangci-lint - Run golangci-lint (comprehensive suite)"
	@echo "  check-mod     - Check go.mod is tidy"
	@echo "  check         - Run essential checks (fmt, vet, staticcheck, mod)"
	@echo "  examples      - Build example programs"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Download and verify dependencies"
	@echo "  release-check - Run all checks for release"
	@echo "  help          - Show this help"
