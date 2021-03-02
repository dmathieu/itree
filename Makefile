TOP_LEVEL = $(shell git rev-parse --show-toplevel)
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
TOOLS_DIR = $(TOP_LEVEL)/.tools
CIRCLECI_DIR = $(TOP_LEVEL)/.circleci

# Make sure this is in-sync with the version in the circle ci config
GOLANGCI_LINT_VERSION := 1.32.2
GOLANGCI_LINT_URL := https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-$(GOOS)-$(GOARCH).tar.gz
GOLANGCI_LINT := $(TOOLS_DIR)/golangci-lint-v$(GOLANGCI_LINT_VERSION)
CIRCLECI_CONFIG := $(CIRCLECI_DIR)/config.yml
PROCESSED_CIRCLECI_CONFIG := $(CIRCLECI_DIR)/.processed.yml
PKG_SPEC := ./...
MOD := -mod=readonly
GOTEST := go test $(MOD)
GOTEST_OPT := -race
COVER_PROFILE = coverage.out
GOTEST_COVERAGE_OPT := -coverprofile=$(COVER_PROFILE) -covermode=atomic

# Additive or overridable variables
override GOTEST_OPT += -timeout 30s
LINT_RUN_OPTS ?= --fix

$(TOOLS_DIR):
	mkdir -p $(TOOLS_DIR)

# Ensures the correct version of golangci-lint is present
$(GOLANGCI_LINT):
	rm -f $(TOOLS_DIR)/golangci-lint*
	mkdir -p $(TOOLS_DIR)
	curl -L $(GOLANGCI_LINT_URL) | tar -zxf - -C $(TOOLS_DIR) --strip=1 golangci-lint-$(GOLANGCI_LINT_VERSION)-$(GOOS)-$(GOARCH)/golangci-lint
	mv $(TOOLS_DIR)/golangci-lint $(GOLANGCI_LINT)

.PHONY: help
help: # Prints out help
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
			IFS=$$':' ; \
			help_split=($$help_line) ; \
			help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			printf '\033[36m'; \
			printf "%-30s %s" $$help_command ; \
			printf '\033[0m'; \
			printf "%s\n" $$help_info; \
	done

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Runs golangci-lint. Override defaults with LINT_RUN_OPTS
	$(GOLANGCI_LINT) run $(LINT_RUN_OPTS) $(PKG_SPEC)

.PHONY: test
test: ## Runs go test. Override defaults with GOTEST_OPT
	$(GOTEST) $(GOTEST_OPT) $(PKG_SPEC)

.PHONY: bench
bench: ## Runs go test with benchmarks
	$(GOTEST) -bench=. $(PKG_SPEC)

.PHONY: coverage
coverage: ## Generates a coverage profile and opens a web browser with the results
	$(GOTEST) $(GOTEST_OPT) $(GOTEST_COVERAGE_OPT) $(PKG_SPEC)
	go tool cover -html=$(COVER_PROFILE)

# Processes the circle ci config locally
$(CIRCLECI_CONFIG):
$(PROCESSED_CIRCLECI_CONFIG): $(CIRCLECI_CONFIG)
	circleci config process $(CIRCLECI_CONFIG) > $(PROCESSED_CIRCLECI_CONFIG)

.PHONY: ci-test
ci-test: ## Runs the ci based test job locally
ci-test: $(PROCESSED_CIRCLECI_CONFIG)
	circleci local execute --job test -c $(PROCESSED_CIRCLECI_CONFIG) -v "$(GOPATH)/pkg":/go/pkg
