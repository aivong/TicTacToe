package game

import "testing"

// TestCellConstants verifies Cell type constants
func TestCellConstants(t *testing.T) {
	tests := []struct {
		name     string
		cell     Cell
		expected int
	}{
		{"Empty cell value", Empty, 0},
		{"X cell value", X, 1},
		{"O cell value", O, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.cell) != tt.expected {
				t.Errorf("Cell value = %d, want %d", int(tt.cell), tt.expected)
			}
		})
	}
}

// TestCellString verifies Cell String() method
func TestCellString(t *testing.T) {
	tests := []struct {
		name     string
		cell     Cell
		expected string
	}{
		{"Empty cell string", Empty, " "},
		{"X cell string", X, "X"},
		{"O cell string", O, "O"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.String()
			if got != tt.expected {
				t.Errorf("Cell.String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestCellIsOccupied verifies Cell IsOccupied() method
func TestCellIsOccupied(t *testing.T) {
	tests := []struct {
		name     string
		cell     Cell
		expected bool
	}{
		{"Empty cell is not occupied", Empty, false},
		{"X cell is occupied", X, true},
		{"O cell is occupied", O, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.IsOccupied()
			if got != tt.expected {
				t.Errorf("Cell.IsOccupied() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestNewBoard verifies NewBoard creates an empty board
func TestNewBoard(t *testing.T) {
	board := NewBoard()

	// Check all cells are empty
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if board[row][col] != Empty {
				t.Errorf("NewBoard()[%d][%d] = %v, want Empty", row, col, board[row][col])
			}
		}
	}
}

// TestBoardGetCell verifies GetCell returns correct cell value
func TestBoardGetCell(t *testing.T) {
	board := Board{
		{X, O, Empty},
		{O, X, O},
		{Empty, Empty, X},
	}

	tests := []struct {
		name     string
		row      int
		col      int
		expected Cell
	}{
		{"Top-left corner (X)", 0, 0, X},
		{"Top-middle (O)", 0, 1, O},
		{"Top-right (Empty)", 0, 2, Empty},
		{"Center (X)", 1, 1, X},
		{"Bottom-right corner (X)", 2, 2, X},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := board.GetCell(tt.row, tt.col)
			if got != tt.expected {
				t.Errorf("GetCell(%d, %d) = %v, want %v", tt.row, tt.col, got, tt.expected)
			}
		})
	}
}

// TestBoardSetCell verifies SetCell returns new board with updated cell
func TestBoardSetCell(t *testing.T) {
	board := NewBoard()

	// Set a cell
	newBoard := board.SetCell(1, 1, X)

	// Original board should be unchanged (immutability)
	if board.GetCell(1, 1) != Empty {
		t.Errorf("Original board was modified, expected Empty at (1,1)")
	}

	// New board should have the updated cell
	if newBoard.GetCell(1, 1) != X {
		t.Errorf("SetCell(1, 1, X) did not update cell, got %v, want X", newBoard.GetCell(1, 1))
	}
}

// TestBoardIsCellEmpty verifies IsCellEmpty correctly identifies empty cells
func TestBoardIsCellEmpty(t *testing.T) {
	board := Board{
		{X, O, Empty},
		{Empty, X, O},
		{Empty, Empty, X},
	}

	tests := []struct {
		name     string
		row      int
		col      int
		expected bool
	}{
		{"Occupied cell (0,0) with X", 0, 0, false},
		{"Occupied cell (0,1) with O", 0, 1, false},
		{"Empty cell (0,2)", 0, 2, true},
		{"Empty cell (1,0)", 1, 0, true},
		{"Occupied cell (1,1) with X", 1, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := board.IsCellEmpty(tt.row, tt.col)
			if got != tt.expected {
				t.Errorf("IsCellEmpty(%d, %d) = %v, want %v", tt.row, tt.col, got, tt.expected)
			}
		})
	}
}

// TestBoardIsFull verifies IsFull correctly identifies full boards
func TestBoardIsFull(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected bool
	}{
		{
			name:     "Empty board",
			board:    NewBoard(),
			expected: false,
		},
		{
			name: "Partially filled board",
			board: Board{
				{X, O, X},
				{O, X, O},
				{Empty, Empty, Empty},
			},
			expected: false,
		},
		{
			name: "Fully filled board",
			board: Board{
				{X, O, X},
				{O, X, O},
				{O, X, O},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.IsFull()
			if got != tt.expected {
				t.Errorf("IsFull() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestBoardCountOccupied verifies CountOccupied returns correct count
func TestBoardCountOccupied(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected int
	}{
		{
			name:     "Empty board has 0 occupied cells",
			board:    NewBoard(),
			expected: 0,
		},
		{
			name: "Board with 3 occupied cells",
			board: Board{
				{X, O, X},
				{Empty, Empty, Empty},
				{Empty, Empty, Empty},
			},
			expected: 3,
		},
		{
			name: "Board with 5 occupied cells",
			board: Board{
				{X, O, X},
				{O, X, Empty},
				{Empty, Empty, Empty},
			},
			expected: 5,
		},
		{
			name: "Full board has 9 occupied cells",
			board: Board{
				{X, O, X},
				{O, X, O},
				{O, X, O},
			},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.CountOccupied()
			if got != tt.expected {
				t.Errorf("CountOccupied() = %d, want %d", got, tt.expected)
			}
		})
	}
}

// TestMakeMove verifies MakeMove applies move and switches players
func TestMakeMove(t *testing.T) {
	game := NewGame()

	// Make first move (Player1/X at position 1,1)
	game, err := game.MakeMove(1, 1)
	if err != nil {
		t.Fatalf("MakeMove(1, 1) returned error: %v", err)
	}

	// Verify board updated
	if game.Board.GetCell(1, 1) != X {
		t.Errorf("After move, board[1][1] = %v, want X", game.Board.GetCell(1, 1))
	}

	// Verify player switched
	if game.CurrentPlayer != Player2 {
		t.Errorf("After move, CurrentPlayer = %v, want Player2", game.CurrentPlayer)
	}

	// Verify move count incremented
	if game.MoveCount != 1 {
		t.Errorf("After move, MoveCount = %d, want 1", game.MoveCount)
	}
}

// TestMakeMoveInvalidPosition verifies errors on invalid positions
func TestMakeMoveInvalidPosition(t *testing.T) {
	game := NewGame()

	tests := []struct {
		name string
		row  int
		col  int
	}{
		{"negative row", -1, 1},
		{"negative col", 1, -1},
		{"row too large", 3, 1},
		{"col too large", 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := game.MakeMove(tt.row, tt.col)
			if err == nil {
				t.Errorf("MakeMove(%d, %d) should return error for invalid position", tt.row, tt.col)
			}
		})
	}
}

// TestMakeMoveOccupiedCell verifies error when cell already occupied
func TestMakeMoveOccupiedCell(t *testing.T) {
	game := NewGame()

	// Make first move
	game, _ = game.MakeMove(1, 1)

	// Try to make move at same position
	_, err := game.MakeMove(1, 1)
	if err == nil {
		t.Error("MakeMove on occupied cell should return error")
	}
}
