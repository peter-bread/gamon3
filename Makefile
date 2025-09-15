.PHONY: build clean goreleaser install

BUILD_DIR = build
LDFLAGS   =

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR) -v -ldflags "$(LDFLAGS)"
	
clean:
	go clean
	rm -rf $(BUILD_DIR)

# This will use '.goreleaser.yaml' and build in 'dist/'.
goreleaser:
	goreleaser release --snapshot --clean

install:
	echo "TODO: Implement 'make install'."
