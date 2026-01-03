# AI Agent Development Rules - Go Project

This document contains all development rules and guidelines for this Go project, applicable to all AI agents (Claude, Gemini, etc.).

## 1. Core Principles

- **Baby Steps**: Always work in baby steps, one at a time. Never go forward more than one step.
- **Test-Driven Development**: Start with a failing test for any new functionality (TDD).
- **Progressive Revelation**: Never show all the code at once; only the next step.
- **Type Safety**: Leverage Go's static type system fully. Avoid `interface{}` unless absolutely necessary. Use type parameters (generics) when appropriate (Go 1.18+).
- **Simplicity First**: Follow Go's philosophy: simplicity over cleverness. Use the simplest working solution; avoid unnecessary abstractions and premature optimization.
- **Small, Focused Functions**: Keep functions focused and readable. Aim for 20-40 lines as a guideline, but prioritize clarity and cohesion. A well-structured 50-line function is better than artificially split code.
- **Clear Naming**: Use clear, descriptive names following Go conventions (e.g., `mixedCaps` for unexported, `MixedCaps` for exported, short names for short scopes like `i` for index).
- **Incremental Changes**: Prefer incremental, focused changes over large, complex modifications.
- **Question Assumptions**: Always question assumptions and inferences.
- **Refactoring Awareness**: Highlight opportunities for refactoring when functions become hard to understand or test.
- **Pattern Detection**: Detect and highlight repeated code patterns that could benefit from abstraction.

## 2. Code Quality & Coverage (Go-Specific)

- **MANDATORY Validation**: Before EVERY commit, run `make validate` and fix ALL errors. Zero tolerance.
- **Quality Requirements**: The project has strict requirements for code quality and maintainability.
- **High Coverage**: All code must have very high test coverage; strive for 80%+ where practical, 100% for critical business logic.
- **Pre-commit Checks**: All code must pass before any commit:
  - `gofmt` (formatting)
  - `go vet` (static analysis)
  - All tests with `-race` flag (race condition detection)
- **TDD Workflow**: Test-Driven Development (TDD) is the default workflow: always write tests first.
- **Idiomatic Go**: Follow Go idioms and best practices:
  - Prefer composition over inheritance (use embedding and interfaces)
  - Accept interfaces, return structs
  - Handle errors explicitly, don't ignore them
  - Use `defer` for cleanup
  - Follow "Effective Go" guidelines
  - Keep the happy path left-aligned (avoid deep nesting)

## 3. Style Guidelines

- **Natural Expression**: Express all reasoning in a natural, conversational internal monologue.
- **Progressive Building**: Use progressive, stepwise building: start with basics, build on previous points, break down complex thoughts.
- **Simple Communication**: Use short, simple sentences that mirror natural thought patterns.
- **Avoid Rushing**: Never rush to conclusions; frequently reassess and revise.
- **Seek Clarification**: If in doubt, always ask for clarification before proceeding.
- **Self-Documenting Code**: Avoid comments in code; rely on self-documenting names. Eliminate superficial comments (Arrange/Act/Assert, describing obvious code behavior, historical references that Git already manages).

## 4. Communication Style

- Think through problems before responding
- Show reasoning when solving complex issues
- Ask for clarification when requirements are unclear
- Keep responses focused and professional

## 5. Process & Key Requirements

- **Show Work**: Show reasoning when solving complex issues.
- **Embrace Uncertainty**: Embrace uncertainty and revision.
- **Persistence**: Persist through multiple attempts until resolution.
- **Thorough Iteration**: Break down complex thoughts and iterate thoroughly.
- **Sequential Questions**: Only one question at a time; each question should build on previous answers.

## 6. Language Standards

- **Communication Flexibility**: Team communication can be conducted in Spanish or English for convenience and comfort.
- **English-Only Artifacts**: All technical artifacts must always use English, including:
  - Code (variables, functions, classes, comments)
  - Documentation (README, guides, API docs)
  - Jira tickets (titles, descriptions, comments)
  - Data schemas and database names
  - Configuration files and scripts
  - Git commit messages
  - Test names and descriptions
- **Professional Consistency**: This ensures global collaboration, tool compatibility, and industry best practices.

## 7. Documentation Standards

- **User-Focused README**: README.md must be user-focused, containing only information relevant to end users.
- **Separate Dev Docs**: All developer, CI, and infrastructure documentation must be placed in a separate development guide (e.g., docs/development_guide.md), with a clear link from the README.
- **Error Examples**: User-facing documentation should include example error messages for common validation failures to help users quickly resolve issues.

