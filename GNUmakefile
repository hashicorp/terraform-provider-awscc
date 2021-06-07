TEST?=./...
TEST_COUNT?=1
ACCTEST_TIMEOUT?=180m
ACCTEST_PARALLELISM?=20

default: build

.PHONY: build download-schemas gen test testacc

build:
	go install

gen:
	rm -f internal/provider/*_gen.go
	go generate ./...

test:
	go test $(TEST) $(TESTARGS) -timeout=5m -parallel=4

testacc:
	TF_ACC=1 go test ./internal/provider -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

download-schemas:
	GOPRIVATE=github.com/hashicorp go run tools/schema-downloader/main.go -config internal/service/cloudformation/schemas-sync/all_schemas.hcl
