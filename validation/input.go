// Package validation provides input validation functions for the tic-tac-toe game
package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Validation error types as sentinel errors
var (
	// ErrInvalidRange indicates coordinates are outside valid range [0-2]
	ErrInvalidRange = errors.New("Invalid position. Row and column must be between 0 and 2")

	// ErrInvalidFormat indicates input contains non-numeric values
	ErrInvalidFormat = errors.New("Invalid input format. Please enter numeric values only")

	// ErrIncompleteInput indicates input doesn't contain two numbers
	ErrIncompleteInput = errors.New("Incomplete input. Please enter two numbers separated by space")
)

const (
	// MinCoordinate is the minimum valid board coordinate
	MinCoordinate = 0
	// MaxCoordinate is the maximum valid board coordinate (3x3 board)
	MaxCoordinate = 2
)

// ValidateRange checks if row and column are within valid board range [0-2]
// Returns ErrInvalidRange if coordinates are out of bounds
func ValidateRange(row, col int) error {
	if row < MinCoordinate || row > MaxCoordinate {
		return ErrInvalidRange
	}
	if col < MinCoordinate || col > MaxCoordinate {
		return ErrInvalidRange
	}
	return nil
}

// ValidateNumeric parses a string into an integer
// Returns ErrInvalidFormat if the string is not a valid number
func ValidateNumeric(input string) (int, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, ErrInvalidFormat
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("%w: %q is not a valid number", ErrInvalidFormat, input)
	}

	return num, nil
}

// ValidateInputFormat validates and parses a complete input string (e.g., "1 2")
// Returns row, col, and an error if validation fails
// Checks for:
// - Exactly two numbers separated by whitespace
// - Both values are numeric
func ValidateInputFormat(input string) (int, int, error) {
	parts := strings.Fields(input)

	// Check for exactly two parts
	if len(parts) == 0 {
		return 0, 0, ErrIncompleteInput
	}
	if len(parts) == 1 {
		return 0, 0, ErrIncompleteInput
	}
	if len(parts) > 2 {
		return 0, 0, fmt.Errorf("%w: expected 2 numbers, got %d", ErrInvalidFormat, len(parts))
	}

	// Parse row
	row, err := ValidateNumeric(parts[0])
	if err != nil {
		return 0, 0, err
	}

	// Parse column
	col, err := ValidateNumeric(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return row, col, nil
}

// ParseAndValidateInput is the main validation pipeline
// It combines format validation and range validation
// Returns validated row and column, or an error describing what went wrong
func ParseAndValidateInput(input string) (int, int, error) {
	// Step 1: Validate format and parse
	row, col, err := ValidateInputFormat(input)
	if err != nil {
		return 0, 0, err
	}

	// Step 2: Validate range
	err = ValidateRange(row, col)
	if err != nil {
		return 0, 0, err
	}

	return row, col, nil
}
