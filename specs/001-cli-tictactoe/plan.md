# Implementation Plan: CLI Tic-Tac-Toe Game

**Branch**: `001-cli-tictactoe` | **Date**: 2025-11-23 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/001-cli-tictactoe/spec.md`

## Summary

Build a terminal-based tic-tac-toe game where two players alternate turns entering row and column positions. The game validates all input, detects win/draw conditions, and provides clear error messages. Implementation uses Go with the tview library for terminal UI, includes comprehensive linting, a Makefile for build automation, and GitHub Actions for cross-platform binary distribution (Windows, macOS, Linux).

## Technical Context

**Language/Version**: Go 1.21+
**Primary Dependencies**:
- tview (terminal UI library for interactive TUI)
- tcell (terminal handling, dependency of tview)
- golangci-lint (comprehensive linting tool)

**Storage**: N/A (in-memory game state only)
**Testing**: Go's built-in testing framework (`go test`) with table-driven tests
**Target Platform**: Cross-platform CLI (Windows, macOS x64/arm64, Linux x64/arm64)
**Project Type**: Single binary application
**Performance Goals**:
- Instant UI response (<16ms per frame for 60fps rendering)
- <100ms move validation
- <1ms win condition checking

**Constraints**:
- <50MB memory footprint
- <10MB binary size per platform
- No external runtime dependencies
- Single-player terminal (no networking)

**Scale/Scope**:
- ~500-1000 lines of Go code
- 4-6 packages (main, game, ui, validation)
- 3x3 board only (no configuration)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify alignment with TicTacToe Constitution (`.specify/memory/constitution.md`):

### Code Quality Standards
- [x] Feature scope allows functions ≤50 lines - Game logic naturally decomposes into small functions (validateInput, checkWin, updateBoard each <30 lines)
- [x] Complexity manageable (cyclomatic complexity ≤10) - Win detection uses 8 checks (3 rows + 3 cols + 2 diags), well under limit
- [x] No obvious duplication or magic numbers planned - BOARD_SIZE=3, player constants (X, O, Empty) will be named
- [x] Code review process defined - GitHub PRs required before merge to main

### Testing Requirements
- [x] Test-first approach planned (Red-Green-Refactor) - Will write tests before implementing game logic, validation, win detection
- [x] Test scenarios documented in spec - 2 user stories with 11 acceptance scenarios total
- [x] 80% coverage target achievable - Core game logic, validation, win detection all testable; UI layer may be integration-tested
- [x] Test pyramid strategy defined (70% unit, 20% integration, 10% e2e) - Unit: game logic/validation, Integration: board state management, E2E: full game playthrough
- [x] Performance: Unit tests <5s, full suite <60s - Expected ~50 unit tests, <100ms total runtime

### User Experience Standards (for user-facing features)
- [x] Response time targets defined (<100ms actions, <1s page loads, <200ms API p95) - Move processing <100ms, instant board display
- [x] Error handling strategy documented - Specific error messages for each validation failure type (see FR-006 in spec)
- [x] Accessibility requirements specified (keyboard nav, screen readers, 4.5:1 contrast) - Terminal UI inherits terminal accessibility; keyboard-only input
- [x] Responsive design planned (mobile, tablet, desktop) - Adapts to terminal size; minimum 80x24 recommended

### Performance Requirements
- [x] Memory budget defined (<50MB for game state) - 3x3 board = 9 cells = <100 bytes; total app <50MB
- [x] Algorithm complexity acceptable (O(1) or O(n)) - Win detection: O(1) fixed 8 checks; validation: O(1) bounds checking
- [x] AI move calculation target <500ms - N/A (no AI opponent in scope)
- [x] No blocking operations on main thread - All I/O through tview event loop; no goroutines needed for this scope

### Documentation & Maintainability
- [x] Architecture approach documented - See Project Structure and Phase 1 deliverables below
- [x] API documentation plan defined - Godoc comments on all exported functions; quickstart.md for users
- [x] Complex decisions explained with rationale - See research.md for tview vs other TUI libraries, Makefile targets, GitHub Actions matrix strategy

## Project Structure

### Documentation (this feature)

```text
specs/001-cli-tictactoe/
├── plan.md              # This file
├── research.md          # Phase 0: Technology choices and patterns
├── data-model.md        # Phase 1: Game state structures
├── quickstart.md        # Phase 1: User guide and developer setup
├── contracts/           # Phase 1: (N/A for this project - no external APIs)
└── tasks.md             # Phase 2: Implementation tasks (created by /speckit.tasks)
```

### Source Code (repository root)

```text
# Go project structure (standard layout)
.
├── main.go                 # Entry point, tview app initialization
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── Makefile                # Build automation (build, test, lint, clean targets)
├── .golangci.yml           # Linter configuration
├── .github/
│   └── workflows/
│       └── build.yml       # GitHub Actions: build, lint, test, release binaries
│
├── game/                   # Core game logic package
│   ├── board.go            # Board state management
│   ├── board_test.go       # Board unit tests
│   ├── player.go           # Player type and turn management
│   ├── player_test.go      # Player unit tests
│   ├── win.go              # Win/draw detection logic
│   └── win_test.go         # Win detection unit tests
│
├── validation/             # Input validation package
│   ├── input.go            # Row/column validation logic
│   └── input_test.go       # Validation unit tests
│
├── ui/                     # TUI package (tview components)
│   ├── app.go              # tview application setup
│   ├── board_view.go       # Board display component
│   ├── input_view.go       # Input form component
│   └── messages.go         # Error and status messages
│
└── integration_test.go     # Full game integration tests

# Build outputs (ignored by git)
bin/
├── tictactoe-linux-amd64
├── tictactoe-linux-arm64
├── tictactoe-darwin-amd64
├── tictactoe-darwin-arm64
└── tictactoe-windows-amd64.exe
```

**Structure Decision**: Using standard Go project layout with packages organized by domain responsibility. Single project structure (Option 1) selected because this is a standalone CLI application with no backend/frontend split. The `game/` package contains pure business logic (testable without UI), `validation/` handles input validation (also pure functions), and `ui/` contains tview-specific UI code.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

No violations. All constitution checks pass.
