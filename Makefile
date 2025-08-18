BINARY=gogrep
BIN_DIR=bin
TEST_SCRIPT=./integration/test_e2e.sh

.PHONY: build test integration lint clean

build:
	@mkdir -p $(BIN_DIR)
	go build -o ${BIN_DIR}/${BINARY} ./cmd/gogrep

test:
	go test -v ./...

integration: build
	bash ${TEST_SCRIPT}

lint:
	go vet ./...
	golangci-lint run ./...

clean:
	rm -rf ${BIN_DIR} out*.txt ref*.txt
	rm -f gogrep