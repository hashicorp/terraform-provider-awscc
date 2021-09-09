TEST?=./...
TEST_COUNT?=1
PKG_NAME=internal/aws/...
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: all build data-sources default golangci-lint lint resources schemas test testacc tools

all: schemas resources build

build:
	go install

data-sources:
	rm -f internal/*/*/*_plural_data_source_gen.go
	rm -f internal/*/*/*_plural_data_source_gen_test.go
	go generate internal/provider/plural_data_sources.go
	# TODO: Generate Singular Data Sources

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
