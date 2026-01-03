# Development Guide

## Prerequisites

- Go 1.21 or later
- Make

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/eduardoferro/md2blogger.git
cd md2blogger
```

2. Install dependencies:
```bash
go mod download
```

## Available Make Targets

Run `make help` to see all available targets:

```
make test      - Run all tests
make build     - Build the binary
make clean     - Remove build artifacts
make install   - Install the binary to GOPATH/bin
make lint      - Run golangci-lint
make fmt       - Format code with gofmt
make vet       - Run go vet
make validate  - Run all quality checks (fmt, vet, test)
```

## Development Workflow

### Running Tests

```bash
make test
```

### Building Locally

```bash
make build
./bin/md2blogger example.md
```

### Code Quality

Before committing, always run:

```bash
make validate
```

This will:
1. Format your code with `gofmt`
2. Run `go vet` for static analysis
3. Run all tests with race detection

## Project Structure

```
.
├── cmd/
│   └── md2blogger/     # CLI entry point
│       ├── main.go
│       └── main_test.go
├── converter/          # Markdown to HTML conversion
│   ├── converter.go
│   └── converter_test.go
├── docs/              # Documentation
│   └── development_guide.md
├── bin/               # Built binaries (gitignored)
├── Makefile           # Build automation
├── go.mod             # Go module definition
└── README.md          # User documentation
```

## Testing

The project uses Go's built-in testing framework. Tests are located alongside the code they test.

### Running Specific Tests

```bash
go test ./converter -v
go test ./cmd/md2blogger -v
```

### Coverage Report

```bash
make test
go tool cover -html=coverage.out
```

## Dependencies

- **goldmark**: Modern Markdown parser
- **goldmark/extension**: GFM (GitHub Flavored Markdown) and table support

## Adding New Features

1. Write a failing test first (TDD)
2. Implement the feature
3. Ensure tests pass: `make test`
4. Validate code quality: `make validate`
5. Update documentation if needed

## Release Process

1. Update version in code
2. Run `make validate` to ensure all checks pass
3. Tag the release: `git tag v1.0.0`
4. Push: `git push origin v1.0.0`
5. GitHub Actions will build and publish binaries (if configured)

## Troubleshooting

### Tests failing after changes

Run `make validate` to see all errors at once.

### Build errors

Ensure you have Go 1.21+ installed:
```bash
go version
```

Clean and rebuild:
```bash
make clean
make build
```
