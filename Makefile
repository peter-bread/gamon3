.PHONY: all build run clean

OUTPUT_DIR = bin
EXE        = gamon

all: clean build run

build:
	mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(EXE) -v
	
run:
	./$(OUTPUT_DIR)/$(EXE)

clean:
	go clean
	rm -rf $(OUTPUT_DIR)
