TEST?=./...
TEST_COUNT?=1
PKG_NAME=internal/provider
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: all build resources schemas test testacc

all: schemas resources build

build:
	go install

resources:
	rm -f internal/*/*/*_gen.go
	go generate internal/provider/resources.go

schemas:
	go generate internal/provider/schemas.go

test:
	go test $(TEST) $(TESTARGS) -timeout=5m -parallel=4

# make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccLogGroup_'
testacc:
	TF_ACC=1 go test ./$(PKG_NAME) -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)