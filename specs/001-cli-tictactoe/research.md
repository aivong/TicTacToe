# Research & Design Decisions: CLI Tic-Tac-Toe Game

**Date**: 2025-11-23
**Feature**: 001-cli-tictactoe

## Overview

This document captures technical research and design decisions for implementing the CLI tic-tac-toe game. All choices prioritize simplicity, testability, and alignment with the constitution requirements.

## Technology Decisions

### 1. Terminal UI Library: tview

**Decision**: Use [rivo/tview](https://github.com/rivo/tview) for terminal UI

**Rationale**:
- **High-level abstractions**: Provides forms, text views, layouts without manual terminal manipulation
- **Built on tcell**: Robust cross-platform terminal handling (Windows, macOS, Linux)
- **Event-driven**: Natural fit for interactive input without blocking operations
- **Active maintenance**: 10k+ stars, actively maintained, good documentation
- **No external dependencies**: Compiles to single binary
- **Performance**: Efficient rendering, meets <16ms frame time requirement

**Alternatives Considered**:
- **bubbletea**: Modern, but more complex for simple use case; better for larger TUIs
- **termui**: Dashboard-focused, overkill for tic-tac-toe
- **Raw tcell**: Too low-level, would require manual component building
- **Standard library only**: Would require complex terminal control sequences, poor cross-platform support

**Implementation Approach**:
- Use `tview.Form` for row/column input with built-in validation
- Use `tview.TextView` for board display with box drawing characters
- Use `tview.Modal` for win/draw announcements
- Single-page layout with board at top, input form below

---

### 2. Linting: golangci-lint

**Decision**: Use golangci-lint as primary linting tool

**Rationale**:
- **Comprehensive**: Runs 50+ linters in parallel (gofmt, govet, staticcheck, etc.)
- **Fast**: Caches results, optimized for CI/CD
- **Configurable**: `.golangci.yml` allows enabling/disabling specific linters
- **Industry standard**: Used by most Go projects
- **GitHub Actions integration**: Official action available

**Configuration Strategy**:
```yaml
# .golangci.yml highlights
linters:
  enable:
    - gofmt          # Code formatting
    - govet          # Suspicious constructs
    - staticcheck    # Advanced static analysis
    - errcheck       # Unchecked errors
    - gosimple       # Simplifications
    - ineffassign    # Unused assignments
    - gocyclo        # Cyclomatic complexity (max 10 per constitution)
    - gofumpt        # Stricter gofmt
    - misspell       # Spelling
    - gocritic       # Opinionated checks
```

**Constitution Alignment**:
- `gocyclo` enforces complexity ≤10
- `funlen` enforces function length ≤50 lines
- `gofmt` ensures consistent formatting

**Alternatives Considered**:
- **Individual linters**: Too manual, golangci-lint aggregates best ones
- **Staticcheck only**: Good but limited scope compared to golangci-lint
- **No linting**: Violates constitution requirement for code quality standards

---

### 3. Build Automation: Makefile

**Decision**: Use GNU Make for build automation

**Rationale**:
- **Universal**: Available on all platforms (Windows via make.exe or WSL)
- **Simple**: Declarative targets for common operations
- **Fast**: Only rebuilds changed files
- **Self-documenting**: `make help` can list targets
- **POSIX standard**: Familiar to all developers

**Makefile Targets**:
```makefile
.PHONY: build test lint clean install help

build:          # Build for current platform
test:           # Run all tests with coverage
lint:           # Run golangci-lint
clean:          # Remove build artifacts
install:        # Install binary to $GOPATH/bin
build-all:      # Build for all platforms (Linux, macOS, Windows x64/arm64)
coverage:       # Generate coverage report (target ≥80%)
fmt:            # Format code with gofmt
vet:            # Run go vet
```

**Alternatives Considered**:
- **Task**: Modern but requires separate installation
- **Just**: Similar to Task, not as widely adopted
- **Shell scripts**: Platform-specific, harder to maintain
- **go build directly**: No automation of lint/test/multi-platform builds

---

### 4. CI/CD: GitHub Actions

**Decision**: Use GitHub Actions for automated builds and releases

**Rationale**:
- **Native GitHub integration**: No external service needed
- **Matrix builds**: Easily build for multiple OS/arch combinations
- **Free for public repos**: No cost consideration
- **Artifact storage**: Can attach binaries to releases
- **Rich ecosystem**: Actions marketplace for golangci-lint, Go setup, etc.

**Workflow Strategy**:
```yaml
# .github/workflows/build.yml structure
on: [push, pull_request]

jobs:
  lint:
    - Setup Go
    - Run golangci-lint

  test:
    - Setup Go
    - Run tests with coverage
    - Upload coverage to artifacts

  build:
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        arch: [amd64, arm64]
    - Build binary for OS/arch
    - Upload binaries as artifacts

  release: (on tag push)
    - Create GitHub release
    - Attach all platform binaries
```

**Build Matrix**:
- Linux: amd64, arm64
- macOS: amd64 (Intel), arm64 (Apple Silicon)
- Windows: amd64

**Alternatives Considered**:
- **goreleaser**: Excellent but adds complexity; overkill for this project
- **Travis CI**: Deprecated for open source
- **CircleCI**: Requires separate service, less integrated
- **GitLab CI**: Would require GitLab, user specified GitHub

---

## Architecture Decisions

### 5. Package Structure

**Decision**: Three-package architecture (game, validation, ui)

**Rationale**:
- **Separation of concerns**: Game logic independent of UI
- **Testability**: Pure functions in game/validation packages easily unit tested
- **Constitution compliance**: Clear single responsibility per package
- **No over-engineering**: Simpler than ports-and-adapters or clean architecture

**Package Responsibilities**:
- `game/`: Board state, player tracking, win detection (pure logic, no I/O)
- `validation/`: Input validation (pure functions, no state)
- `ui/`: tview components, user interaction (only package with side effects)
- `main`: Wiring and application entry point

**Rejected Patterns**:
- **Single package**: Would violate single responsibility, hard to test
- **Clean architecture layers**: Over-engineered for 500-1000 LOC project
- **Repository pattern**: No persistence layer needed

---

### 6. Game State Management

**Decision**: Immutable board state with copy-on-write

**Rationale**:
- **Thread safety**: No concurrent modifications (not needed here, but good practice)
- **Testability**: Pure functions easier to test
- **Predictability**: No hidden state changes
- **Constitution alignment**: Avoids blocking operations

**Data Structure**:
```go
type Cell int
const (
    Empty Cell = iota
    X
    O
)

type Board [3][3]Cell

// Pure function: returns new board, doesn't modify input
func MakeMove(board Board, row, col int, player Cell) Board
```

**Alternatives Considered**:
- **Mutable board**: Simpler but harder to test, risk of unintended mutations
- **Pointer-based**: Unnecessary complexity for 9-cell array

---

### 7. Win Detection Algorithm

**Decision**: Exhaustive check of 8 win conditions

**Rationale**:
- **O(1) complexity**: Fixed 8 checks regardless of board size (constitution requirement)
- **Simple**: Easy to understand and maintain
- **Fast**: <1ms execution (constitution requirement)
- **Comprehensive**: No missed edge cases

**Algorithm**:
```go
// Check all 8 win conditions:
// - 3 rows (0,0)-(0,2), (1,0)-(1,2), (2,0)-(2,2)
// - 3 cols (0,0)-(2,0), (0,1)-(2,1), (0,2)-(2,2)
// - 2 diags (0,0)-(2,2), (0,2)-(2,0)
func CheckWin(board Board, player Cell) bool
```

**Alternatives Considered**:
- **Generic n-in-a-row**: Over-engineered for fixed 3x3 board
- **Dynamic checking from last move**: More complex, negligible performance gain

---

### 8. Error Handling Strategy

**Decision**: Sentinel errors with descriptive messages

**Rationale**:
- **Constitution requirement**: Clear, actionable error messages (FR-006)
- **Type safety**: Can distinguish error types programmatically
- **User-friendly**: Maps directly to UI error display

**Error Types**:
```go
var (
    ErrInvalidRange = errors.New("Invalid position. Row and column must be between 1 and 3")
    ErrInvalidFormat = errors.New("Invalid input format. Please enter numeric values")
    ErrCellOccupied = errors.New("Position already occupied. Please choose an empty cell")
    ErrIncompleteInput = errors.New("Incomplete input. Please provide both row and column")
)
```

**Alternatives Considered**:
- **Custom error types**: Over-engineered for this scope
- **String errors**: Not type-safe, harder to test
- **Panic**: Violates constitution (graceful error handling required)

---

## Testing Strategy

### 9. Test Organization

**Decision**: Table-driven tests for core logic, integration tests for full game

**Rationale**:
- **Constitution requirement**: 80% coverage, test pyramid (70/20/10)
- **Go best practice**: Table-driven tests idiomatic for Go
- **Comprehensive**: Easy to add many test cases without duplication
- **Fast**: Unit tests run in <100ms total

**Test Structure**:
```go
// game/win_test.go
func TestCheckWin(t *testing.T) {
    tests := []struct {
        name   string
        board  Board
        player Cell
        want   bool
    }{
        {"empty board", Board{}, X, false},
        {"horizontal win row 0", Board{{X,X,X},{O,O,Empty},{Empty,Empty,Empty}}, X, true},
        {"vertical win col 0", Board{{X,O,Empty},{X,O,Empty},{X,Empty,Empty}}, X, true},
        {"diagonal win", Board{{X,O,Empty},{O,X,Empty},{Empty,Empty,X}}, X, true},
        {"no win yet", Board{{X,O,X},{O,X,O},{O,Empty,Empty}}, X, false},
        // ... more cases
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := CheckWin(tt.board, tt.player)
            if got != tt.want {
                t.Errorf("CheckWin() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

**Coverage Targets** (per constitution):
- `game/` package: 90%+ (pure logic, fully testable)
- `validation/` package: 95%+ (pure functions, all edge cases)
- `ui/` package: 50-60% (tview integration, harder to unit test)
- Overall: 80%+ target

---

## Build and Deployment

### 10. Cross-Platform Binary Strategy

**Decision**: Use `GOOS`/`GOARCH` environment variables for cross-compilation

**Rationale**:
- **Go native**: Built into Go toolchain, no external tools needed
- **Simple**: Single command per platform
- **Reliable**: Official Go feature, well-tested

**Build Commands**:
```makefile
BINARY_NAME=tictactoe

build-linux-amd64:
    GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64

build-linux-arm64:
    GOOS=linux GOARCH=arm64 go build -o bin/$(BINARY_NAME)-linux-arm64

build-darwin-amd64:
    GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-darwin-amd64

build-darwin-arm64:
    GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)-darwin-arm64

