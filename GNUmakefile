TEST                ?= ./...
TEST_COUNT          ?= 1
PKG_NAME            ?= internal/aws/...
ACCTEST_TIMEOUT     ?= 180m
ACCTEST_PARALLELISM ?= 20
GO_VER              ?= go

default: build

.PHONY: all build default docs golangci-lint lint plural-data-sources resources schemas singular-data-sources test testacc tools

all: schemas resources singular-data-sources plural-data-sources build docs

build:
	$(GO_VER) install

plural-data-sources:
	rm -f internal/*/*/*_plural_data_source_gen.go
	rm -f internal/*/*/*_plural_data_source_gen_test.go
	$(GO_VER) generate internal/provider/plural_data_sources.go
	goimports -w internal/*/*/*_plural_data_source_gen.go
	goimports -w internal/*/*/*_plural_data_source_gen_test.go

singular-data-sources:
	rm -f internal/*/*/*_singular_data_source_gen.go
	rm -f internal/*/*/*_singular_data_source_gen_test.go
	$(GO_VER) generate internal/provider/singular_data_sources.go
	goimports -w internal/*/*/*_singular_data_source_gen.go
	goimports -w internal/*/*/*_singular_data_source_gen_test.go

resources:
	rm -f internal/*/*/*_resource_gen.go
	rm -f internal/*/*/*_resource_gen_test.go
	$(GO_VER) generate internal/provider/resources.go
	goimports -w internal/*/*/*_resource_gen.go
	goimports -w internal/*/*/*_resource_gen_test.go

schemas:
	$(GO_VER) generate internal/provider/schemas.go

test:
	$(GO_VER) test $(TEST) $(TESTARGS) -timeout=5m

# make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_basic'
testacc:
	TF_ACC=1 $(GO_VER) test ./$(PKG_NAME) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

lint: golangci-lint importlint

golangci-lint:
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./internal/...

importlint:
	@echo "==> Checking source code with importlint..."
	@impi --local . --scheme stdThirdPartyLocal --ignore-generated=true ./...

tools:
	cd tools && $(GO_VER) install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd tools && $(GO_VER) install github.com/pavius/impi/cmd/impi
	cd tools && $(GO_VER) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	cd tools && $(GO_VER) install golang.org/x/tools/cmd/goimports@latest

docs:
	$(GO_VER) run internal/provider/generators/import-examples/main.go
	rm -f docs/data-sources/*.md
	rm -f docs/resources/*.md
	@tfplugindocs generate
