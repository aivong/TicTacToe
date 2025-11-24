// Package game contains the core tic-tac-toe game logic
package game

// Cell represents a single position on the game board
type Cell int

const (
	// Empty represents an unoccupied cell
	Empty Cell = iota
	// X represents a cell occupied by Player 1
	X
	// O represents a cell occupied by Player 2
	O
)

// String returns the string representation of a cell
func (c Cell) String() string {
	switch c {
	case Empty:
		return " "
	case X:
		return "X"
	case O:
		return "O"
	default:
		return "?"
	}
}

// IsOccupied returns true if the cell is not empty
func (c Cell) IsOccupied() bool {
	return c != Empty
}

// BOARD_SIZE defines the dimensions of the tic-tac-toe board (3x3)
const BOARD_SIZE = 3

// Board represents a 3x3 tic-tac-toe game board
type Board [BOARD_SIZE][BOARD_SIZE]Cell

// NewBoard creates and returns a new empty board
func NewBoard() Board {
	return Board{} // All cells initialized to Empty (zero value)
}

// GetCell returns the cell value at the specified position
// row and col must be in range [0, 2]
func (b Board) GetCell(row, col int) Cell {
	return b[row][col]
}

// SetCell returns a new board with the cell at the specified position set to the given value
// This function is immutable - it does not modify the original board
// row and col must be in range [0, 2]
func (b Board) SetCell(row, col int, cell Cell) Board {
	newBoard := b
	newBoard[row][col] = cell
	return newBoard
}

// IsCellEmpty returns true if the cell at the specified position is empty
// row and col must be in range [0, 2]
func (b Board) IsCellEmpty(row, col int) bool {
	return b[row][col] == Empty
}

// IsFull returns true if all cells on the board are occupied
func (b Board) IsFull() bool {
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if b[row][col] == Empty {
				return false
			}
		}
	}
	return true
}

// CountOccupied returns the number of non-empty cells on the board
func (b Board) CountOccupied() int {
	count := 0
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if b[row][col] != Empty {
				count++
			}
		}
	}
	return count
}

// GameState represents the current state of the game
type GameState int

const (
	// InProgress indicates the game is still being played
	InProgress GameState = iota
	// Player1Won indicates Player 1 (X) has won
	Player1Won
	// Player2Won indicates Player 2 (O) has won
	Player2Won
	// Draw indicates the game ended in a draw
	Draw
)

// Game represents the complete game context
type Game struct {
	Board         Board     // Current board state
	CurrentPlayer Player    // Whose turn it is
	State         GameState // Current game status
	MoveCount     int       // Number of moves made (0-9)
}

// NewGame creates and returns a new game instance
func NewGame() Game {
	return Game{
		Board:         NewBoard(),
		CurrentPlayer: Player1,
		State:         InProgress,
		MoveCount:     0,
	}
}

// MakeMove applies a move at the specified position and returns the new game state
// Returns an error if the move is invalid (out of bounds or cell occupied)
func (g Game) MakeMove(row, col int) (Game, error) {
	// Validate position is within bounds
	if row < 0 || row >= BOARD_SIZE || col < 0 || col >= BOARD_SIZE {
		return g, ErrInvalidRange
	}

	// Validate cell is empty
	if !g.Board.IsCellEmpty(row, col) {
		return g, ErrCellOccupied
	}

	// Create new game state (immutable)
	newGame := g

	// Apply move to board
	newGame.Board = g.Board.SetCell(row, col, g.CurrentPlayer.GetMark())

	// Increment move count
	newGame.MoveCount++

	// Check for win
	if CheckWin(newGame.Board, g.CurrentPlayer.GetMark()) {
		if g.CurrentPlayer == Player1 {
			newGame.State = Player1Won
		} else {
			newGame.State = Player2Won
		}
		return newGame, nil
	}

	// Check for draw
	if CheckDraw(newGame.Board) {
		newGame.State = Draw
		return newGame, nil
	}

	// Switch player
	newGame.CurrentPlayer = g.CurrentPlayer.Other()

	return newGame, nil
}

// Error types for move validation
var (
	ErrInvalidRange = &GameError{"Invalid position. Row and column must be between 0 and 2"}
	ErrCellOccupied = &GameError{"Position already occupied. Please choose an empty cell"}
)

// GameError represents a game-specific error
type GameError struct {
	Message string
}

func (e *GameError) Error() string {
	return e.Message
}