build-windows-amd64:
    GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64.exe
```

**Alternatives Considered**:
- **goreleaser**: Excellent tool but adds dependency and complexity
- **Docker multi-stage builds**: Overkill for Go's native cross-compilation
- **Platform-specific build machines**: Unnecessary, Go cross-compiles natively

---

## Performance Considerations

### 11. Memory Optimization

**Decision**: Stack allocation for board state, minimal heap usage

**Rationale**:
- **Constitution requirement**: <50MB memory footprint
- **Performance**: Stack allocation faster than heap
- **Garbage collection**: Minimal GC pressure

**Approach**:
- Board is `[3][3]Cell` (value type, stack-allocated, ~9 bytes)
- No dynamic allocations in game loop
- tview components reused, not recreated per frame

**Expected Memory Usage**:
- Game state: <100 bytes
- tview buffers: ~1-2 MB
- Go runtime: ~5-10 MB
- Total: <20 MB (well under 50MB limit)

---

### 12. Build Size Optimization

**Decision**: Use standard Go build, optional UPX compression

**Rationale**:
- **Constitution constraint**: <10MB binary size target
- **Standard build**: Go produces reasonably small binaries (~5-8MB for simple CLI)
- **Optional compression**: UPX can reduce to ~2-3MB if needed

**Build Flags**:
```bash
# Standard build (~6MB)
go build -o tictactoe

# Stripped build (~5MB)
go build -ldflags="-s -w" -o tictactoe

# With UPX compression (~2MB) - optional
upx --best tictactoe
```

**Decision**: Start with stripped build (`-ldflags="-s -w"`), add UPX only if needed

---

## Summary

All technology and architecture decisions align with constitution requirements:
- **Code Quality**: Package structure enforces separation of concerns
- **Testing**: Table-driven tests support 80% coverage target
- **Performance**: O(1) algorithms, <50MB memory, <100ms response
- **Maintainability**: Standard Go idioms, clear package boundaries
- **Tooling**: golangci-lint, Makefile, GitHub Actions provide automation

All decisions favor simplicity over complexity, testability over cleverness, and standard practices over custom solutions.
