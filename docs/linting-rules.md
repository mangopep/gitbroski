# Linting Rules and Standards

This document describes the linting rules, standards, and expectations for the GitBroski project.

## Code Quality Standards

### Documentation Requirements

**This project follows a minimal documentation approach to keep production code lean.**

#### Package Documentation (REQUIRED)

Every package must have **one** package-level comment:

```go
// Package commands provides CLI command handlers for gitbroski operations.
package commands
```

**Requirements:**

- Must start with "Package <name>"
- Should describe the package's purpose in one line
- Place directly above `package` declaration
- **Only one comment per package is required**

#### Function, Variable, and Type Documentation (NOT REQUIRED)

**⚠️ Important:** This project does NOT require comments for:

- Exported functions
- Exported variables
- Exported constants
- Exported types

The linter has been configured to skip these checks (`revive` exported rule is disabled).

**Example - No comments needed:**

```go
// Package git provides utilities for interacting with git repositories.
package git

// No comment needed for exported function
func GetRemoteURL() string {
    // ...
}

// No comment needed for exported variable
var Registry = make(map[string]func(args ...string))

// No comment needed for exported constant
const DefaultTimeout = 30 * time.Second
```

This approach keeps the codebase clean while maintaining essential package-level documentation.

### Error Handling

All errors must be handled appropriately:

#### Required Error Checking

```go
// ❌ Bad - error ignored
result, _ := SomeFunction()

// ✅ Good - error handled
result, err := SomeFunction()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// ✅ Acceptable - error explicitly ignored with reason
result, _ := fmt.Println("message") // Print errors are non-critical
```

#### Error Wrapping

Use error wrapping for context:

```go
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

### Security Standards

#### File Permissions

Use appropriate file permissions:

```go
// Sensitive files (configs, keys, etc.)
os.WriteFile(configPath, data, 0o600) // Owner read/write only

// Regular files
os.WriteFile(dataPath, data, 0o644)   // Owner r/w, others read

// Executables
os.WriteFile(scriptPath, data, 0o755) // Owner all, others r/x
```

#### Subprocess Execution

Be careful with subprocess execution:

```go
// ❌ Dangerous - user input in command
cmd := exec.Command("sh", "-c", userInput)

// ✅ Safe - controlled arguments
cmd := exec.Command("git", "status")

// ✅ Safe with validation
if !isValidBranch(branch) {
    return errors.New("invalid branch name")
}
cmd := exec.Command("git", "checkout", branch)
```

### Code Style

#### Modern Go Syntax

Use modern Go syntax (Go 1.13+):

```go
// ❌ Old octal literals
perm := 0644

// ✅ New octal literals
perm := 0o644

// ❌ Old error handling
if err != nil {
    return nil, err
}

// ✅ Modern error wrapping
if err != nil {
    return nil, fmt.Errorf("operation failed: %w", err)
}
```

#### Unused Parameters

Mark unused parameters explicitly:

```go
// ❌ Misleading
func Handler(args ...string) {
    // args never used
}

// ✅ Clear intent
func Handler(_ ...string) {
    // Parameter required by interface but not used
}
```

#### Import Organization

Imports should be organized in groups:

```go
import (
    // Standard library
    "fmt"
    "os"

    // External dependencies
    "github.com/fatih/color"

    // Internal packages
    "gitbroski/internal/commands"
    "gitbroski/utils/logger"
)
```

### Code Complexity

Keep functions simple and focused:

- **Cyclomatic complexity**: Maximum 15 (enforced by gocyclo)
- **Function length**: Keep under 50 lines when possible
- **Function parameters**: Maximum 5 parameters

If a function exceeds these limits, consider refactoring:

```go
// ❌ Too complex
func ProcessEverything(a, b, c, d, e, f string) error {
    // 100 lines of code
    // Multiple nested if statements
}

