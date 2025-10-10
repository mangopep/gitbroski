# Quick Linting Reference

A quick reference guide for common linting tasks and fixes in the GitBroski project.

## Quick Commands

```bash
# Run linting
golangci-lint run

# Auto-fix issues
golangci-lint --fix

# Install tools
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0
```

## Common Fixes Cheatsheet

### Add Package Comment (Required)

```go
// Package <name> describes what this package does.
package name
```

**Note:** Only package comments are required. Function and variable comments are NOT required in this project.

### Handle Errors

```go
// Ignore error safely
_, _ = function()

// Handle error properly
if err := function(); err != nil {
    // handle error
}
```

### Fix File Permissions

```go
// Use 0o600 for secure files
os.WriteFile(path, data, 0o600)
```

### Mark Unused Parameters

```go
// Before
func Handler(unused string) {}

// After
func Handler(_ string) {}
```

## Issue Categories

| Category                | Count  | Priority |
| ----------------------- | ------ | -------- |
| Missing package comment | Rare   | Medium   |
| Unchecked errors        | Common | High     |
| Security issues         | Rare   | Critical |
| Style issues            | Common | Low      |

**Note:** Function and variable comments are NOT required in this project.

## When to Use #nosec

```go
// Only when you're certain the code is safe
// #nosec G204 - explain why it's safe
cmd := exec.Command("git", safeArg)
```

## Exclude Patterns

Add to `.golangci.yml`:

```yaml
issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck
```

## Getting Help

```bash
# List all linters
golangci-lint help linters

# Validate config
golangci-lint config --validate

# Run specific linter
golangci-lint run --disable-all -E errcheck
```
