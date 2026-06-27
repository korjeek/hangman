package formatter_test

import (
	"hangman/internal/app/formatter"
	"hangman/internal/domain/session"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatResult(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name     string
		hidden   string
		guessed  map[rune]bool
		status   session.Status
		expected string
	}

	tests := []TestCase{
		{
			name:   "all guessed",
			hidden: "test",
			guessed: map[rune]bool{
				't': true,
				'e': true,
				's': true,
			},
			status:   session.Win,
			expected: "test;POS",
		},
		{
			name:     "nothing guessed",
			hidden:   "test",
			guessed:  map[rune]bool{},
			status:   session.Lose,
			expected: "****;NEG",
		},
		{
			name:   "some guessed",
			hidden: "test",
			guessed: map[rune]bool{
				't': true,
			},
			status:   session.Lose,
			expected: "t**t;NEG",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			actual := formatter.FormatResult(test.hidden, test.guessed, test.status)
			assert.Equal(tt, test.expected, actual)
		})
	}
}
