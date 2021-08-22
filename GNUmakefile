TEST?=./...
TEST_COUNT?=1
PKG_NAME=internal/aws/...
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: all build default golangci-lint lint resources schemas test testacc tools

all: schemas resources build

build:
	go install

data-sources:
	rm -f internal/*/*/*_gen.go
	rm -f internal/*/*/*_gen_test.go
	go generate internal/provider/plural-data-sources.go
	# TODO: add singular data source generation

resources:
	rm -f internal/*/*/*_gen.go
	rm -f internal/*/*/*_gen_test.go
	go generate internal/provider/resources.go

schemas:
	go generate internal/provider/schemas.go

test:
	go test $(TEST) $(TESTARGS) -timeout=5m

# make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_basic'
testacc:
	TF_ACC=1 go test ./$(PKG_NAME) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

lint: golangci-lint

golangci-lint:
	@golangci-lint run ./internal/...

tools:
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint
