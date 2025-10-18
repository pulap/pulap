# Makefile for the Pulap monorepo

# Variables
PROJECT_NAME=pulap
SERVICES=authn authz estate admin
BASE_PORTS=8081 8082 8083 8080
PKG_LIBS=auth core fake

# Go related commands
GOFUMPT=gofumpt
GCI=gci
GOLANGCI_LINT=golangci-lint
GO_TEST=go test
GO_VET=go vet
GO_VULNCHECK=govulncheck

# Phony targets
.PHONY: all build run test test-v test-short coverage coverage-html coverage-func coverage-profile coverage-check coverage-100 clean fmt lint vet check ci run-all stop-all help build-all test-all lint-all

all: build-all

help:
	@echo "Available targets:"
	@echo "  build-all    - Build all services"
	@echo "  run-all      - Kill ports, build, and start all services"
	@echo "  stop-all     - Stop all running services"
	@echo "  test         - Run tests for all components"
	@echo "  test-v       - Run tests with verbose output"
	@echo "  test-short   - Run tests in short mode"
	@echo "  test-all     - Run tests for all services and pkg libs"
	@echo "  coverage     - Run tests with coverage report"
	@echo "  coverage-html - Generate HTML coverage report"
	@echo "  coverage-func - Show function-level coverage"
	@echo "  coverage-check - Check coverage meets 80% threshold"
	@echo "  coverage-100 - Check 100% test coverage"
	@echo "  lint         - Run golangci-lint on all code"
	@echo "  lint-all     - Run lint on services and pkg libs"
	@echo "  fmt          - Format all Go code"
	@echo "  vet          - Run go vet on all code"
	@echo "  clean        - Clean all generated files and binaries"
	@echo "  check        - Run all quality checks"
	@echo "  ci           - Run CI pipeline with strict checks"
	@echo ""
	@echo "Individual service targets (replace <service> with authn/authz/estate/admin):"
	@echo "  build-<service>  - Build specific service"
	@echo "  test-<service>   - Test specific service"
	@echo "  lint-<service>   - Lint specific service"
	@echo "  run-<service>    - Run specific service"

# Build all services
build-all:
	@echo "🏗️  Building all services..."
	@for service in $(SERVICES); do \
		echo "   📦 Building $$service..."; \
		cd services/$$service && go build -o $$service . || exit 1; \
		cd ../..; \
	done
	@echo "✅ All services built successfully"

# Build individual services
build-authn:
	@echo "📦 Building authn service..."
	@cd services/authn && go build -o authn .

build-authz:
	@echo "📦 Building authz service..."
	@cd services/authz && go build -o authz .

build-estate:
	@echo "📦 Building estate service..."
	@cd services/estate && go build -o estate .

build-admin:
	@echo "📦 Building admin service..."
	@cd services/admin && go build -o admin .

# Test all components
test:
	@echo "🧪 Running tests for all components..."
	@$(GO_TEST) ./...

test-v:
	@echo "🧪 Running tests with verbose output..."
	@$(GO_TEST) -v ./...

test-short:
	@echo "🧪 Running tests in short mode..."
	@$(GO_TEST) -short ./...

# Test all services and pkg libs individually
test-all:
	@echo "🧪 Running tests for all services and pkg libs..."
	@for lib in $(PKG_LIBS); do \
		echo "   🧪 Testing pkg/lib/$$lib..."; \
		cd pkg/lib/$$lib && go test ./... || exit 1; \
		cd ../../..; \
	done
	@for service in $(SERVICES); do \
		echo "   🧪 Testing $$service service..."; \
		cd services/$$service && go test ./... || exit 1; \
		cd ../..; \
	done

# Test individual services
test-authn:
	@cd services/authn && go test ./...

test-authz:
	@cd services/authz && go test ./...

test-estate:
	@cd services/estate && go test ./...

test-admin:
	@cd services/admin && go test ./...

# Coverage targets
coverage:
	@echo "📊 Running tests with coverage..."
	@$(GO_TEST) -cover ./...

coverage-profile:
	@echo "📊 Generating coverage profile..."
	@$(GO_TEST) -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | tail -1

coverage-html: coverage-profile
	@echo "📊 Generating HTML coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

coverage-func: coverage-profile
	@echo "📊 Function-level coverage:"
	@go tool cover -func=coverage.out

coverage-check: coverage-profile
	@COVERAGE=$$(go tool cover -func=coverage.out | tail -1 | awk '{print $$3}' | sed 's/%//'); \
	echo "📊 Coverage: $$COVERAGE%"; \
	if [ $$(echo "$$COVERAGE < 80" | bc -l) -eq 1 ]; then \
		echo "❌ Coverage $$COVERAGE% is below 80% threshold"; \
		exit 1; \
	else \
		echo "✅ Coverage $$COVERAGE% meets the 80% threshold"; \
	fi

coverage-100: coverage-profile
	@COVERAGE=$$(go tool cover -func=coverage.out | tail -1 | awk '{print $$3}' | sed 's/%//'); \
	echo "📊 Coverage: $$COVERAGE%"; \
	if [ "$$COVERAGE" != "100.0" ]; then \
		echo "❌ Coverage $$COVERAGE% is not 100%"; \
		go tool cover -func=coverage.out | grep -v "100.0%"; \
		exit 1; \
	else \
		echo "🎉 Perfect! 100% test coverage!"; \
	fi

