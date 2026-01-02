.PHONY: build test cover clean install goreleaser

override VALID_BUILD_MODES := release debug

# Release builds by default, must explicitly set BUILD_MODE=debug for dev builds.
BUILD_MODE        ?= release

ifneq ($(filter $(BUILD_MODE),$(VALID_BUILD_MODES)),$(BUILD_MODE))
$(error Invalid BUILD_MODE '$(BUILD_MODE)'; must be one of: [$(VALID_BUILD_MODES)])
endif

BUILD_DIR       = build
CGO_ENABLED     = 0

LDFLAGS_COMMON :=
GOFLAGS_COMMON := -v -buildvcs=true

ifeq ($(BUILD_MODE), release)
LDFLAGS := -s -w $(LDFLAGS_COMMON)
GOFLAGS := -trimpath $(GOFLAGS_COMMON)
else  ifeq ($(BUILD_MODE), debug)
LDFLAGS := $(LDFLAGS_COMMON)
GOFLAGS := $(GOFLAGS_COMMON)
endif

build:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -o "$(BUILD_DIR)" -ldflags "$(LDFLAGS)"

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
	
clean:
	go clean
	rm -rf $(BUILD_DIR)


PREFIX ?= /usr/local

install: build
	install -d $(PREFIX)/bin
	install $(BUILD_DIR)/gamon3 $(PREFIX)/bin

################################################################################

# This will use '.goreleaser.yaml' and build in 'dist/'.
goreleaser:
	goreleaser release --snapshot --clean
