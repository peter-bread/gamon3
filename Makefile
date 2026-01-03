.PHONY: build test cover clean install goreleaser

override VALID_BUILD_MODES := release debug

# Release builds by default, must explicitly set BUILD_MODE=debug for dev builds.
BUILD_MODE        ?= release

ifneq ($(filter $(BUILD_MODE),$(VALID_BUILD_MODES)),$(BUILD_MODE))
$(error Invalid BUILD_MODE '$(BUILD_MODE)'; must be one of: [$(VALID_BUILD_MODES)])
endif

# This is the default value. In contexts where VCS information is unavaliable,
# this needs to be set manually.
#
# For example, in a Homebrew formula, this set to 'version.to_s'.
GAMON3_VERSION ?= $(shell git describe --tags --dirty --always)

BIN            := bin
CGO_ENABLED    := 0

LDFLAGS_COMMON := -X github.com/peter-bread/gamon3/v2/internal/build.Version=$(GAMON3_VERSION)
GOFLAGS_COMMON := -v -buildvcs=auto

ifeq ($(BUILD_MODE), release)
LDFLAGS := -s -w $(LDFLAGS_COMMON)
GOFLAGS := -trimpath $(GOFLAGS_COMMON)
else ifeq ($(BUILD_MODE), debug)
LDFLAGS := $(LDFLAGS_COMMON)
GOFLAGS := $(GOFLAGS_COMMON)
endif

build:
	mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -o "$(BIN)" -ldflags "$(LDFLAGS)"

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
	
clean:
	go clean
	rm -rf $(BIN)


PREFIX ?= /usr/local

install: build
	install -d $(PREFIX)/bin
	install $(BIN)/gamon3 $(PREFIX)/bin

################################################################################

# This will use '.goreleaser.yaml' and build in 'dist/'.
# GAMON3_VERSION is passed to ensure version output is the same as using 'make'.
# This is just for checking that GoReleaser and Makefile builds are the same.
goreleaser:
	GAMON3_VERSION=$(GAMON3_VERSION) goreleaser release --snapshot --clean
	# goreleaser release --snapshot --clean
