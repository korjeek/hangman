package words_test

import (
	"hangman/internal/domain/words"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalText(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name          string
		text          []byte
		expectedHide  string
		expectedGuess string
		expectedErr   error
	}

	tests := []TestCase{
		{
			name:        "empty",
			text:        []byte{},
			expectedErr: words.ErrUnmarshal{Words: []string{""}},
		},
		{
			name:        "one",
			text:        []byte("one"),
			expectedErr: words.ErrUnmarshal{Words: []string{"one"}},
		},
		{
			name:          "two",
			text:          []byte("one two"),
			expectedHide:  "one",
			expectedGuess: "two",
			expectedErr:   nil,
		},
		{
			name:        "three",
			text:        []byte("one two three"),
			expectedErr: words.ErrUnmarshal{Words: []string{"one", "two", "three"}},
		},
		{
			name:          "random size of letters",
			text:          []byte("OnE TWo"),
			expectedHide:  "one",
			expectedGuess: "two",
			expectedErr:   nil,
		},
		{
			name:          "two but empty",
			text:          []byte(" "),
			expectedHide:  "",
			expectedGuess: "",
			expectedErr:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			var w words.Words
			err := w.UnmarshalText(test.text)

			if test.expectedErr != nil {
				var targetErr words.ErrUnmarshal
				if assert.ErrorAs(tt, err, &targetErr) {
					assert.Equal(tt, test.expectedErr.(words.ErrUnmarshal).Words, targetErr.Words)
				}
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, test.expectedHide, w.GetHidden())
			assert.Equal(tt, test.expectedGuess, w.GetGuessed())
		})
	}
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name     string
		text     []byte
		expected bool
	}

	tests := []TestCase{
		{
			name:     "empty",
			text:     []byte(" "),
			expected: true,
		},
		{
			name:     "hidden only",
			text:     []byte("one "),
			expected: false,
		},
		{
			name:     "guessed only",
			text:     []byte(" two"),
			expected: false,
		},
		{
			name:     "full",
			text:     []byte("one two"),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			var w words.Words
			err := w.UnmarshalText(test.text)
			assert.NoError(tt, err)

			actual := w.IsEmpty()
			assert.Equal(tt, test.expected, actual)
		})
	}
}
