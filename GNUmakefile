TEST?=./...
TEST_COUNT?=1
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: all build resources schemas test testacc

all: schemas resources build

build:
	go install

resources:
	rm -f internal/provider/*_gen.go
	go generate internal/provider/resources.go

schemas:
	go generate internal/provider/schemas.go

test:
	go test $(TEST) $(TESTARGS) -timeout=5m -parallel=4

testacc:
	TF_ACC=1 go test ./internal/provider -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)
