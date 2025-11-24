package main

import (
	"testing"

	"github.com/YOUR_USERNAME/tictactoe/game"
)

// TestCompleteGameWithWinner simulates a complete game ending in a win
func TestCompleteGameWithWinner(t *testing.T) {
	g := game.NewGame()

	// Simulate game: X wins with diagonal
	// Move 1: X at (0,0)
	var err error
	g, err = g.MakeMove(0, 0)
	if err != nil {
		t.Fatalf("Move 1 failed: %v", err)
	}

	// Move 2: O at (0, 1)
	g, err = g.MakeMove(0, 1)
	if err != nil {
		t.Fatalf("Move 2 failed: %v", err)
	}

	// Move 3: X at (1,1)
	g, err = g.MakeMove(1, 1)
	if err != nil {
		t.Fatalf("Move 3 failed: %v", err)
	}

	// Move 4: O at (0,2)
	g, err = g.MakeMove(0, 2)
	if err != nil {
		t.Fatalf("Move 4 failed: %v", err)
	}

	// Move 5: X at (2,2) - wins with diagonal
	g, err = g.MakeMove(2, 2)
	if err != nil {
		t.Fatalf("Move 5 failed: %v", err)
	}

	// Verify X won
	if !game.CheckWin(g.Board, game.X) {
		t.Error("X should have won with diagonal")
	}

	// Verify game recognizes win
	if g.State != game.Player1Won {
		t.Errorf("Game state = %v, want Player1Won", g.State)
	}
}

// TestCompleteGameWithDraw simulates a complete game ending in a draw
func TestCompleteGameWithDraw(t *testing.T) {
	g := game.NewGame()

	// Simulate draw game
	moves := [][2]int{
		{0, 0}, // X
		{0, 1}, // O
		{0, 2}, // X
		{1, 1}, // O
		{1, 0}, // X
		{1, 2}, // O
		{2, 1}, // X
		{2, 0}, // O
		{2, 2}, // X - draw
	}

	var err error
	for i, move := range moves {
		g, err = g.MakeMove(move[0], move[1])
		if err != nil {
			t.Fatalf("Move %d at (%d,%d) failed: %v", i+1, move[0], move[1], err)
		}
	}

	// Verify it's a draw
	if !game.CheckDraw(g.Board) {
		t.Error("Game should be a draw")
	}

	// Verify game state
	if g.State != game.Draw {
		t.Errorf("Game state = %v, want Draw", g.State)
	}
}

// TestGameRejectsInvalidMoves verifies error handling
func TestGameRejectsInvalidMoves(t *testing.T) {
	g := game.NewGame()

	// Make valid move
	g, _ = g.MakeMove(1, 1)

	// Try to move on occupied cell
	_, err := g.MakeMove(1, 1)
	if err == nil {
		t.Error("Should reject move on occupied cell")
	}

	// Try invalid coordinates
	_, err = g.MakeMove(5, 5)
	if err == nil {
		t.Error("Should reject out-of-bounds move")
	}
}

// TestEnhancedErrorHandling tests the validation layer error handling
func TestEnhancedErrorHandling(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldError bool
		description string
	}{
		{
			name:        "Valid input accepted",
			input:       "1 1",
			shouldError: false,
			description: "Normal valid input should be accepted",
		},
		{
			name:        "Empty input rejected",
			input:       "",
			shouldError: true,
			description: "Empty input should trigger incomplete input error",
		},
		{
			name:        "Single number rejected",
			input:       "1",
			shouldError: true,
			description: "Incomplete input (single number) should be rejected",
		},
		{
			name:        "Non-numeric rejected",
			input:       "a b",
			shouldError: true,
			description: "Non-numeric input should trigger format error",
		},
		{
			name:        "Out of range rejected",
			input:       "5 5",
			shouldError: true,
			description: "Out-of-range coordinates should be rejected",
		},
		{
			name:        "Negative values rejected",
			input:       "-1 1",
			shouldError: true,
			description: "Negative coordinates should be rejected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This test documents expected behavior for validation layer
			// Actual integration with game will be implemented in T027
			t.Log(tt.description)
			// Placeholder for validation integration
		})
	}
}

// TestGameLoopErrorRecovery verifies game continues after invalid input
func TestGameLoopErrorRecovery(t *testing.T) {
	g := game.NewGame()

	// Make invalid move (out of range)
	_, err := g.MakeMove(10, 10)
	if err == nil {
		t.Fatal("Expected error for out-of-range move")
	}

	// Verify game state unchanged
	if g.CurrentPlayer != game.Player1 {
		t.Error("Player should not switch after invalid move")
	}
	if g.MoveCount != 0 {
		t.Error("Move count should not increment after invalid move")
	}

	// Valid move should work after error
	g, err = g.MakeMove(1, 1)
	if err != nil {
		t.Fatalf("Valid move after error failed: %v", err)
	}

	// Verify game state advanced
	if g.CurrentPlayer != game.Player2 {
		t.Error("Player should switch after valid move")
	}
	if g.MoveCount != 1 {
		t.Errorf("Move count = %d, want 1", g.MoveCount)
	}
}

// TestMultipleInvalidInputsBeforeValid verifies resilience
func TestMultipleInvalidInputsBeforeValid(t *testing.T) {
	g := game.NewGame()

	// Try multiple invalid moves
	invalidMoves := [][2]int{
		{-1, 0},  // Negative
		{3, 3},   // Out of range
		{10, 10}, // Way out of range
	}

	for _, move := range invalidMoves {
		_, err := g.MakeMove(move[0], move[1])
		if err == nil {
			t.Errorf("Move (%d, %d) should have failed", move[0], move[1])
		}

		// Verify state unchanged
		if g.CurrentPlayer != game.Player1 {
			t.Error("Player should remain Player1 after invalid moves")
		}
		if g.MoveCount != 0 {
			t.Error("Move count should remain 0 after invalid moves")
		}
	}

	// Finally make valid move
	g, err := g.MakeMove(0, 0)
	if err != nil {
		t.Fatalf("Valid move after multiple errors failed: %v", err)
	}

	// Verify game progressed correctly
	if g.Board.GetCell(0, 0) != game.X {
		t.Error("Board should have X at (0,0)")
	}
	if g.CurrentPlayer != game.Player2 {
		t.Error("Player should switch to Player2 after valid move")
	}
}
