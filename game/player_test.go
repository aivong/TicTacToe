package game

import "testing"

// TestPlayerConstants verifies Player constants
func TestPlayerConstants(t *testing.T) {
	if Player1 != Player(X) {
		t.Errorf("Player1 = %v, want Player(X)", Player1)
	}
	if Player2 != Player(O) {
		t.Errorf("Player2 = %v, want Player(O)", Player2)
	}
}

// TestPlayerGetMark verifies GetMark returns correct Cell
func TestPlayerGetMark(t *testing.T) {
	tests := []struct {
		name     string
		player   Player
		expected Cell
	}{
		{"Player1 has mark X", Player1, X},
		{"Player2 has mark O", Player2, O},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.player.GetMark()
			if got != tt.expected {
				t.Errorf("Player.GetMark() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestPlayerName verifies Name returns correct display name
func TestPlayerName(t *testing.T) {
	tests := []struct {
		name     string
		player   Player
		expected string
	}{
		{"Player1 name", Player1, "Player 1 (X)"},
		{"Player2 name", Player2, "Player 2 (O)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.player.Name()
			if got != tt.expected {
				t.Errorf("Player.Name() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestPlayerOther verifies Other returns opposite player
func TestPlayerOther(t *testing.T) {
	tests := []struct {
		name     string
		player   Player
		expected Player
	}{
		{"Player1's other is Player2", Player1, Player2},
		{"Player2's other is Player1", Player2, Player1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.player.Other()
			if got != tt.expected {
				t.Errorf("Player.Other() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestPlayerOtherIsSymmetric verifies Other is symmetric
func TestPlayerOtherIsSymmetric(t *testing.T) {
	// Player1.Other().Other() should return Player1
	if Player1.Other().Other() != Player1 {
		t.Error("Player1.Other().Other() != Player1, Other() is not symmetric")
	}

	// Player2.Other().Other() should return Player2
	if Player2.Other().Other() != Player2 {
		t.Error("Player2.Other().Other() != Player2, Other() is not symmetric")
	}
}
