TEST                ?= ./...
TEST_COUNT          ?= 1
PKG_NAME            ?= internal/aws/...
ACCTEST_TIMEOUT     ?= 180m
ACCTEST_PARALLELISM ?= 20
GO_VER              ?= go

default: build

.PHONY: all build default docs docs-all docs-import golangci-lint help lint plural-data-sources resources schemas singular-data-sources test testacc tools

all: schemas resources singular-data-sources plural-data-sources build docs-all ## Generate all schemas, resources, data sources, documentation, and build the provider

help: ## Display this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-21s\033[0m %s\n", $$1, $$2}'

build: prereq-go ## Build the provider
	$(GO_VER) install

plural-data-sources: prereq-go ## Generate plural data sources
	rm -f internal/*/*/*_plural_data_source_gen.go
	rm -f internal/*/*/*_plural_data_source_gen_test.go
	$(GO_VER) generate internal/provider/plural_data_sources.go
	goimports -w internal/*/*/*_plural_data_source_gen.go
	goimports -w internal/*/*/*_plural_data_source_gen_test.go

singular-data-sources: prereq-go ## Generate singular data sources
	rm -f internal/*/*/*_singular_data_source_gen.go
	rm -f internal/*/*/*_singular_data_source_gen_test.go
	$(GO_VER) generate internal/provider/singular_data_sources.go
	goimports -w internal/*/*/*_singular_data_source_gen.go
	goimports -w internal/*/*/*_singular_data_source_gen_test.go

resources: prereq-go ## Generate resources
	rm -f internal/*/*/*_resource_gen.go
	rm -f internal/*/*/*_resource_gen_test.go
	$(GO_VER) generate internal/provider/resources.go
	goimports -w internal/*/*/*_resource_gen.go
	goimports -w internal/*/*/*_resource_gen_test.go

schemas: prereq-go ## Generate schemas
	$(GO_VER) generate internal/provider/schemas.go

test: prereq-go ## Run unit tests
	$(GO_VER) test $(TEST) $(TESTARGS) -timeout=5m

# make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_basic'
testacc: prereq-go ## Run acceptance tests
	TF_ACC=1 $(GO_VER) test ./$(PKG_NAME) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

lint: golangci-lint importlint ## Run all linters

golangci-lint: ## Run golangci-lint
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./internal/...

importlint: ## Run importlint
	@echo "==> Checking source code with importlint..."
	@impi --local . --scheme stdThirdPartyLocal --ignore-generated=true ./...

tools: prereq-go ## Install tools
	cd tools && $(GO_VER) install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd tools && $(GO_VER) install github.com/pavius/impi/cmd/impi
	cd tools && $(GO_VER) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	cd tools && $(GO_VER) install golang.org/x/tools/cmd/goimports@latest

docs-all: docs-import docs-fmt docs ## Generate all documentation

docs: prereq-go ## Generate documentation
	rm -f docs/data-sources/*.md
	rm -f docs/resources/*.md
	@tfplugindocs generate

docs-fmt: prereq-go
	cd examples/resources/ && terraform fmt -recursive

docs-import: prereq-go ## Generate import documentation
	$(GO_VER) run internal/provider/generators/import-examples/main.go -file=internal/provider/import_examples_gen.json

prereq-go: # If $(GO_VER) is not installed, install it
	@if ! type "$(GO_VER)" > /dev/null 2>&1 ; then \
		echo "make: $(GO_VER) not found" ; \
		echo "make: installing $(GO_VER)..." ; \
		echo "make: if you get an error, see https://go.dev/doc/manage-install to locally install various Go versions" ; \
		go install golang.org/dl/$(GO_VER)@latest ; \
		$(GO_VER) download ; \
		echo "make: $(GO_VER) ready" ; \
	fi
