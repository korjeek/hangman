package ctg_test

import (
	"hangman/internal/domain/ctg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalText(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name     string
		text     []byte
		expected ctg.Category
		error    bool
	}

	tests := []TestCase{
		{
			name:  "empty",
			text:  []byte{},
			error: true,
		},
		{
			name:     "animals",
			text:     []byte("animals"),
			expected: ctg.Animals,
			error:    false,
		},
		{
			name:     "countries",
			text:     []byte("countries"),
			expected: ctg.Countries,
			error:    false,
		},
		{
			name:     "nature",
			text:     []byte("nature"),
			expected: ctg.Nature,
			error:    false,
		},
		{
			name:     "food",
			text:     []byte("food"),
			expected: ctg.Food,
			error:    false,
		},
		{
			name:  "non-existent",
			text:  []byte("non-existent"),
			error: true,
		},
		{
			name:     "random size of letters",
			text:     []byte("AnImALS"),
			expected: ctg.Animals,
			error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			var c ctg.Category
			err := c.UnmarshalText(test.text)

			if test.error {
				assert.Error(tt, err)
			} else {
				assert.NoError(tt, err)
				assert.Equal(tt, test.expected, c)
			}
		})
	}
}
