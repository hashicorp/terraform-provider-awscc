TEST                ?= ./...
TEST_COUNT          ?= 1
PKG_NAME            ?= internal/aws/...
ACCTEST_TIMEOUT     ?= 180m
ACCTEST_PARALLELISM ?= 20
GO_VER              ?= go

default: build

all: schemas resources singular-data-sources plural-data-sources build docs-all ## Generate all schemas, resources, data sources, documentation, and build the provider

help: ## Display this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-21s\033[0m %s\n", $$1, $$2}'

build: prereq-go ## Build the provider
	$(GO_VER) install

plural-data-sources: prereq-go ## Generate plural data sources
	@echo "==> Counting existing plural data source files..."
	@EXPECTED=$$(find internal -name '*_plural_data_source_gen.go' -o -name '*_plural_data_source_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
	if [ "$$EXPECTED" -eq 0 ]; then EXPECTED=200; fi; \
	echo "==> Expected ~$$EXPECTED files to generate"; \
	rm -f internal/*/*/*_plural_data_source_gen.go; \
	rm -f internal/*/*/*_plural_data_source_gen_test.go; \
	echo "==> Generating plural data sources..."; \
	( \
		while sleep 3; do \
			CURRENT=$$(find internal -name '*_plural_data_source_gen.go' -o -name '*_plural_data_source_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
			if [ "$$CURRENT" -gt 0 ]; then \
				PERCENT=$$((CURRENT * 100 / EXPECTED)); \
				if [ "$$PERCENT" -gt 100 ]; then PERCENT=100; fi; \
				printf "\r==> Progress: $$CURRENT/$$EXPECTED files ($$PERCENT%%) "; \
			fi; \
		done \
	) & PROGRESS_PID=$$!; \
	$(GO_VER) generate internal/provider/plural_data_sources.go > .make-plural-data-sources.log 2>&1; \
	GEN_EXIT=$$?; \
	kill $$PROGRESS_PID 2>/dev/null; wait $$PROGRESS_PID 2>/dev/null; \
	FINAL=$$(find internal -name '*_plural_data_source_gen.go' -o -name '*_plural_data_source_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
	printf "\r==> Generated $$FINAL plural data source files          \n"; \
	if [ $$GEN_EXIT -ne 0 ]; then \
		echo "==> ⚠️  Generation failed - check .make-plural-data-sources.log"; \
		exit $$GEN_EXIT; \
	fi; \
	echo "==> Running goimports on generated files..."; \
	goimports -w internal/*/*/*_plural_data_source_gen.go >> .make-plural-data-sources.log 2>&1; \
	goimports -w internal/*/*/*_plural_data_source_gen_test.go >> .make-plural-data-sources.log 2>&1; \
	if grep -iE '(error|fatal|panic|failed)' .make-plural-data-sources.log > /dev/null; then \
		echo "==> ⚠️  ERRORS DETECTED:"; \
		grep -iE '(error|fatal|panic|failed)' .make-plural-data-sources.log; \
		echo ""; \
		echo "==> Full output in .make-plural-data-sources.log"; \
		exit 1; \
	else \
		echo "==> ✓ Plural data sources generated successfully, no errors detected!"; \
		echo "==> Full output available in .make-plural-data-sources.log"; \
	fi

singular-data-sources: prereq-go ## Generate singular data sources
	@echo "==> Counting existing singular data source files..."
	@EXPECTED=$$(find internal -name '*_singular_data_source_gen.go' -o -name '*_singular_data_source_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
	if [ "$$EXPECTED" -eq 0 ]; then EXPECTED=1000; fi; \
	echo "==> Expected ~$$EXPECTED files to generate"; \
	rm -f internal/*/*/*_singular_data_source_gen.go; \
	rm -f internal/*/*/*_singular_data_source_gen_test.go; \
	echo "==> Generating singular data sources..."; \
	( \
		while sleep 3; do \
			CURRENT=$$(find internal -name '*_singular_data_source_gen.go' -o -name '*_singular_data_source_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
			if [ "$$CURRENT" -gt 0 ]; then \
				PERCENT=$$((CURRENT * 100 / EXPECTED)); \
				if [ "$$PERCENT" -gt 100 ]; then PERCENT=100; fi; \
				printf "\r==> Progress: $$CURRENT/$$EXPECTED files ($$PERCENT%%) "; \
			fi; \
		done \
	) & PROGRESS_PID=$$!; \
	$(GO_VER) generate internal/provider/singular_data_sources.go > .make-singular-data-sources.log 2>&1; \
	GEN_EXIT=$$?; \
	kill $$PROGRESS_PID 2>/dev/null; wait $$PROGRESS_PID 2>/dev/null; \
	FINAL=$$(find internal -name '*_singular_data_source_gen.go' -o -name '*_singular_data_source_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
	printf "\r==> Generated $$FINAL singular data source files          \n"; \
	if [ $$GEN_EXIT -ne 0 ]; then \
		echo "==> ⚠️  Generation failed - check .make-singular-data-sources.log"; \
		exit $$GEN_EXIT; \
	fi; \
	echo "==> Running goimports on generated files..."; \
	goimports -w internal/*/*/*_singular_data_source_gen.go >> .make-singular-data-sources.log 2>&1; \
	goimports -w internal/*/*/*_singular_data_source_gen_test.go >> .make-singular-data-sources.log 2>&1; \
	if grep -iE '(error|fatal|panic|failed)' .make-singular-data-sources.log > /dev/null; then \
		echo "==> ⚠️  ERRORS DETECTED:"; \
		grep -iE '(error|fatal|panic|failed)' .make-singular-data-sources.log; \
		echo ""; \
		echo "==> Full output in .make-singular-data-sources.log"; \
		exit 1; \
	else \
		echo "==> ✓ Singular data sources generated successfully, no errors detected!"; \
		echo "==> Full output available in .make-singular-data-sources.log"; \
	fi

resources: prereq-go ## Generate resources
	@echo "==> Counting existing resource files..."
	@EXPECTED=$$(find internal -name '*_resource_gen.go' -o -name '*_resource_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
	if [ "$$EXPECTED" -eq 0 ]; then EXPECTED=1000; fi; \
	echo "==> Expected ~$$EXPECTED files to generate"; \
	rm -f internal/*/*/*_resource_gen.go; \
	rm -f internal/*/*/*_resource_gen_test.go; \
	echo "==> Generating resources..."; \
	( \
		while sleep 3; do \
			CURRENT=$$(find internal -name '*_resource_gen.go' -o -name '*_resource_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
			if [ "$$CURRENT" -gt 0 ]; then \
				PERCENT=$$((CURRENT * 100 / EXPECTED)); \
				if [ "$$PERCENT" -gt 100 ]; then PERCENT=100; fi; \
				printf "\r==> Progress: $$CURRENT/$$EXPECTED files ($$PERCENT%%) "; \
			fi; \
		done \
	) & PROGRESS_PID=$$!; \
	$(GO_VER) generate internal/provider/resources.go > .make-resources.log 2>&1; \
	GEN_EXIT=$$?; \
	kill $$PROGRESS_PID 2>/dev/null; wait $$PROGRESS_PID 2>/dev/null; \
	FINAL=$$(find internal -name '*_resource_gen.go' -o -name '*_resource_gen_test.go' 2>/dev/null | wc -l | tr -d ' '); \
	printf "\r==> Generated $$FINAL resource files          \n"; \
	if [ $$GEN_EXIT -ne 0 ]; then \
		echo "==> ⚠️  Generation failed - check .make-resources.log"; \
		exit $$GEN_EXIT; \
	fi; \
	echo "==> Running goimports on generated files..."; \
	goimports -w internal/*/*/*_resource_gen.go >> .make-resources.log 2>&1; \
	goimports -w internal/*/*/*_resource_gen_test.go >> .make-resources.log 2>&1; \
	if grep -iE '(error|fatal|panic|failed)' .make-resources.log > /dev/null; then \
		echo "==> ⚠️  ERRORS DETECTED:"; \
		grep -iE '(error|fatal|panic|failed)' .make-resources.log; \
		echo ""; \
		echo "==> Full output in .make-resources.log"; \
		exit 1; \
	else \
		echo "==> ✓ Resources generated successfully, no errors detected!"; \
		echo "==> Full output available in .make-resources.log"; \
	fi

schemas: prereq-go ## Generate schemas
	@echo "==> Generating schemas..."
	@$(GO_VER) generate internal/provider/schemas.go > .make-schemas.log 2>&1; \
	GEN_EXIT=$$?; \
	if [ $$GEN_EXIT -ne 0 ]; then \
		echo "==> ⚠️  Generation failed - check .make-schemas.log"; \
		cat .make-schemas.log; \
		exit $$GEN_EXIT; \
	fi; \
	if grep -iE '(error|fatal|panic|failed)' .make-schemas.log > /dev/null; then \
		echo "==> ⚠️  ERRORS DETECTED:"; \
		grep -iE '(error|fatal|panic|failed)' .make-schemas.log; \
		echo ""; \
		echo "==> Full output in .make-schemas.log"; \
		exit 1; \
	else \
		echo "==> ✓ Schemas generated successfully, no errors detected!"; \
		echo "==> Full output available in .make-schemas.log"; \
	fi

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
	rm -f docs/list-resources/*.md
	@tfplugindocs generate --provider-name "terraform-provider-awscc"

docs-fmt: prereq-go ## Format example Terraform files in documentation
	@echo "==> Formatting example Terraform files in documentation..."
	cd examples/resources/
	terraform fmt -recursive

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

update: prereq-go ## Update Schema
	echo "==> Updating Schema..."
	$(GO_VER) run $$(find internal/update -name "*.go" -not -name "*_test.go")

smoke: prereq-go ## Run smoke tests
	@echo "make: Running smoke tests..."
	@echo "make: NOTE: All tests should pass. Error output for sdk.proto, \"Response contains error diagnostic\" can be ignored."
	@TF_LOG=ERROR make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_\|TestAccAWSLogsLogGroupDataSource_' ACCTEST_PARALLELISM=3

commitdatas: ## Commit data source changes
	@git add -A && git commit -m "$$(date -I) CloudFormation schemas in us-east-1; Generate Terraform data source schemas"

commitresources: ## Commit resource changes
	@git add -A && git commit -m "$$(date -I) CloudFormation schemas in us-east-1; Generate Terraform resource schemas"

commitschemas: ## Commit schema changes
	@git add -A && git commit -m "$$(date -I) CloudFormation schemas in us-east-1; New schemas"

commitrefresh: ## Commit schema refresh changes
	@git add -A && git commit -m "$$(date -I) CloudFormation schemas in us-east-1; Refresh existing schemas"

commitdocs: ## Commit documentation changes
	@git add -A && git commit -m "$$(date -I) Documentation; Update generated documentation"

bigdiffer: ## Show big diff between current branch and main
	@echo "==> Showing big diff between this schema generation and last"
	@echo "==> Manually add each new resource/data source type to internal/provider/all_schemas.hcl"
	@LAST_VERSION_DATE=$$(git log -1 --format=%cd --date=format:%Y-%m-%d version/VERSION); \
	diff internal/provider/generators/allschemas/available_schemas.$$LAST_VERSION_DATE.hcl internal/provider/generators/allschemas/available_schemas.$$(date -I).hcl || true

newbranch: ## Create a new branch for schema updates
	@NEW_BRANCH="$(date -I)-schema-updates"; \
	echo "==> Creating and switching to new branch $$NEW_BRANCH"; \
	git checkout -b $$NEW_BRANCH

cleanschemas: ## Clean generated schema files
	@echo "==> Cleaning generated schema files..."
	rm internal/service/cloudformation/schemas/AWS_*.json

suppressions: ## Checkout suppressed schema files
	@echo "==> Checking out suppressions..."
	@echo "==> Currently updates to some schemas should be suppressed as they have changes which"
	@echo "==> prevent Terraform schema generation (or they no longer exist and are pending major version removal)"
	@echo "==> Add any new suppressions to the internal/update/suppressions_checkout.txt file"
	@SUPPRESSION_FILES=$$(grep -v '^#' internal/update/suppressions_checkout.txt | grep -v '^$$' | tr '\n' ' '); \
	if [ -n "$$SUPPRESSION_FILES" ]; then \
		echo "==> Reverting $$(echo $$SUPPRESSION_FILES | wc -w | tr -d ' ') suppressed schema files..."; \
		git checkout $$SUPPRESSION_FILES; \
	else \
		echo "==> No suppressions found"; \
	fi

biglister: prereq-go ## List all resources and data sources
	@echo "==> Listing all currently available resources and data sources..."
	@OUTPUT_FILE="internal/provider/generators/allschemas/available_schemas.$$(date -I).hcl"; \
	( \
		SECONDS=0; \
		while sleep 1; do \
			MINS=$$((SECONDS / 60)); \
			SECS=$$((SECONDS % 60)); \
			printf "\r==> Elapsed time: %02d:%02d " $$MINS $$SECS; \
			SECONDS=$$((SECONDS + 1)); \
		done \
	) & TIMER_PID=$$!; \
	$(GO_VER) run internal/provider/generators/allschemas/manual_allschemas/main.go > $$OUTPUT_FILE 2>&1; \
	EXIT_CODE=$$?; \
	kill $$TIMER_PID 2>/dev/null; wait $$TIMER_PID 2>/dev/null; \
	if [ $$EXIT_CODE -ne 0 ]; then \
		printf "\r==> ⚠️  Failed after %02d:%02d - check output\n" $$((SECONDS / 60)) $$((SECONDS % 60)); \
		cat $$OUTPUT_FILE; \
		exit $$EXIT_CODE; \
	else \
		printf "\r==> ✓ Completed in %02d:%02d\n" $$((SECONDS / 60)) $$((SECONDS % 60)); \
		echo "==> Output written to $$OUTPUT_FILE"; \
	fi

.PHONY: all
.PHONY: bigdiffer
.PHONY: biglister
.PHONY: build
.PHONY: cleanschemas
.PHONY: commitdatas
.PHONY: commitdocs
.PHONY: commitresources
.PHONY: commitschemas
.PHONY: commitrefresh
.PHONY: default
.PHONY: docs
.PHONY: docs-all
.PHONY: docs-fmt
.PHONY: docs-import
.PHONY: golangci-lint
.PHONY: help
.PHONY: importlint
.PHONY: lint
.PHONY: newbranch
.PHONY: plural-data-sources
.PHONY: resources
.PHONY: schemas
.PHONY: singular-data-sources
.PHONY: smoke
.PHONY: suppressions
.PHONY: test
.PHONY: testacc
.PHONY: tools
.PHONY: update
