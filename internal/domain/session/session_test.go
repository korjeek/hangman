package session_test

import (
	"testing"

	"hangman/internal/domain/session"

	"github.com/stretchr/testify/assert"
)

func TestSession_Guess(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name           string
		setup          func(t *testing.T) *session.Session
		rune           rune
		expectedResult session.GuessResult
		expectedStatus session.Status
		expectedLeft   int
	}

	tests := []TestCase{
		{
			name: "not in progress",
			setup: func(t *testing.T) *session.Session {
				sess, err := session.NewSession("test", 1)
				assert.NoError(t, err)
				sess.Guess('z')
				return sess
			},
			rune:           'a',
			expectedResult: session.Incorrect,
			expectedStatus: session.Lose,
			expectedLeft:   0,
		},
		{
			name: "already guessed",
			setup: func(t *testing.T) *session.Session {
				sess, err := session.NewSession("test", 5)
				assert.NoError(t, err)
				sess.Guess('t')
				return sess
			},
			rune:           't',
			expectedResult: session.AlreadyGuessed,
			expectedStatus: session.InProgress,
			expectedLeft:   5,
		},
		{
			name: "correct guess - game continues",
			setup: func(t *testing.T) *session.Session {
				sess, err := session.NewSession("test", 5)
				assert.NoError(t, err)
				return sess
			},
			rune:           'e',
			expectedResult: session.Correct,
			expectedStatus: session.InProgress,
			expectedLeft:   5,
		},
		{
			name: "incorrect guess - lose attempt",
			setup: func(t *testing.T) *session.Session {
				sess, err := session.NewSession("test", 5)
				assert.NoError(t, err)
				return sess
			},
			rune:           'z',
			expectedResult: session.Incorrect,
			expectedStatus: session.InProgress,
			expectedLeft:   4,
		},
		{
			name: "correct guess - win game",
			setup: func(t *testing.T) *session.Session {
				sess, err := session.NewSession("test", 5)
				assert.NoError(t, err)
				sess.Guess('t')
				sess.Guess('s')
				return sess
			},
			rune:           'e',
			expectedResult: session.Correct,
			expectedStatus: session.Win,
			expectedLeft:   5,
		},
		{
			name: "incorrect guess - lose game",
			setup: func(t *testing.T) *session.Session {
				sess, err := session.NewSession("test", 1)
				assert.NoError(t, err)
				return sess
			},
			rune:           'z',
			expectedResult: session.Incorrect,
			expectedStatus: session.Lose,
			expectedLeft:   0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			sess := test.setup(tt)
			actual := sess.Guess(test.rune)

			assert.Equal(tt, test.expectedResult, actual)
			assert.Equal(tt, test.expectedStatus, sess.GetStatus())
			assert.Equal(tt, test.expectedLeft, sess.GetLeftAttempts())
		})
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name             string
		hidden           string
		attempts         int
		opts             []session.Option
		expectedHidden   string
		expectedAttempts int
		expectedStatus   session.Status
		expectedHint     string
		expectedErr      error
	}

	tests := []TestCase{
		{
			name:             "success",
			hidden:           "test",
			attempts:         5,
			opts:             nil,
			expectedHidden:   "test",
			expectedAttempts: 5,
			expectedStatus:   session.InProgress,
			expectedHint:     "",
			expectedErr:      nil,
		},
		{
			name:             "success with opts",
			hidden:           "test",
			attempts:         5,
			opts:             []session.Option{session.WithHint("hint")},
			expectedHidden:   "test",
			expectedAttempts: 5,
			expectedStatus:   session.InProgress,
			expectedHint:     "hint",
			expectedErr:      nil,
		},
		{
			name:        "empty hidden",
			hidden:      "",
			attempts:    5,
			expectedErr: session.ErrValidation{},
		},
		{
			name:        "attempts less or equals 0",
			hidden:      "test",
			attempts:    0,
			expectedErr: session.ErrValidation{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			actual, err := session.NewSession(test.hidden, test.attempts, test.opts...)

			if test.expectedErr != nil {
				assert.Nil(tt, actual)
				assert.ErrorAs(tt, err, &test.expectedErr)
				return
			}

			assert.NoError(tt, err)
			assert.NotNil(tt, actual)
			assert.Equal(tt, test.expectedHidden, actual.GetHidden())
			assert.NotNil(tt, actual.GetGuessed())
			assert.Equal(tt, test.expectedAttempts, actual.GetLeftAttempts())
			assert.Equal(tt, test.expectedStatus, actual.GetStatus())
			assert.Equal(tt, test.expectedHint, actual.GetHint())
		})
	}
}
