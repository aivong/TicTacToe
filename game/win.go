package game

// CheckWin determines if the specified player has won the game
// by achieving three marks in a row (horizontal, vertical, or diagonal)
// Complexity: O(1) - fixed 8 checks regardless of board size
func CheckWin(board Board, player Cell) bool {
	return checkRows(board, player) ||
		checkColumns(board, player) ||
		checkDiagonals(board, player)
}

// checkRows checks if player has won in any row
func checkRows(board Board, player Cell) bool {
	for row := 0; row < BOARD_SIZE; row++ {
		if board[row][0] == player &&
			board[row][1] == player &&
			board[row][2] == player {
			return true
		}
	}
	return false
}

// checkColumns checks if player has won in any column
func checkColumns(board Board, player Cell) bool {
	for col := 0; col < BOARD_SIZE; col++ {
		if board[0][col] == player &&
			board[1][col] == player &&
			board[2][col] == player {
			return true
		}
	}
	return false
}

// checkDiagonals checks if player has won in either diagonal
func checkDiagonals(board Board, player Cell) bool {
	// Main diagonal (top-left to bottom-right)
	mainDiag := board[0][0] == player &&
		board[1][1] == player &&
		board[2][2] == player

	// Anti-diagonal (top-right to bottom-left)
	antiDiag := board[0][2] == player &&
		board[1][1] == player &&
		board[2][0] == player

	return mainDiag || antiDiag
}

// CheckDraw determines if the game is a draw
// A draw occurs when the board is full and neither player has won
func CheckDraw(board Board) bool {
	// Not a draw if board isn't full
	if !board.IsFull() {
		return false
	}

	// Not a draw if either player has won
	if CheckWin(board, X) || CheckWin(board, O) {
		return false
	}

	return true
}
