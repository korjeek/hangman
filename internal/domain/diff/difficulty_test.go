package diff_test

import (
	"hangman/internal/domain/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalText(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name     string
		text     []byte
		expected diff.Difficulty
		error    bool
	}

	tests := []TestCase{
		{
			name:  "empty",
			text:  []byte{},
			error: true,
		},
		{
			name:     "easy",
			text:     []byte("easy"),
			expected: diff.Easy,
			error:    false,
		},
		{
			name:     "normal",
			text:     []byte("normal"),
			expected: diff.Normal,
			error:    false,
		},
		{
			name:     "hard",
			text:     []byte("hard"),
			expected: diff.Hard,
			error:    false,
		},
		{
			name:  "non-existent",
			text:  []byte("non-existent"),
			error: true,
		},
		{
			name:     "random size of letters",
			text:     []byte("EaSY"),
			expected: diff.Easy,
			error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			var d diff.Difficulty
			err := d.UnmarshalText(test.text)

			if test.error {
				assert.Error(tt, err)
			} else {
				assert.NoError(tt, err)
				assert.Equal(tt, test.expected, d)
			}
		})
	}
}

func TestGetMaxAttempts(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name     string
		diff     diff.Difficulty
		expected int
	}

	tests := []TestCase{
		{
			name:     "Easy",
			diff:     diff.Easy,
			expected: 10,
		},
		{
			name:     "Normal",
			diff:     diff.Normal,
			expected: 7,
		},
		{
			name:     "Hard",
			diff:     diff.Hard,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			att := test.diff.GetMaxAttempts()
			assert.Equal(tt, test.expected, att)
		})
	}
}
