package ui_test

import (
	"bytes"
	"hangman/internal/domain/session"
	"hangman/internal/ui"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleUI_Clear(t *testing.T) {
	t.Parallel()

	outputBuf := &bytes.Buffer{}
	console := ui.New(nil, outputBuf, nil)

	console.Clear()

	assert.Equal(t, "\033[H\033[2J\n", outputBuf.String())
}

func TestConsoleUI_AskRune(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name           string
		inputData      string
		expectedRune   rune
		expectedOutput string
	}

	tests := []TestCase{
		{
			name:           "success single letter",
			inputData:      "a\n",
			expectedRune:   'a',
			expectedOutput: "enter the suggested letter > \n",
		},
		{
			name:         "validation retry on multiple letters",
			inputData:    "abc\nf\n",
			expectedRune: 'f',
			expectedOutput: "enter the suggested letter > \n" +
				"you should enter only one letter!\n" +
				"enter the suggested letter > \n",
		},
		{
			name:         "validation retry on empty input",
			inputData:    "\nf\n",
			expectedRune: 'f',
			expectedOutput: "enter the suggested letter > \n" +
				"you should enter only one letter!\n" +
				"enter the suggested letter > \n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			inputBuf := bytes.NewBufferString(test.inputData)
			outputBuf := &bytes.Buffer{}
			console := ui.New(inputBuf, outputBuf, nil)

			actualRune := console.AskRune()

			assert.Equal(tt, test.expectedRune, actualRune)
			assert.Equal(tt, test.expectedOutput, outputBuf.String())
		})
	}
}

func TestConsoleUI_PrintGuessResult(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name           string
		result         session.GuessResult
		leftAttempts   int
		expectedOutput string
	}

	tests := []TestCase{
		{
			name:           "correct guess",
			result:         session.Correct,
			leftAttempts:   5,
			expectedOutput: "the letter is guessed\nattempts left > 5\n",
		},
		{
			name:           "incorrect guess",
			result:         session.Incorrect,
			leftAttempts:   3,
			expectedOutput: "this letter is not in the hidden word\nattempts left > 3\n",
		},
		{
			name:           "already guessed",
			result:         session.AlreadyGuessed,
			leftAttempts:   4,
			expectedOutput: "this letter is already guessed\nattempts left > 4\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			outputBuf := &bytes.Buffer{}
			console := ui.New(nil, outputBuf, nil)

			console.PrintGuessResult(test.result, test.leftAttempts)

			assert.Equal(tt, test.expectedOutput, outputBuf.String())
		})
	}
}

func TestConsoleUI_PrintGameStatus(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name           string
		status         session.Status
		expectedOutput string
	}

	tests := []TestCase{
		{
			name:           "win status",
			status:         session.Win,
			expectedOutput: "you win, congratulations!",
		},
		{
			name:           "lose status",
			status:         session.Lose,
			expectedOutput: "you lose, better luck next time!",
		},
		{
			name:           "in progress status",
			status:         session.InProgress,
			expectedOutput: "game still in progress...",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			outputBuf := &bytes.Buffer{}
			console := ui.New(nil, outputBuf, nil)

			console.PrintGameStatus(test.status)

			assert.Equal(tt, test.expectedOutput, outputBuf.String())
		})
	}
}

func TestConsoleUI_PrintHint(t *testing.T) {
	t.Parallel()

	outputBuf := &bytes.Buffer{}
	console := ui.New(nil, outputBuf, nil)

	console.PrintHint("it flies")

	assert.Equal(t, "hint > it flies\n", outputBuf.String())
}

func TestConsoleUI_PrintHangman(t *testing.T) {
	t.Parallel()

	fakeStates := []string{"[state 0]", "[state 1]", "[state 2]"}

	type TestCase struct {
		name           string
		leftAttempts   int
		expectedOutput string
	}

	tests := []TestCase{
		{
			name:           "exact index",
			leftAttempts:   1,
			expectedOutput: "[state 1]\n",
		},
		{
			name:           "negative bounds protection",
			leftAttempts:   -5,
			expectedOutput: "[state 0]\n",
		},
		{
			name:           "overflow bounds protection",
			leftAttempts:   10,
			expectedOutput: "[state 2]\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			outputBuf := &bytes.Buffer{}
			console := ui.New(nil, outputBuf, fakeStates)

			console.PrintHangman(test.leftAttempts)

			assert.Equal(tt, test.expectedOutput, outputBuf.String())
		})
	}
}
