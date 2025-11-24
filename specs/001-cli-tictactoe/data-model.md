# Data Model: CLI Tic-Tac-Toe Game

**Date**: 2025-11-23
**Feature**: 001-cli-tictactoe

## Overview

This document defines the core data structures for the tic-tac-toe game. All structures are designed to be immutable where possible, support efficient validation, and align with constitution performance requirements (O(1) operations, <50MB memory).

## Core Entities

### 1. Cell

Represents a single position on the game board.

**Type**: Enumeration (3 possible values)

**Values**:
- `Empty`: Unoccupied cell (initial state)
- `X`: Cell occupied by Player 1
- `O`: Cell occupied by Player 2

**Representation** (Go):
```go
type Cell int

const (
    Empty Cell = iota  // 0
    X                  // 1
    O                  // 2
)
```

**Properties**:
- **Memory**: 1 byte per cell (int8 sufficient)
- **Validation**: Only three valid states
- **Rendering**: Maps to display characters ('_' or ' ' for Empty, 'X' for X, 'O' for O)

**Operations**:
- `String()`: Convert to display character
- `IsOccupied()`: Check if cell is not Empty

---

### 2. Board

Represents the 3x3 game board containing all cells.

**Type**: Fixed-size 2D array

**Structure**:
```
     Col 0   Col 1   Col 2
Row 0  [0,0] [0,1] [0,2]
Row 1  [1,0] [1,1] [1,2]
Row 2  [2,0] [2,1] [2,2]
```

**Representation** (Go):
```go
type Board [3][3]Cell
```

**Properties**:
- **Size**: Fixed 3x3 (BOARD_SIZE constant = 3)
- **Memory**: 9 bytes (9 cells × 1 byte each)
- **Immutability**: Pass by value or use copy-on-write for modifications
- **Indexing**: Zero-based (0-2), but user input is 1-based (1-3)

**Operations**:
- `NewBoard()`: Create empty board (all cells Empty)
- `GetCell(row, col)`: Get cell value at position
- `SetCell(row, col, cell)`: Return new board with cell updated (immutable)
- `IsCellEmpty(row, col)`: Check if position is Empty
- `IsFull()`: Check if all 9 cells are occupied
- `CountOccupied()`: Count non-Empty cells
- `String()`: Render board as ASCII art

**Invariants**:
- Rows and columns always in range [0, 2]
- Board never contains invalid Cell values
- Once a cell is set to X or O, it cannot change during game

---

### 3. Player

Represents one of the two players in the game.

**Type**: Enumeration (mapped to Cell type)

**Values**:
- `Player1`: Uses mark X, goes first
- `Player2`: Uses mark O, goes second

**Representation** (Go):
```go
// Player is represented by their mark (Cell type)
type Player Cell

const (
    Player1 = Player(X)
    Player2 = Player(O)
)
```

**Properties**:
- **Identity**: Defined by mark (X or O)
- **Turn order**: Player1 always starts
- **Memory**: 1 byte (same as Cell)

**Operations**:
- `GetMark()`: Return Cell type (X or O)
- `Name()`: Return display name ("Player 1 (X)" or "Player 2 (O)")
- `Other()`: Return opposite player (Player1 → Player2, Player2 → Player1)

---

### 4. Move

Represents a player's action (placing their mark on the board).

**Type**: Struct containing position and player

**Structure**:
```go
type Move struct {
    Row    int    // 0-2 (internal representation)
    Col    int    // 0-2 (internal representation)
    Player Player // Which player made the move
}
```

**Properties**:
- **Validation**: Row and Col must be in range [0, 2]
- **User input**: Users enter 1-3, converted to 0-2 internally
- **Memory**: ~12 bytes (2 ints + 1 byte player)

**Operations**:
- `NewMove(row, col, player)`: Create validated move
- `IsValid(board)`: Check if move can be applied (position empty)
- `Apply(board)`: Return new board with move applied

**Validation Rules**:
- Row ∈ [0, 2]
- Col ∈ [0, 2]
- Target cell must be Empty
- Player must be valid (Player1 or Player2)

---

### 5. GameState

Represents the current state of the game (in progress, won, or draw).

**Type**: Enumeration

**Values**:
- `InProgress`: Game ongoing, no winner yet
- `Player1Won`: Player 1 achieved three in a row
- `Player2Won`: Player 2 achieved three in a row
- `Draw`: All cells filled, no winner

**Representation** (Go):
```go
type GameState int

const (
    InProgress GameState = iota
    Player1Won
    Player2Won
    Draw
)
```

**Properties**:
- **Terminal states**: Player1Won, Player2Won, Draw (game ends)
- **Transitions**: InProgress → (Player1Won | Player2Won | Draw)
- **No reverse**: Once terminal state reached, game cannot continue

**Operations**:
- `IsTerminal()`: Check if game has ended
- `Winner()`: Return winning player (if applicable)
- `String()`: Display message ("Player 1 wins!", "It's a draw!", etc.)

---

### 6. Game

Represents the complete game context (board + state + current player).

**Type**: Struct containing all game information

**Structure**:
```go
type Game struct {
    Board         Board      // Current board state
    CurrentPlayer Player     // Whose turn it is
    State         GameState  // Current game status
    MoveCount     int        // Number of moves made (0-9)
}
```

