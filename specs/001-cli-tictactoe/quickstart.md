# Quickstart Guide: CLI Tic-Tac-Toe

**Last Updated**: 2025-11-23
**Version**: 1.0.0

## For Users

### Installation

#### Option 1: Download Pre-built Binary

1. Go to the [Releases page](https://github.com/YOUR_USERNAME/tictactoe/releases)
2. Download the binary for your platform:
   - **Linux (64-bit)**: `tictactoe-linux-amd64`
   - **Linux (ARM64)**: `tictactoe-linux-arm64`
   - **macOS (Intel)**: `tictactoe-darwin-amd64`
   - **macOS (Apple Silicon)**: `tictactoe-darwin-arm64`
   - **Windows (64-bit)**: `tictactoe-windows-amd64.exe`

3. Make the binary executable (Linux/macOS only):
   ```bash
   chmod +x tictactoe-*
   ```

4. Run the game:
   ```bash
   ./tictactoe-linux-amd64    # Linux
   ./tictactoe-darwin-arm64   # macOS
   tictactoe-windows-amd64.exe  # Windows
   ```

#### Option 2: Install from Source

**Prerequisites**: Go 1.21 or later

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/tictactoe.git
cd tictactoe

# Build and install
make install

# Run the game
tictactoe
```

### How to Play

1. **Start the game**: Run the executable or type `tictactoe` if installed
2. **Game begins**: An empty 3x3 board appears
3. **Player 1 (X) goes first**: Enter row and column numbers (1-3)
4. **Player 2 (O) goes next**: Enter row and column numbers (1-3)
5. **Alternate turns** until someone wins or the board is full
6. **Win condition**: Get three marks in a row (horizontal, vertical, or diagonal)
7. **Draw condition**: All cells filled with no winner

### Input Format

Enter moves as **row number** (1-3) then **column number** (1-3).

**Examples of valid input**:
- `1 2` (row 1, column 2)
- `1,2` (row 1, column 2)
- `1` then `2` (prompted separately)

**Board positions**:
```
     Col 1  Col 2  Col 3
Row 1  [1,1] [1,2] [1,3]
Row 2  [2,1] [2,2] [2,3]
Row 3  [3,1] [3,2] [3,3]
```

### Common Errors

**"Invalid position. Row and column must be between 1 and 3"**
- You entered a number outside the valid range (1-3)
- Solution: Enter numbers 1, 2, or 3 only

**"Position already occupied. Please choose an empty cell"**
- That cell already has an X or O in it
- Solution: Choose a different, empty cell

**"Invalid input format. Please enter numeric values for row and column"**
- You entered non-numeric characters
- Solution: Enter only numbers (1, 2, or 3)

**"Incomplete input. Please provide both row and column"**
- You only entered one number
- Solution: Enter both row and column

### Controls

- **Enter move**: Type row and column, press Enter
- **Quit game**: Press `Ctrl+C` or close terminal window

### System Requirements

- **Operating System**: Windows 10+, macOS 10.15+, or Linux (any recent distribution)
- **Terminal**: Any terminal emulator (80x24 minimum recommended)
- **Dependencies**: None (self-contained binary)
- **Memory**: <50MB RAM
- **Disk Space**: <10MB

---

## For Developers

### Prerequisites

- **Go**: Version 1.21 or later ([download](https://go.dev/dl/))
- **Make**: GNU Make (usually pre-installed on Linux/macOS, [Windows installation](https://gnuwin32.sourceforge.net/packages/make.htm))
- **Git**: For cloning the repository
- **golangci-lint**: (optional, for linting) ([installation](https://golangci-lint.run/usage/install/))

### Setup Development Environment

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/tictactoe.git
cd tictactoe

# Install dependencies
go mod download

# Verify setup
make test
```

### Project Structure

```
tictactoe/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Dependency lock file
├── Makefile                # Build automation
├── .golangci.yml           # Linter configuration
├── .github/workflows/      # CI/CD configuration
│   └── build.yml
├── game/                   # Core game logic
│   ├── board.go
│   ├── board_test.go
│   ├── player.go
│   ├── player_test.go
│   ├── win.go
│   └── win_test.go
├── validation/             # Input validation
│   ├── input.go
│   └── input_test.go
├── ui/                     # Terminal UI components
│   ├── app.go
│   ├── board_view.go
│   ├── input_view.go
│   └── messages.go
└── integration_test.go     # Full game tests
```

### Development Workflow

#### 1. Run Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run specific package tests
go test ./game/...
go test ./validation/...
```

**Constitution Requirement**: 80% coverage minimum, tests pass before committing

#### 2. Lint Code

```bash
# Run all linters
make lint

# Auto-fix formatting issues
make fmt
```

**Constitution Requirement**: No linting errors before committing

#### 3. Build Binary

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Output location: ./bin/
```

#### 4. Run Locally

```bash
# Build and run
make build
./bin/tictactoe

# Or run without building
go run main.go
```

### Makefile Targets

| Target | Description |
|--------|-------------|
| `make build` | Build binary for current platform (output: `./bin/tictactoe`) |
| `make test` | Run all tests with coverage report |
| `make lint` | Run golangci-lint (checks code quality) |
| `make fmt` | Format code with gofmt |
| `make clean` | Remove build artifacts (`./bin/` directory) |
| `make install` | Build and install to `$GOPATH/bin` |
| `make build-all` | Build binaries for all platforms (Linux, macOS, Windows) |
| `make coverage` | Generate HTML coverage report |
| `make help` | Show all available targets |

### Testing Guidelines

**Constitution Requirements**:
- ✅ Write tests BEFORE implementation (Red-Green-Refactor)
- ✅ Maintain 80%+ code coverage
- ✅ Test pyramid: 70% unit, 20% integration, 10% e2e
- ✅ All tests must pass in <60 seconds

**Test Structure**:
```go
// Table-driven test example
func TestCheckWin(t *testing.T) {
    tests := []struct {
        name   string
        board  Board
        player Cell
        want   bool
    }{
        {
            name: "empty board no win",
            board: Board{},
            player: X,
            want: false,
        },
        {
            name: "horizontal win row 0",
            board: Board{{X, X, X}, {O, O, Empty}, {Empty, Empty, Empty}},
            player: X,
            want: true,
        },
        // ... more test cases
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

**Coverage Target by Package**:
- `game/`: 90%+ (pure logic)
- `validation/`: 95%+ (pure functions)
- `ui/`: 50-60% (integration-heavy)
- **Overall**: 80%+

### Code Quality Standards

**Constitution Requirements**:
- ✅ Functions ≤50 lines
- ✅ Cyclomatic complexity ≤10
- ✅ No magic numbers (use named constants)
- ✅ No code duplication >3 lines
- ✅ All exported functions have godoc comments

**Example**:
```go
// BOARD_SIZE defines the dimensions of the tic-tac-toe board (3x3)
const BOARD_SIZE = 3

// CheckWin determines if the specified player has achieved three marks in a row
// on the given board. Returns true if the player has won, false otherwise.
//
// Complexity: O(1) - checks exactly 8 win conditions (3 rows, 3 cols, 2 diagonals)
func CheckWin(board Board, player Cell) bool {
    // Implementation...
}
```

### Contributing

1. **Create feature branch**: `git checkout -b feature/your-feature`
2. **Write tests first**: Ensure they fail (Red)
3. **Implement feature**: Make tests pass (Green)
4. **Refactor**: Clean up code while keeping tests green
5. **Run quality checks**:
   ```bash
   make test    # Must pass with 80%+ coverage
   make lint    # Must have no errors
   ```
6. **Commit changes**: `git commit -m "feat: your feature description"`
7. **Push and create PR**: `git push origin feature/your-feature`

**Pull Request Checklist**:
- [ ] All tests passing
- [ ] Coverage ≥80%
- [ ] Linting passed (no errors)
- [ ] Functions ≤50 lines
- [ ] Complexity ≤10
- [ ] Godoc comments on exported functions
- [ ] Constitution check passed

### CI/CD Pipeline

**GitHub Actions** runs automatically on push and PR:

1. **Lint Job**: Runs golangci-lint
2. **Test Job**: Runs tests with coverage report
3. **Build Job**: Builds binaries for all platforms
4. **Release Job** (on tag push): Creates GitHub release with binaries

**Local Testing** (before pushing):
```bash
# Run the same checks as CI
make lint
make test
make build-all
```

### Architecture Overview

```
┌─────────────────────────────────────────────────┐
│                   main.go                        │
│            (Application Entry Point)             │
└───────────────────┬─────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────┐
│                  ui/ Package                     │
│        (tview components, user interaction)      │
└───────┬──────────────────────────┬───────────────┘
        │                          │
        ▼                          ▼
┌───────────────────┐    ┌────────────────────────┐
│  game/ Package    │    │  validation/ Package   │
│   (Pure Logic)    │    │    (Pure Functions)    │
│ - Board state     │    │ - Input parsing        │
│ - Win detection   │    │ - Range checking       │
│ - Player tracking │    │ - Error generation     │
└───────────────────┘    └────────────────────────┘
```

**Package Responsibilities**:
- `game/`: Core business logic (no I/O, fully testable)
- `validation/`: Input validation (pure functions)
- `ui/`: Terminal UI (tview integration, event handling)
- `main`: Wiring and startup

### Debugging

**Enable verbose logging** (if implemented):
```bash
DEBUG=1 ./tictactoe
```

**Run with delve debugger**:
```bash
dlv debug
(dlv) break main.main
(dlv) continue
```

**Common issues**:
- **Import errors**: Run `go mod tidy`
- **Test failures**: Check coverage with `make coverage`, open `coverage.html`
- **Lint errors**: Run `make lint` for details, `make fmt` to auto-fix formatting

### Performance Profiling

```bash
# CPU profile
go test -cpuprofile=cpu.prof ./...
go tool pprof cpu.prof

# Memory profile
go test -memprofile=mem.prof ./...
go tool pprof mem.prof
```

**Constitution targets**:
- Move processing: <100ms
- Win detection: <1ms
- Total memory: <50MB

### Release Process

1. **Tag version**: `git tag v1.0.0`
2. **Push tag**: `git push origin v1.0.0`
3. **GitHub Actions**: Automatically builds and creates release
4. **Download**: Binaries attached to release

**Versioning**: Semantic versioning (MAJOR.MINOR.PATCH)
- MAJOR: Breaking changes
- MINOR: New features (backward compatible)
- PATCH: Bug fixes

---

## Support

- **Issues**: [GitHub Issues](https://github.com/YOUR_USERNAME/tictactoe/issues)
- **Documentation**: [Project README](../README.md)
- **Specification**: [Feature Spec](./spec.md)
- **Technical Plan**: [Implementation Plan](./plan.md)

---

## License

[Specify license here, e.g., MIT, Apache 2.0]
