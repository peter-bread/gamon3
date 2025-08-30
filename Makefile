.PHONY: all build clean

OUTPUT_DIR = bin
EXE        = gamon

all: clean build

build:
	mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(EXE) -v
	
clean:
	go clean
	rm -rf $(OUTPUT_DIR)
