package game

import "testing"

// TestCheckWinHorizontal verifies horizontal win detection
func TestCheckWinHorizontal(t *testing.T) {
	tests := []struct {
		name   string
		board  Board
		player Cell
		want   bool
	}{
		{
			name: "Row 0 win for X",
			board: Board{
				{X, X, X},
				{O, O, Empty},
				{Empty, Empty, Empty},
			},
			player: X,
			want:   true,
		},
		{
			name: "Row 1 win for O",
			board: Board{
				{X, X, Empty},
				{O, O, O},
				{X, Empty, Empty},
			},
			player: O,
			want:   true,
		},
		{
			name: "Row 2 win for X",
			board: Board{
				{O, O, Empty},
				{Empty, O, Empty},
				{X, X, X},
			},
			player: X,
			want:   true,
		},
		{
			name: "No horizontal win",
			board: Board{
				{X, O, X},
				{O, X, O},
				{O, X, O},
			},
			player: X,
			want:   false,
		},
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

// TestCheckWinVertical verifies vertical win detection
func TestCheckWinVertical(t *testing.T) {
	tests := []struct {
		name   string
		board  Board
		player Cell
		want   bool
	}{
		{
			name: "Column 0 win for X",
			board: Board{
				{X, O, O},
				{X, O, Empty},
				{X, Empty, Empty},
			},
			player: X,
			want:   true,
		},
		{
			name: "Column 1 win for O",
			board: Board{
				{X, O, X},
				{Empty, O, X},
				{Empty, O, Empty},
			},
			player: O,
			want:   true,
		},
		{
			name: "Column 2 win for X",
			board: Board{
				{O, O, X},
				{O, Empty, X},
				{Empty, Empty, X},
			},
			player: X,
			want:   true,
		},
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

// TestCheckWinDiagonal verifies diagonal win detection
func TestCheckWinDiagonal(t *testing.T) {
	tests := []struct {
		name   string
		board  Board
		player Cell
		want   bool
	}{
		{
			name: "Main diagonal win for X",
			board: Board{
				{X, O, O},
				{O, X, Empty},
				{Empty, Empty, X},
			},
			player: X,
			want:   true,
		},
		{
			name: "Anti-diagonal win for O",
			board: Board{
				{X, X, O},
				{X, O, Empty},
				{O, Empty, Empty},
			},
			player: O,
			want:   true,
		},
		{
			name: "No diagonal win",
			board: Board{
				{X, O, X},
				{O, O, X},
				{O, X, O},
			},
			player: X,
			want:   false,
		},
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

// TestCheckWinEmptyBoard verifies no win on empty board
func TestCheckWinEmptyBoard(t *testing.T) {
	board := NewBoard()

	if CheckWin(board, X) {
		t.Error("CheckWin(empty board, X) = true, want false")
	}
	if CheckWin(board, O) {
		t.Error("CheckWin(empty board, O) = true, want false")
	}
}

// TestCheckWinPartialBoard verifies no false positives on partial boards
func TestCheckWinPartialBoard(t *testing.T) {
	tests := []struct {
		name   string
		board  Board
		player Cell
	}{
		{
			name: "Two in a row, not three",
			board: Board{
				{X, X, O},
				{O, O, X},
				{Empty, Empty, Empty},
			},
			player: X,
		},
		{
			name: "Two in column, not three",
			board: Board{
				{X, O, Empty},
				{X, O, Empty},
				{Empty, Empty, Empty},
			},
			player: X,
		},
		{
			name: "Two in diagonal, not three",
			board: Board{
				{X, O, Empty},
				{O, X, Empty},
				{Empty, Empty, O},
			},
			player: X,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CheckWin(tt.board, tt.player) {
				t.Errorf("CheckWin() = true, want false (partial match should not win)")
			}
		})
	}
}

// TestCheckDrawFullBoardNoWinner verifies draw detection on full board
func TestCheckDrawFullBoardNoWinner(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  bool
	}{
		{
			name: "Full board with no winner is draw",
			board: Board{
				{X, O, X},
				{O, X, O},
				{O, X, O},
			},
			want: true,
		},
		{
			name: "Another draw pattern",
			board: Board{
				{X, X, O},
				{O, O, X},
				{X, O, X},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckDraw(tt.board)
			if got != tt.want {
				t.Errorf("CheckDraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestCheckDrawPartialBoard verifies no draw on partial board
func TestCheckDrawPartialBoard(t *testing.T) {
	tests := []struct {
		name  string
		board Board
	}{
		{
			name:  "Empty board is not draw",
			board: NewBoard(),
		},
		{
			name: "Partially filled board is not draw",
			board: Board{
				{X, O, X},
				{O, X, O},
				{Empty, Empty, Empty},
			},
		},
		{
			name: "Almost full board is not draw",
			board: Board{
				{X, O, X},
				{O, X, O},
				{O, X, Empty},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CheckDraw(tt.board) {
				t.Error("CheckDraw() = true, want false (board not full)")
			}
		})
	}
}

// TestCheckDrawWithWinner verifies no draw when there's a winner
func TestCheckDrawWithWinner(t *testing.T) {
	tests := []struct {
		name  string
		board Board
	}{
		{
			name: "Full board with horizontal win is not draw",
			board: Board{
				{X, X, X},
				{O, O, X},
				{O, X, O},
			},
		},
		{
			name: "Full board with diagonal win is not draw",
			board: Board{
				{X, O, O},
				{O, X, O},
				{X, O, X},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if CheckDraw(tt.board) {
				t.Error("CheckDraw() = true, want false (there's a winner)")
			}
		})
	}
}
