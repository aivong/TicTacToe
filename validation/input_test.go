package validation

import (
	"errors"
	"testing"
)

// TestValidateRange verifies range validation for board coordinates
func TestValidateRange(t *testing.T) {
	tests := []struct {
		name      string
		row       int
		col       int
		wantError bool
		errType   error
	}{
		// Valid cases
		{"Valid top-left (0,0)", 0, 0, false, nil},
		{"Valid center (1,1)", 1, 1, false, nil},
		{"Valid bottom-right (2,2)", 2, 2, false, nil},
		{"Valid all corners", 0, 2, false, nil},

		// Invalid cases - negative
		{"Negative row", -1, 1, true, ErrInvalidRange},
		{"Negative col", 1, -1, true, ErrInvalidRange},
		{"Both negative", -5, -3, true, ErrInvalidRange},

		// Invalid cases - too large
		{"Row too large", 3, 1, true, ErrInvalidRange},
		{"Col too large", 1, 3, true, ErrInvalidRange},
		{"Both too large", 10, 10, true, ErrInvalidRange},

		// Edge cases
		{"Large negative row", -999, 1, true, ErrInvalidRange},
		{"Large positive col", 1, 999, true, ErrInvalidRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRange(tt.row, tt.col)
			if tt.wantError {
				if err == nil {
					t.Errorf("ValidateRange(%d, %d) expected error, got nil", tt.row, tt.col)
				}
				if tt.errType != nil && !errors.Is(err, tt.errType) {
					t.Errorf("ValidateRange(%d, %d) error = %v, want %v", tt.row, tt.col, err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateRange(%d, %d) unexpected error: %v", tt.row, tt.col, err)
				}
			}
		})
	}
}

// TestValidateNumeric verifies numeric input parsing
func TestValidateNumeric(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errType   error
	}{
		// Valid cases
		{"Valid single digit", "1", false, nil},
		{"Valid zero", "0", false, nil},
		{"Valid two", "2", false, nil},

		// Invalid cases - non-numeric
		{"Letter", "a", true, ErrInvalidFormat},
		{"Word", "one", true, ErrInvalidFormat},
		{"Special char", "@", true, ErrInvalidFormat},
		{"Mixed alphanumeric", "1a", true, ErrInvalidFormat},

		// Edge cases
		{"Empty string", "", true, ErrInvalidFormat},
		{"Whitespace only", " ", true, ErrInvalidFormat},
		{"Multiple digits", "12", false, nil},        // Should parse but fail range check later
		{"Negative number string", "-1", false, nil}, // Parses but fails range
		{"Float", "1.5", true, ErrInvalidFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateNumeric(tt.input)
			if tt.wantError {
				if err == nil {
					t.Errorf("ValidateNumeric(%q) expected error, got nil", tt.input)
				}
				if tt.errType != nil && !errors.Is(err, tt.errType) {
					t.Errorf("ValidateNumeric(%q) error = %v, want %v", tt.input, err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateNumeric(%q) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

// TestValidateInputFormat verifies full input string format validation
func TestValidateInputFormat(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errType   error
	}{
		// Valid cases
		{"Valid single space", "1 1", false, nil},
		{"Valid multiple spaces", "1  1", false, nil},
		{"Valid with tab", "1\t1", false, nil},
		{"Valid zeros", "0 0", false, nil},

		// Invalid cases - incomplete
		{"Single number", "1", true, ErrIncompleteInput},
		{"Empty string", "", true, ErrIncompleteInput},
		{"Only spaces", "   ", true, ErrIncompleteInput},

		// Invalid cases - too many values
		{"Three numbers", "1 1 1", true, ErrInvalidFormat},
		{"Four numbers", "0 1 2 3", true, ErrInvalidFormat},

		// Invalid cases - non-numeric
		{"Letters only", "a b", true, ErrInvalidFormat},
		{"Mixed valid and invalid", "1 a", true, ErrInvalidFormat},
		{"Special characters", "@ #", true, ErrInvalidFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := ValidateInputFormat(tt.input)
			if tt.wantError {
				if err == nil {
					t.Errorf("ValidateInputFormat(%q) expected error, got nil", tt.input)
				}
				if tt.errType != nil && !errors.Is(err, tt.errType) {
					t.Errorf("ValidateInputFormat(%q) error = %v, want %v", tt.input, err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateInputFormat(%q) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

// TestValidateInputFormatReturnValues verifies correct row/col extraction
func TestValidateInputFormatReturnValues(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRow int
		expectedCol int
	}{
		{"Standard input", "1 2", 1, 2},
		{"Zero values", "0 0", 0, 0},
		{"Max values", "2 2", 2, 2},
		{"Mixed values", "0 2", 0, 2},
		{"Extra whitespace", "1  2", 1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row, col, err := ValidateInputFormat(tt.input)
			if err != nil {
				t.Fatalf("ValidateInputFormat(%q) unexpected error: %v", tt.input, err)
			}
			if row != tt.expectedRow {
				t.Errorf("ValidateInputFormat(%q) row = %d, want %d", tt.input, row, tt.expectedRow)
			}
			if col != tt.expectedCol {
				t.Errorf("ValidateInputFormat(%q) col = %d, want %d", tt.input, col, tt.expectedCol)
			}
		})
	}
}

// TestParseAndValidateInput verifies full input validation pipeline
func TestParseAndValidateInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errType   error
	}{
		// Valid cases
		{"Valid input 0 0", "0 0", false, nil},
		{"Valid input 1 1", "1 1", false, nil},
		{"Valid input 2 2", "2 2", false, nil},

		// Invalid format
		{"Empty input", "", true, ErrIncompleteInput},
		{"Single value", "1", true, ErrIncompleteInput},
		{"Non-numeric", "a b", true, ErrInvalidFormat},

		// Invalid range
		{"Negative row", "-1 1", true, ErrInvalidRange},
		{"Negative col", "1 -1", true, ErrInvalidRange},
		{"Row too large", "3 1", true, ErrInvalidRange},
		{"Col too large", "1 3", true, ErrInvalidRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := ParseAndValidateInput(tt.input)
			if tt.wantError {
				if err == nil {
					t.Errorf("ParseAndValidateInput(%q) expected error, got nil", tt.input)
				}
				if tt.errType != nil && !errors.Is(err, tt.errType) {
					t.Errorf("ParseAndValidateInput(%q) error = %v, want %v", tt.input, err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("ParseAndValidateInput(%q) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

// TestErrorMessages verifies error messages are descriptive
func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name          string
		err           error
		shouldContain string
	}{
		{"Range error message", ErrInvalidRange, "0 and 2"},
		{"Format error message", ErrInvalidFormat, "numeric"},
		{"Incomplete error message", ErrIncompleteInput, "two numbers"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.err.Error()
			if msg == "" {
				t.Errorf("%s has empty error message", tt.name)
			}
			// Basic check that error has some descriptive text
			if len(msg) < 10 {
				t.Errorf("%s error message too short: %q", tt.name, msg)
			}
		})
	}
}
