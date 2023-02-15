VERSION ?= 0.0.1

# We want to automatically embed some build information into our executables:
BUILD_VERSION ?= $(shell git branch --show-current)
BUILD_COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP ?= $(shell date -u '+%Y-%m-%d %H:%M:%S')

GOLDFLAGS += -X 'stackit/secrets-manager-api/internal/version.buildVersion=$(BUILD_VERSION)'
GOLDFLAGS += -X 'stackit/secrets-manager-api/internal/version.buildCommit=$(BUILD_COMMIT)'
GOLDFLAGS += -X 'stackit/secrets-manager-api/internal/version.buildTimestamp=$(BUILD_TIMESTAMP)'

.PHONY: all
all: build

.PHONY: prepare
prepare: golangci-lint openapi
	go mod tidy
	go fmt ./...
	go vet ./...
	$(GOLANGCI_LINT) run

.PHONY: build
build: prepare
	go build -ldflags "$(GOLDFLAGS)" ./cmd/stackit-secrets-manager

.PHONY: update-openapi-spec
update-openapi-spec:
	curl -o scripts/openapi.yaml https://docs.api.eu01.stackit.cloud/oas/secrets-manager

.PHONY: release
release:
	git tag -a v$(VERSION) -m "Release v$(VERSION)"
	git push origin v$(VERSION)
	$(GORELEASER) check
	$(GORELEASER) release

.PHONY: release-local
release-local:
	$(GORELEASER) release --snapshot --clean

# ===== auto generate client from openapi yaml spec =====
.PHONY: openapi
openapi: internal/api/secrets_manager.gen.go
internal/api/secrets_manager.gen.go: scripts/oapi-codegen.yaml scripts/openapi.yaml oapi-codegen
	$(OAPI_CODEGEN) --config $(word 1,$^) $(word 2,$^) > $@

# ===== golangci-lint =====
GOLANGCI_LINT = bin/golangci-lint
.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT)
$(GOLANGCI_LINT):
	scripts/install_golangci-lint.sh

# ===== oapi-codegen =====
OAPI_CODEGEN = bin/oapi-codegen
.PHONY: oapi-codegen
oapi-codegen: $(OAPI_CODEGEN)
$(OAPI_CODEGEN):
	scripts/install_oapi-codegen.sh

# ===== goreleaser =====
GORELEASER = bin/goreleaser
.PHONY: goreleaser
goreleaser: $(GORELEASER)
$(GORELEASER):
	scripts/install_goreleaser.sh
