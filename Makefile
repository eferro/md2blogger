.PHONY: help test build clean install lint fmt vet validate

help:
	@echo "Available targets:"
	@echo "  test      - Run all tests"
	@echo "  build     - Build the binary"
	@echo "  clean     - Remove build artifacts"
	@echo "  install   - Install the binary to GOPATH/bin"
	@echo "  lint      - Run golangci-lint"
	@echo "  fmt       - Format code with gofmt"
	@echo "  vet       - Run go vet"
	@echo "  validate  - Run all quality checks (fmt, vet, test)"

test:
	go test -v -race -coverprofile=coverage.out ./...

build:
	go build -o bin/md2blogger ./cmd/md2blogger

clean:
	rm -rf bin/ dist/ coverage.out coverage.html

install:
	go install ./cmd/md2blogger

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

vet:
	go vet ./...

validate: fmt vet test
	@echo "All validation checks passed!"
