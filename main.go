package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/YOUR_USERNAME/tictactoe/game"
	"github.com/YOUR_USERNAME/tictactoe/validation"
)

func main() {
	fmt.Println("=== Tic-Tac-Toe ===")
	fmt.Println()

	g := game.NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	for g.State == game.InProgress {
		// Display board
		displayBoard(g.Board)

		// Display current player
		fmt.Printf("\n%s's turn\n", g.CurrentPlayer.Name())
		fmt.Print("Enter row and column (0-2), e.g., '1 1': ")

		// Read input
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		// Validate and parse input using validation package
		row, col, err := validation.ParseAndValidateInput(input)
		if err != nil {
			displayError(err)
			continue
		}

		// Make move
		newGame, err := g.MakeMove(row, col)
		if err != nil {
			displayError(err)
			continue
		}

		g = newGame
	}

	// Display final board
	displayBoard(g.Board)
	fmt.Println()

	// Display result
	switch g.State {
	case game.Player1Won:
		fmt.Println("🎉 Player 1 (X) wins!")
	case game.Player2Won:
		fmt.Println("🎉 Player 2 (O) wins!")
	case game.Draw:
		fmt.Println("It's a draw!")
	}
}

func displayBoard(board game.Board) {
	fmt.Println()
	fmt.Println("  0   1   2")
	for row := 0; row < 3; row++ {
		fmt.Printf("%d ", row)
		for col := 0; col < 3; col++ {
			cell := board.GetCell(row, col)
			fmt.Printf(" %s ", cell.String())
			if col < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if row < 2 {
			fmt.Println("  -----------")
		}
	}
	fmt.Println()
}

// displayError shows user-friendly error messages based on error type
func displayError(err error) {
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════╗")

	// Check for specific validation errors
	switch {
	case errors.Is(err, validation.ErrInvalidRange):
		fmt.Println("║  ❌ Invalid Position                      ║")
		fmt.Println("║                                            ║")
		fmt.Println("║  Row and column must be between 0 and 2    ║")
		fmt.Println("║  Example: '1 1' for center position        ║")
	case errors.Is(err, validation.ErrInvalidFormat):
		fmt.Println("║  ❌ Invalid Format                        ║")
		fmt.Println("║                                            ║")
		fmt.Println("║  Please enter numeric values only          ║")
		fmt.Println("║  Example: '0 2' or '1 1'                   ║")
	case errors.Is(err, validation.ErrIncompleteInput):
		fmt.Println("║  ❌ Incomplete Input                      ║")
		fmt.Println("║                                            ║")
		fmt.Println("║  Please enter two numbers separated by     ║")
		fmt.Println("║  space (row and column)                    ║")
	case errors.Is(err, game.ErrCellOccupied):
		fmt.Println("║  ❌ Cell Already Occupied                 ║")
		fmt.Println("║                                            ║")
		fmt.Println("║  That position is already taken            ║")
		fmt.Println("║  Please choose an empty cell               ║")
	default:
		// Generic error display
		fmt.Println("║  ❌ Error                                 ║")
		fmt.Println("║                                            ║")
		fmt.Printf("║  %s\n", err.Error())
	}

	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println()
}
