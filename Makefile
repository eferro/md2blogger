.PHONY: help test build clean install lint fmt vet validate

# Installation directory (can be overridden: make install INSTALL_DIR=/usr/local/bin)
INSTALL_DIR ?= $(HOME)/.local/bin
BINARY_NAME = md2blogger

help:
	@echo "Available targets:"
	@echo "  test      - Run all tests"
	@echo "  build     - Build the binary"
	@echo "  clean     - Remove build artifacts"
	@echo "  install   - Install the binary to $(INSTALL_DIR)"
	@echo "              Override with: make install INSTALL_DIR=/custom/path"
	@echo "  lint      - Run golangci-lint"
	@echo "  fmt       - Format code with gofmt"
	@echo "  vet       - Run go vet"
	@echo "  validate  - Run all quality checks (fmt, vet, test)"

test:
	go test -v -race -coverprofile=coverage.out ./...

build:
	go build -o bin/$(BINARY_NAME) ./cmd/md2blogger

clean:
	rm -rf bin/ dist/ coverage.out coverage.html

install: build
	@mkdir -p $(INSTALL_DIR)
	@cp bin/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)"
	@echo "Make sure $(INSTALL_DIR) is in your PATH"

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

vet:
	go vet ./...

validate: fmt vet test
	@echo "All validation checks passed!"
