.PHONY: all build clean

OUTPUT_DIR = dist

all: clean build

build:
	goreleaser release --snapshot --clean
	
clean:
	go clean
	rm -rf $(OUTPUT_DIR)
