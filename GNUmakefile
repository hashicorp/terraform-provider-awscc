TEST                ?= ./...
TEST_COUNT          ?= 1
PKG_NAME            ?= internal/aws/...
ACCTEST_TIMEOUT     ?= 180m
ACCTEST_PARALLELISM ?= 20
GO_VER              ?= go

default: build

.PHONY: all build default docs golangci-lint lint plural-data-sources resources schemas singular-data-sources test testacc tools

all: schemas resources singular-data-sources plural-data-sources build docs-all

build: prereq-go
	$(GO_VER) install

plural-data-sources: prereq-go
	rm -f internal/*/*/*_plural_data_source_gen.go
	rm -f internal/*/*/*_plural_data_source_gen_test.go
	$(GO_VER) generate internal/provider/plural_data_sources.go
	goimports -w internal/*/*/*_plural_data_source_gen.go
	goimports -w internal/*/*/*_plural_data_source_gen_test.go

singular-data-sources: prereq-go
	rm -f internal/*/*/*_singular_data_source_gen.go
	rm -f internal/*/*/*_singular_data_source_gen_test.go
	$(GO_VER) generate internal/provider/singular_data_sources.go
	goimports -w internal/*/*/*_singular_data_source_gen.go
	goimports -w internal/*/*/*_singular_data_source_gen_test.go

resources: prereq-go
	rm -f internal/*/*/*_resource_gen.go
	rm -f internal/*/*/*_resource_gen_test.go
	$(GO_VER) generate internal/provider/resources.go
	goimports -w internal/*/*/*_resource_gen.go
	goimports -w internal/*/*/*_resource_gen_test.go

schemas: prereq-go
	$(GO_VER) generate internal/provider/schemas.go

test: prereq-go
	$(GO_VER) test $(TEST) $(TESTARGS) -timeout=5m

# make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_basic'
testacc: prereq-go
	TF_ACC=1 $(GO_VER) test ./$(PKG_NAME) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

lint: golangci-lint importlint

golangci-lint:
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./internal/...

importlint:
	@echo "==> Checking source code with importlint..."
	@impi --local . --scheme stdThirdPartyLocal --ignore-generated=true ./...

tools: prereq-go
	cd tools && $(GO_VER) install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd tools && $(GO_VER) install github.com/pavius/impi/cmd/impi
	cd tools && $(GO_VER) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	cd tools && $(GO_VER) install golang.org/x/tools/cmd/goimports@latest

docs-all: docs-import docs

docs: prereq-go
	rm -f docs/data-sources/*.md
	rm -f docs/resources/*.md
	@tfplugindocs generate

docs-import: prereq-go
	$(GO_VER) generate internal/provider/import_examples.go

prereq-go: ## If $(GO_VER) is not installed, install it
	@if ! type "$(GO_VER)" > /dev/null 2>&1 ; then \
		echo "make: $(GO_VER) not found" ; \
		echo "make: installing $(GO_VER)..." ; \
		echo "make: if you get an error, see https://go.dev/doc/manage-install to locally install various Go versions" ; \
		go install golang.org/dl/$(GO_VER)@latest ; \
		$(GO_VER) download ; \
		echo "make: $(GO_VER) ready" ; \
	fi
