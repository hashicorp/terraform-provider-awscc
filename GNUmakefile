TEST?=./...
TEST_COUNT?=1
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: build gen schemas test testacc

build:
	go install

gen:
	rm -f internal/provider/*_gen.go
	go generate ./...

test:
	go test $(TEST) $(TESTARGS) -timeout=5m -parallel=4

testacc:
	TF_ACC=1 go test ./internal/provider -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

schemas:
	go generate internal/provider/schemas.go
