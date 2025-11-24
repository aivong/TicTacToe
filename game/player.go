package game

// Player represents one of the two players in the game
type Player Cell

const (
	// Player1 uses mark X and goes first
	Player1 = Player(X)
	// Player2 uses mark O and goes second
	Player2 = Player(O)
)

// GetMark returns the Cell mark (X or O) for this player
func (p Player) GetMark() Cell {
	return Cell(p)
}

// Name returns the display name for this player
func (p Player) Name() string {
	switch p {
	case Player1:
		return "Player 1 (X)"
	case Player2:
		return "Player 2 (O)"
	default:
		return "Unknown Player"
	}
}

// Other returns the opposite player
func (p Player) Other() Player {
	switch p {
	case Player1:
		return Player2
	case Player2:
		return Player1
	default:
		return Player1 // Fallback, should not happen
	}
}