## 8. Development Best Practices

### Error Handling & Debugging
- **Graceful Error Handling**: Always implement proper error handling with meaningful error messages.
- **Debugging First**: When encountering issues, use debugging tools and logging before asking for help.
- **Error Context**: Provide sufficient context in error messages to enable quick problem resolution.
- **Fail Fast**: Design code to fail fast and fail clearly when errors occur.

### Code Review & Collaboration  
- **Pair Programming**: Prefer pairing sessions for complex features and knowledge sharing.
- **Small Pull Requests**: Keep changes small and focused for easier review and faster integration.
- **Code Review Standards**: All code must be reviewed before merging, following project quality standards.
- **Knowledge Sharing**: Document decisions and share context with team members.

### Security Considerations
- **Security by Design**: Consider security implications in all design decisions.
- **Input Validation**: Always validate and sanitize user inputs and external data.
- **Secrets Management**: Never hardcode secrets; use proper secret management systems.
- **Dependency Security**: Regularly update dependencies and monitor for security vulnerabilities.

### Testing Strategy Distinction
- **Unit Tests**: Fast, isolated tests for individual components (majority of test suite).
- **Integration Tests**: Test interactions between components and external systems (limited, focused).
- **E2E Tests**: Full system validation (minimal, critical user paths only).
- **Test Pyramid**: Follow the test pyramid - many unit tests, some integration tests, few E2E tests.

## 9. Test-Driven Development Rules

### TDD Approach
- **Failing Test First**: Always start with a failing test before implementing new functionality.
- **Single Test**: Write only one test at a time; never create more than one test per change.
- **Complete Coverage**: Ensure every new feature or bugfix is covered by a test.

### Test Structure & Style (Go-Specific)
- **Consistent Tooling**: Use Go's standard `testing` package. Use testify/assert only if the project already has it.
- **Test File Organization**: Tests go in `*_test.go` files in the same package. Integration tests can use `package_test` to test public API.
- **Focused Tests**: Keep each test focused and readable. Aim for 20-30 lines, but prioritize clarity over strict line limits.
- **Clear Naming**: Use Go's test naming convention: `TestFunctionName` for unit tests, `TestFunctionName_WhenCondition` for specific scenarios.
- **Table-Driven Tests**: Use table-driven tests for testing multiple scenarios of the same behavior.
- **No Comments**: Avoid comments; make code self-documenting through naming.
- **Helper Functions**: Extract common setup into helper functions. Prefix with `test` for unexported helpers.
- **Strategic Mocking**: Prefer testing real implementations where practical. Use interfaces for dependencies that need mocking (databases, APIs, file systems). Define interfaces in the consumer package, not the provider.

### Test Simplicity & Maintainability
- **Simplest Setup**: Prefer the simplest test setup that covers the requirement.
- **Refactor Tests**: Refactor tests to remove duplication and improve readability.
- **Consistent Assertions**: Use one assertion style consistently throughout the suite.
- **Extract Helpers**: If a test setup is repeated, extract a helper or fixture.
- **Readable Tests**: Always keep tests readable and easy to modify.

### Test Process & Output
- **Single Test Display**: Only show one test at a time; never present multiple tests in a single step.
- **Single File Display**: Never show more than one file at a time.
- **Self-Contained Tests**: Each test should be self-contained and not depend on the order of execution.
- **Clarify Requirements**: If in doubt about requirements, ask for clarification before writing the test.
- **Verify Failure**: After writing a test, run it to ensure it fails before implementing the feature.
- **Automatic Test Running**: After every code or test change, always run the relevant tests using the project's task runner. Do not ask for permission to run tests—just do it.

### Test Naming & Coverage
- **Descriptive Names**: Test function names should clearly describe the scenario and expected outcome.
- **Purpose-Driven Variables**: Use descriptive variable names that reflect their purpose in the test.
- **Incremental Coverage**: Ensure all code paths and edge cases are eventually covered by tests, but add them incrementally.

### Test Review & Refactoring
- **Post-Pass Review**: After a test passes, review for opportunities to simplify or clarify.
- **Helper Refactoring**: Refactor test helpers and fixtures as needed to keep the suite DRY and maintainable.

## 10. Task Runner Usage

