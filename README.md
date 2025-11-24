# Tic-Tac-Toe CLI Game

A command-line tic-tac-toe game for two players, built with Go following TDD principles.

- Instructions for running the app
Run `make build ` and then run `./bin/tictactoe`
- A brief description of your approach
Leveraged spec driven development with an AI coding assistant
- What AI tools you used and how
Spec kit for spec driven development
Claude Code with the Sonnet 4.5 LLM model
- Anything that didn’t go as planned or you'd improve with more time
I went for a simple terminal based app. With more time, I'd look into a buildout of a fullstack application with a frontend and backend with a network mode for players to play over the Internet.
The UI for the column indices on the board display could be centered. With more time, I'd revisit the original spec with spec kit to include centered labels for the board display and see what new tasks are generated from the plan.

## Features

- Two-player turn-based gameplay
- Clean command-line interface
- Comprehensive input validation with helpful error messages
- Automatic win and draw detection
- Immutable game state architecture
- Cross-platform support (Linux, macOS, Windows)

## Requirements

- Go 1.21 or higher

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/tictactoe.git
cd tictactoe

# Build the binary
make build

# Or build for all platforms
make build-all
```

### Pre-built Binaries

Download pre-built binaries from the `bin/` directory after building:

- `bin/tictactoe-linux-amd64` - Linux (AMD64)
- `bin/tictactoe-linux-arm64` - Linux (ARM64)
- `bin/tictactoe-darwin-amd64` - macOS (Intel)
- `bin/tictactoe-darwin-arm64` - macOS (Apple Silicon)
- `bin/tictactoe-windows-amd64.exe` - Windows (AMD64)

## Usage

Run the game:

```bash
./bin/tictactoe
```

### How to Play

1. The game displays a 3x3 grid with row and column numbers (0-2)
2. Players alternate turns, starting with Player 1 (X)
3. Enter your move as two numbers separated by space: `row column`
   - Example: `1 1` for center position
   - Example: `0 0` for top-left corner
4. The game automatically detects wins and draws
5. Invalid inputs show clear error messages with examples

### Example Game Session

```
=== Tic-Tac-Toe ===

  0   1   2
0     |     |
  -----------
1     |     |
  -----------
2     |     |

Player 1 (X)'s turn
Enter row and column (0-2), e.g., '1 1': 1 1

  0   1   2
0     |     |
  -----------
1     |  X  |
  -----------
2     |     |

Player 2 (O)'s turn
Enter row and column (0-2), e.g., '1 1': 0 0
...
```

## Development

### Build Commands

```bash
make build        # Build for current platform
make build-all    # Build for all platforms
make test         # Run tests with coverage
make lint         # Run linters
make coverage     # Generate HTML coverage report
make fmt          # Format code
make clean        # Remove build artifacts
```

### Project Structure

```
.
├── game/                  # Core game logic
│   ├── board.go          # Board and game state
│   ├── player.go         # Player types
│   ├── win.go            # Win/draw detection
│   ├── board_test.go     # Board tests
│   ├── player_test.go    # Player tests
│   └── win_test.go       # Win detection tests
├── validation/            # Input validation
│   ├── input.go          # Validation functions
│   └── input_test.go     # Validation tests
├── main.go               # CLI interface
├── integration_test.go   # Integration tests
├── Makefile              # Build automation
└── README.md             # This file
```

### Architecture

The game follows these design principles:

- **Immutability**: Game state changes return new instances rather than mutating existing state
- **Test-First Development**: All features developed with Red-Green-Refactor TDD cycle
- **Separation of Concerns**: Clear boundaries between game logic, validation, and UI
- **Constitution-Driven**: Adheres to strict code quality standards:
  - 80%+ test coverage (achieved: 85.1% in game, 100% in validation)
  - Functions ≤50 lines
  - Cyclomatic complexity ≤10
  - O(1) win detection algorithm

### Running Tests

```bash
# Run all tests
make test

# Run specific package tests
go test ./game/
go test ./validation/

# Run with verbose output
go test -v ./...

# Generate coverage report
make coverage
open coverage.html
```

### Test Coverage

Current coverage:
- `game/` package: 85.1%
- `validation/` package: 100%
- Overall: Exceeds 80% target

### Linting

The project uses golangci-lint with strict settings:

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linters
make lint
```

Enabled linters include:
- gofmt - Code formatting
- govet - Suspicious constructs
- staticcheck - Static analysis
- errcheck - Unchecked errors
- gocyclo - Cyclomatic complexity (max 10)
- funlen - Function length (max 50 lines)

## Code Quality

This project follows the TicTacToe Constitution v1.0.0 principles:

1. **Code Quality**: Max 50 lines per function, complexity ≤10, no magic numbers
2. **Test-First Development**: Red-Green-Refactor cycle, 80% coverage minimum
3. **User Experience Consistency**: Clear error messages, <100ms response time
4. **Performance Requirements**: O(1) algorithms, <50MB memory usage
5. **Maintainability**: Comprehensive documentation, DRY principles

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Follow TDD: Write tests first, then implementation
4. Ensure tests pass (`make test`)
5. Ensure linting passes (`make lint`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## License

This project is licensed under the MIT License.

## Acknowledgments

- Built following Test-Driven Development methodology
- Architecture inspired by functional programming principles
- Error handling patterns from Go best practices