**Properties**:
- **Initialization**: Empty board, Player1 starts, InProgress state, 0 moves
- **Immutability**: Each move returns new Game instance
- **Memory**: ~32 bytes total

**Operations**:
- `NewGame()`: Initialize new game
- `MakeMove(row, col)`: Apply move, switch player, check win/draw
- `ValidateMove(row, col)`: Check if move is legal
- `GetWinner()`: Return winning player if game ended
- `IsDraw()`: Check if game ended in draw

**Invariants**:
- If State is terminal (won/draw), CurrentPlayer doesn't change
- MoveCount ∈ [0, 9]
- MoveCount == 9 ⇒ (State == Draw ∨ State == Player1Won ∨ State == Player2Won)
- State == InProgress ⇒ MoveCount < 9

**State Transitions**:
```
NewGame() → InProgress (Player1's turn, empty board, 0 moves)

InProgress → InProgress:  MakeMove() when no winner and board not full
                         (switch CurrentPlayer, increment MoveCount)

InProgress → Player1Won:  MakeMove() by Player1 creates three-in-a-row
InProgress → Player2Won:  MakeMove() by Player2 creates three-in-a-row
InProgress → Draw:        MakeMove() fills last cell with no winner
```

---

## Win Conditions

Win condition detection checks 8 possible three-in-a-row patterns:

**Horizontal Wins** (3 patterns):
- Row 0: (0,0), (0,1), (0,2)
- Row 1: (1,0), (1,1), (1,2)
- Row 2: (2,0), (2,1), (2,2)

**Vertical Wins** (3 patterns):
- Col 0: (0,0), (1,0), (2,0)
- Col 1: (0,1), (1,1), (2,1)
- Col 2: (0,2), (1,2), (2,2)

**Diagonal Wins** (2 patterns):
- Main diagonal: (0,0), (1,1), (2,2)
- Anti-diagonal: (0,2), (1,1), (2,0)

**Algorithm** (O(1) complexity per constitution):
```go
func CheckWin(board Board, player Cell) bool {
    // Check rows
    for row := 0; row < 3; row++ {
        if board[row][0] == player && board[row][1] == player && board[row][2] == player {
            return true
        }
    }

    // Check columns
    for col := 0; col < 3; col++ {
        if board[0][col] == player && board[1][col] == player && board[2][col] == player {
            return true
        }
    }

    // Check diagonals
    if board[0][0] == player && board[1][1] == player && board[2][2] == player {
        return true
    }
    if board[0][2] == player && board[1][1] == player && board[2][0] == player {
        return true
    }

    return false
}
```

**Complexity**: Fixed 8 checks = O(1) time, O(1) space

---

## Input Validation Model

User input is transformed through validation pipeline:

**Input**: String (e.g., "1 2", "1,2", "1")
**Step 1**: Parse to integers (row, col)
**Step 2**: Validate format (two integers present)
**Step 3**: Validate range (1-3 for user input)
**Step 4**: Convert to internal coordinates (1-3 → 0-2)
**Step 5**: Validate cell is empty
**Output**: Valid Move or specific error

**Error Mapping**:
```
Parse failure → ErrInvalidFormat
Single value → ErrIncompleteInput
Out of range → ErrInvalidRange (row or col < 1 or > 3)
Cell occupied → ErrCellOccupied
```

---

## Memory Analysis

**Total Memory per Game**:
- Board: 9 bytes
- CurrentPlayer: 1 byte
- State: 1 byte
- MoveCount: 4 bytes (int32)
- Total: ~16 bytes

**UI Buffers** (tview):
- Display buffer: ~1-2 MB
- Event queue: <100 KB

**Go Runtime**:
- Base: ~5-10 MB
- GC overhead: ~2-5 MB

**Total Application Memory**: ~10-20 MB (well under 50MB constitution limit)

---

## Persistence

**Decision**: No persistence layer

**Rationale**:
- Spec explicitly excludes save/load game state
- In-memory only aligns with simplicity principle
- Game sessions < 2 minutes (per success criteria)

If persistence added later:
- Serialize Game struct to JSON
- Store in user home directory (~/.tictactoe/saves/)
- Load on startup with prompt

---

## Testing Considerations

**Test Data Builders**:
```go
// Helper functions for tests
func NewEmptyBoard() Board
func NewBoardWithMoves(moves ...Move) Board
func NewWinningBoard(player Player, pattern string) Board  // pattern: "row0", "col1", "diag"
func NewDrawBoard() Board
```

**Property-Based Tests**:
- Any sequence of 5+ valid moves results in valid GameState
- CheckWin is symmetric (same result regardless of check order)
- MakeMove maintains board cell count invariant (occupied cells only increase)

**Edge Cases**:
- All 8 win conditions
- Draw with various patterns (no three-in-a-row)
- Boundary inputs (1 and 3 are valid, 0 and 4 are not)
- Occupied cell attempts

---

## Summary

All data structures are:
- **Simple**: No unnecessary complexity
- **Immutable**: Where possible (Board, Move)
- **Performant**: O(1) operations, minimal memory (<50MB total)
- **Testable**: Pure functions, no hidden state
- **Type-safe**: Enums for player marks and game state

This model supports all functional requirements (FR-001 through FR-013) while adhering to constitution standards.