# Code quality targets
fmt:
	@echo "🎨 Formatting Go code..."
	@$(GOFUMPT) -l -w .
	@$(GCI) -w .
	@echo "✅ Go code formatted"

vet:
	@echo "🔍 Running go vet..."
	@$(GO_VET) ./...

lint:
	@echo "🔍 Running golangci-lint..."
	@$(GOLANGCI_LINT) run
	@echo "✅ golangci-lint finished"

# Lint all services and pkg libs
lint-all:
	@echo "🔍 Running golangci-lint on all components..."
	@for lib in $(PKG_LIBS); do \
		echo "   🔍 Linting pkg/lib/$$lib..."; \
		cd pkg/lib/$$lib && golangci-lint run || exit 1; \
		cd ../../..; \
	done
	@for service in $(SERVICES); do \
		echo "   🔍 Linting $$service service..."; \
		cd services/$$service && golangci-lint run || exit 1; \
		cd ../..; \
	done

# Individual service lint
lint-authn:
	@cd services/authn && golangci-lint run

lint-authz:
	@cd services/authz && golangci-lint run

lint-estate:
	@cd services/estate && golangci-lint run

lint-admin:
	@cd services/admin && golangci-lint run

# Quality checks
check: fmt vet test coverage-check lint
	@echo "✅ All quality checks passed!"

ci: fmt vet test coverage-100 lint
	@echo "🚀 CI pipeline passed!"

# Service management
run-all:
	@echo "🚀 Starting full Pulap environment..."
	@$(MAKE) stop-all
	@$(MAKE) build-all
	@echo "🚀 Starting services..."
	@echo "   📦 Starting AuthN on :8081..."
	@cd services/authn && nohup ./authn > authn.log 2>&1 & echo $$! > authn.pid; sleep 2
	@echo "   📦 Starting AuthZ on :8082..."
	@cd services/authz && nohup ./authz > authz.log 2>&1 & echo $$! > authz.pid; sleep 2
	@echo "   📦 Starting Estate on :8083..."
	@cd services/estate && nohup ./estate > estate.log 2>&1 & echo $$! > estate.pid; sleep 2
	@echo "   📦 Starting Admin on :8080..."
	@cd services/admin && nohup ./admin > admin.log 2>&1 & echo $$! > admin.pid; sleep 2
	@echo ""
	@echo "🎉 All Pulap services started!"
	@echo "📡 Services running:"
	@echo "   • Admin:  http://localhost:8080 (web interface)"
	@echo "   • AuthN:  http://localhost:8081 (authentication)"
	@echo "   • AuthZ:  http://localhost:8082 (authorization)"
	@echo "   • Estate: http://localhost:8083 (real estate)"
	@echo ""
	@echo "🛑 To stop all services: make stop-all"

# Individual service runners
run-authn: build-authn
	@cd services/authn && ./authn

run-authz: build-authz
	@cd services/authz && ./authz

run-estate: build-estate
	@cd services/estate && ./estate

run-admin: build-admin
	@cd services/admin && ./admin

stop-all:
	@echo "🛑 Stopping all Pulap services..."
	@for port in 8080 8081 8082 8083 8084 8085; do \
		if lsof -ti:$$port >/dev/null 2>&1; then \
			echo "🛑 Stopping process on port $$port"; \
			lsof -ti:$$port | xargs -r kill -9 || true; \
		fi; \
	done
	@for service in $(SERVICES); do \
		pid_file="services/$$service/$$service.pid"; \
		if [ -f "$$pid_file" ]; then \
			pid=$$(cat "$$pid_file"); \
			if ps -p "$$pid" >/dev/null 2>&1; then \
				echo "🛑 Stopping $$service (PID: $$pid)"; \
				kill -9 "$$pid" 2>/dev/null || true; \
			fi; \
			rm -f "$$pid_file"; \
		fi; \
	done
	@echo "✅ All Pulap services stopped"

# Clean targets
clean:
	@echo "🧹 Cleaning up..."
	@for service in $(SERVICES); do \
		cd services/$$service && rm -f $$service *.log *.pid; \
		cd ../..; \
	done
	@go clean -testcache
	@rm -f coverage.out coverage.html
	@echo "✅ Cleanup complete"

# Security check (if govulncheck is available)
security:
	@echo "🔒 Running security checks..."
	@$(GO_VULNCHECK) ./... || echo "⚠️  govulncheck not available, install with: go install golang.org/x/vuln/cmd/govulncheck@latest"

# Development helpers
dev-deps:
	@echo "📥 Installing development dependencies..."
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/daixiang0/gci@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "✅ Development dependencies installed"

# Tidy all modules
tidy:
	@echo "🧹 Running go mod tidy on all modules..."
	@for lib in $(PKG_LIBS); do \
		echo "   📦 Tidying pkg/lib/$$lib..."; \
		cd pkg/lib/$$lib && go mod tidy; \
		cd ../../..; \
	done
	@for service in $(SERVICES); do \
		echo "   📦 Tidying $$service service..."; \
		cd services/$$service && go mod tidy; \
		cd ../..; \
	done
	@echo "✅ All modules tidied"