// ✅ Refactored
func ProcessEverything(opts Options) error {
    if err := validateOptions(opts); err != nil {
        return err
    }

    if err := processStep1(opts); err != nil {
        return fmt.Errorf("step 1 failed: %w", err)
    }

    return processStep2(opts)
}
```

## Linter-Specific Rules

### errcheck

**Purpose:** Ensures all errors are checked

**Rule:** Every error return value must be handled

**Exceptions:**

- Print functions where errors are non-critical
- Deferred Close() calls (use `defer func() { _ = f.Close() }()`)

### govet

**Purpose:** Catches suspicious constructs

**Common Issues:**

- Composite literals without field names
- Printf format string issues
- Unreachable code
- Shadow variables

### staticcheck

**Purpose:** Advanced static analysis

**Enforces:**

- Package-level comments (ST1000)
- Proper error handling patterns
- Deprecation warnings
- Performance issues

### revive

**Purpose:** Style and best practices

**Key Rules:**

- Package comments required ✅
- Proper naming conventions (camelCase, not snake_case)
- No unused parameters without `_` prefix
- Error strings should not be capitalized or end with punctuation

### gocritic

**Purpose:** Code quality diagnostics

**Checks:**

- Unnecessary code
- Style issues
- Performance problems
- Common mistakes

### gosec

**Purpose:** Security vulnerability detection

**Checks:**

- Weak random number generation
- SQL injection possibilities
- Path traversal vulnerabilities
- Insecure file permissions
- Unsafe subprocess execution

### misspell

**Purpose:** Spell checking

**Checks:**

- Common English misspellings in:
  - Comments
  - Documentation
  - String literals
  - Variable names

## Severity Levels

| Level   | Meaning         | Action Required |
| ------- | --------------- | --------------- |
| Error   | Must be fixed   | Block merge     |
| Warning | Should be fixed | Review required |
| Info    | Nice to fix     | Optional        |

Current configuration: All issues are treated as **errors** by default.

## Exclusion Rules

Some files/patterns are excluded from specific linters:

### Test Files

Test files (`*_test.go`) are excluded from:

- `gocyclo` - Tests can be complex
- `errcheck` - Test errors often checked differently
- `dupl` - Test code often repeats
- `gosec` - Tests may use insecure patterns safely
- `gocritic` - Tests have different style needs

### Generated Files

Generated files are automatically excluded:

- `*.pb.go` - Protocol buffer generated code
- `*.gen.go` - Other generated code

## Auto-fixable Issues

These issues can be automatically fixed with `make lint-fix`:

- Import organization (goimports)
- Code formatting (gofmt)
- Some gocritic issues
- Some staticcheck suggestions
- Unused imports

## Non-fixable Issues

These require manual intervention:

- Missing package documentation (rare - only one comment per package needed)
- Logic errors
- Security vulnerabilities
- Complex refactoring needs

## Configuration Changes

### When to Modify Configuration

Consider modifying `.golangci.yml` when:

- A linter produces too many false positives
- Project needs change
- New linters become available
- Performance issues arise

### How to Propose Changes

1. Open an issue explaining the need
2. Provide examples of the problem
3. Suggest specific configuration changes
4. Test changes locally first
5. Submit PR with updated configuration

### Testing Configuration Changes

```bash
# Validate new configuration
golangci-lint config --validate

# Run with new configuration
golangci-lint run ./...

# Compare results
golangci-lint run ./... > new-results.txt
git checkout main
golangci-lint run ./... > old-results.txt
diff old-results.txt new-results.txt
```

## Performance Considerations

### Fast Mode

For quick checks during development:

```bash
golangci-lint run --fast ./...
```

This runs only fast linters, skipping:

- Expensive type checking
- Some static analysis
- Code duplication detection

### Parallel Execution

Linters run in parallel by default. Controlled by:

```yaml
run:
  allow-parallel-runners: true
```

### Caching

golangci-lint caches results. Clear cache if needed:

```bash
golangci-lint cache clean
```

## Version Compatibility

| Tool          | Version | Notes               |
| ------------- | ------- | ------------------- |
| golangci-lint | 2.5.0+  | Required            |
| Go            | 1.25.1+ | Project requirement |
| Configuration | v2      | Current format      |

## Future Improvements

Planned linting enhancements:

1. **Additional Linters**

   - `nilnil` - No nil return with nil error
   - `exhaustive` - Exhaustive enum switch checks
   - `goheader` - License header validation

2. **CI/CD Integration**

   - Automated PR comments with issues
   - Trend tracking over time
   - Quality gates

3. **Documentation**
   - Video tutorials
   - Common mistake examples
   - Best practice guide

## Resources

- [golangci-lint Linters](https://golangci-lint.run/usage/linters/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://go.dev/doc/effective_go)
- [Uber Go Style Guide](https://github.com/uber-go/guide)

---

**Maintained by:** GitBroski Team  
**Last Updated:** October 7, 2025  
**Version:** 1.0
