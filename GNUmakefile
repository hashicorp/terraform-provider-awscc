TEST?=./...
TEST_COUNT?=1
PKG_NAME=internal/aws/...
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: all build default docs golangci-lint lint plural-data-sources resources schemas singular-data-sources test testacc tools

all: schemas resources singular-data-sources plural-data-sources build

build:
	go install

plural-data-sources:
	rm -f internal/*/*/*_plural_data_source_gen.go
	rm -f internal/*/*/*_plural_data_source_gen_test.go
	go generate internal/provider/plural_data_sources.go

singular-data-sources:
	rm -f internal/*/*/*_singular_data_source_gen.go
	rm -f internal/*/*/*_singular_data_source_gen_test.go
	go generate internal/provider/singular_data_sources.go

resources:
	rm -f internal/*/*/*_resource_gen.go
	rm -f internal/*/*/*_resource_gen_test.go
	go generate internal/provider/resources.go

schemas:
	go generate internal/provider/schemas.go

test:
	go test $(TEST) $(TESTARGS) -timeout=5m

# make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_basic'
testacc:
	TF_ACC=1 go test ./$(PKG_NAME) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

lint: golangci-lint importlint

golangci-lint:
	@golangci-lint run ./internal/...

importlint:
	@impi --local . --scheme stdThirdPartyLocal --ignore-generated=true ./...

tools:
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd tools && go install github.com/pavius/impi/cmd/impi
	cd tools && go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

docs:
	go run internal/provider/generators/import-examples/main.go
	@tfplugindocs generate