### Core Rule
**NEVER** call tools like `pytest`, `black`, `mypy`, or similar directly. Always use the project's task runner (e.g., Make, npm scripts, Poetry scripts, Gradle, etc.).

### Discovering Available Tasks (Go Projects)
Before starting work, identify available tasks:
- **Makefile**: Run `make help` or inspect the `Makefile` (most common for Go projects)
- Check for `mage` (magefile.go) or `task` (Taskfile.yml) if no Makefile exists

### Common Task Categories (Go Projects)
Go projects typically provide these tasks:
- **test**: Run all tests with coverage (`go test -v -race -coverprofile=coverage.out ./...`)
- **build**: Build the binary (`go build -o bin/...`)
- **install**: Install to GOPATH/bin (`go install`)
- **fmt**: Format code (`gofmt -s -w .`)
- **vet**: Static analysis (`go vet ./...`)
- **validate**: Run all quality checks (fmt, vet, test)
- **clean**: Remove build artifacts
- **lint**: Run golangci-lint (if configured)

### Usage Rules
1. **Discover first**: Always check what tasks are available before running any tool.
2. **Use task runner**: Never call underlying tools directly; use the configured task.
3. **Add new tasks**: If a new operation is needed, prefer adding a new task rather than running a tool directly.

### Good vs Bad Examples (Go)
```sh
# Good: Use task runner
make test             # Run tests via Makefile
make validate         # Run all quality checks
make build            # Build the binary

# Bad: Call tools directly
go test ./...
gofmt -w .
go vet ./...
```

## 11. Pre-Commit Validation (MANDATORY)

Before ANY commit:
1. Run the project's validation task (e.g., `make validate`, `npm run validate`, `./gradlew check`)
2. If errors exist: fix them and re-run
3. Only commit when validation passes with ZERO errors

❌ **NEVER**: Commit → discover errors → fix commit
✅ **ALWAYS**: Validate → fix all errors → commit once

## 12. Go Project Structure

### Standard Layout
Follow the Go community's standard project layout:
- **`cmd/`**: Main applications for this project. Directory name should match the binary name.
  - `cmd/myapp/main.go` → builds `myapp` binary
- **`pkg/`**: Library code that's ok to import by external applications (optional, use sparingly)
- **`internal/`**: Private application and library code. Cannot be imported by other projects.
- **Root directory**: Library packages if this is a library project
- **`docs/`**: Documentation beyond README
- **`scripts/`**: Build, install, analysis, etc. scripts (if not using Make)

### Package Organization
- Keep packages small and focused
- Avoid circular dependencies
- Group by domain/functionality, not by layer (avoid packages named `models`, `controllers`, `services`)
- Example good structure:
  ```
  cmd/md2blogger/     # CLI entrypoint
  converter/          # Markdown conversion logic
  renderer/           # HTML rendering (if needed)
  internal/parser/    # Internal parsing utilities
  ```

### File Organization
- One concept per file (e.g., `user.go`, `user_test.go`)
- Keep `main.go` minimal in `cmd/` directories
- Tests in same package as code (`*_test.go`)
- Integration tests can use `package_test` suffix

### Dependency Management
- Use Go modules (`go.mod`, `go.sum`)
- Run `go mod tidy` regularly to clean up dependencies
- Vendor dependencies only if necessary (`go mod vendor`)

## 13. Quick Reference for Go Development

When working on this Go project:

1. **Take baby steps** - one test, one file, one change at a time 👣
2. **Always write the failing test first** (TDD) ❌➡️✅
3. **Use Makefile** - never call go commands directly (`make test`, not `go test`) 🔧
4. **Keep functions focused** - aim for 20-40 lines, prioritize clarity 📏
5. **Show reasoning for complex issues** - be clear and professional 💭
6. **Question everything** - assumptions, requirements, design choices ❓
7. **Run `make validate` before EVERY commit** - zero tolerance ✅
8. **Run tests automatically** after every change (`make test`) 🧪
9. **Follow Go idioms** - simplicity, composition, explicit errors ✨
10. **Ask for clarification** when in doubt 🤔

### Go-Specific Reminders
- Accept interfaces, return concrete types
- Handle errors explicitly, never ignore them
- Use `defer` for cleanup
- Keep the happy path left-aligned (minimal nesting)
- Prefer table-driven tests
- Use `gofmt`, `go vet`, and race detector

Remember: This is a high-quality, test-driven, incremental Go development environment. Simplicity over cleverness, clarity over complexity, idiomatic Go over clever hacks. 