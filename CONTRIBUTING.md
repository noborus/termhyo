# Contributing to termhyo

Thank you for your interest in contributing to termhyo! This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Create a new branch for your feature or bug fix
4. Make your changes
5. Test your changes
6. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git

### Setup

```bash
git clone https://github.com/your-username/termhyo.git
cd termhyo
go mod download
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Update golden files (when output format changes)
make test-update
```

### Code Quality

Before submitting a pull request, ensure your code passes all checks:

```bash
# Run all quality checks
make check

# Individual checks
make fmt         # Format code
make vet         # Run go vet
make lint        # Run golint
make check-mod   # Check go.mod is tidy
```

## Code Style

- Follow standard Go formatting (`gofmt`)
- Write clear, descriptive variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and reasonably sized
- Use meaningful commit messages

## Testing

- Write tests for new functionality
- Ensure existing tests continue to pass
- Use golden files for output testing when appropriate
- Test with different character sets (ASCII, Unicode, Japanese, etc.)
- Test both buffered and streaming modes

### Golden File Testing

When adding new features that change table output:

```bash
# Generate new golden files
go test -update

# Verify golden files are correct
git diff testdata/
```

## Documentation

- Update README.md if adding new features
- Add examples for new functionality
- Document public APIs with Go doc comments
- Keep examples simple and focused

## Submitting Changes

### Pull Request Process

1. Update documentation as needed
2. Add tests for new functionality
3. Ensure all tests pass
4. Run `make check` to verify code quality
5. Create a pull request with a clear description

### Pull Request Guidelines

- Use a clear, descriptive title
- Describe what changes you made and why
- Link to any relevant issues
- Include screenshots for visual changes
- Keep pull requests focused and atomic

### Commit Messages

Use clear, descriptive commit messages:

- Start with a capital letter
- Use imperative mood ("Add feature" not "Added feature")
- Limit first line to 50 characters
- Add detailed description if needed

Good examples:

```text
Add support for 24-bit true color headers
Fix width calculation for ANSI escape sequences
Update examples with new header styling options
```

## Reporting Issues

When reporting bugs or requesting features:

1. Check existing issues first
2. Use issue templates when available
3. Provide clear reproduction steps for bugs
4. Include environment information (OS, Go version)
5. Add code samples when relevant

## Feature Requests

Before implementing new features:

1. Open an issue to discuss the feature
2. Ensure it fits the project's scope and goals
3. Consider backward compatibility
4. Plan the API design carefully

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help newcomers learn and contribute
- Maintain a welcoming environment

## Questions?

If you have questions about contributing:

- Open an issue for discussion
- Check existing documentation
- Look at existing code for examples

Thank you for contributing to termhyo!
