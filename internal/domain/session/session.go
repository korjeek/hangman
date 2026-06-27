package session

import (
	"github.com/google/uuid"
)

// Session represents an isolated, stateful instance of a running hangman game match.
type Session struct {
	id             uuid.UUID
	hidden         string
	guessedLetters map[rune]bool
	leftAttempts   int
	status         Status

	hint string
}

type Option func(*Session)

func NewSession(hidden string, attempts int, opts ...Option) (*Session, error) {
	if hidden == "" {
		return nil, ErrValidation{
			sessionId: uuid.Nil,
			msg:       "hidden session id should not be empty",
		}
	}

	if attempts <= 0 {
		return nil, ErrValidation{
			sessionId: uuid.Nil,
			msg:       "attempts should bigger than 0",
		}
	}

	sess := &Session{
		id:             uuid.New(),
		hidden:         hidden,
		guessedLetters: make(map[rune]bool),
		leftAttempts:   attempts,
		status:         InProgress,
	}

	for _, opt := range opts {
		opt(sess)
	}

	return sess, nil
}

// Options

// WithHint applies a functional configuration choice to populate a clue string into the target Session instance.
func WithHint(hint string) Option {
	return func(s *Session) {
		s.hint = hint
	}
}

// Getters

// GetHidden returns the secret hidden word bound to the current game session.
func (s *Session) GetHidden() string {
	return s.hidden
}

// GetGuessed returns a lookup map containing all unique letters processed as turn choices so far.
func (s *Session) GetGuessed() map[rune]bool {
	return s.guessedLetters
}

// GetLeftAttempts returns the total number of remaining wrong letter guess opportunities allowed.
func (s *Session) GetLeftAttempts() int {
	return s.leftAttempts
}

// GetHint returns the descriptive clue text string bound to the current game match session.
func (s *Session) GetHint() string {
	return s.hint
}

// GetStatus returns the current finite state machine execution state of the game session match.
func (s *Session) GetStatus() Status {
	return s.status
}
