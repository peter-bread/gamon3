.PHONY: build clean install goreleaser

BUILD_DIR = build

VERSION  	?= $(shell git describe --tags --dirty --always)
LDFLAGS		?= -X main.version=$(VERSION)

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR) -v -ldflags "$(LDFLAGS)"

test:
	go test -v ./...
	
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
