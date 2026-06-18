package game

type Session struct {
	hidden         string
	guessedLetters map[rune]bool
	leftAttempts   int
}

func NewSession(hidden string, attempts int) (*Session, error) {
	if hidden == "" {
		return nil, &ValidationError{
			Message: "hidden word must not be empty",
		}
	}

	if attempts <= 0 {
		return nil, &ValidationError{
			Message: "attempts must be greater than zero",
		}
	}

	sess := &Session{
		hidden:         hidden,
		guessedLetters: make(map[rune]bool),
		leftAttempts:   attempts,
	}

	return sess, nil
}

func (s *Session) Hidden() string {
	return s.hidden
}

func (s *Session) Guessed() map[rune]bool {
	return s.guessedLetters
}
