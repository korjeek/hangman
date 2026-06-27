package session

// Status represents the current execution state of a game match session.
type Status int

const (
	Win Status = iota
	Lose
	InProgress
)

func (s *Session) updateStatus() {
	if s.leftAttempts <= 0 {
		s.status = Lose
		return
	}

	for _, char := range s.hidden {
		if !s.guessedLetters[char] {
			s.status = InProgress
			return
		}
	}

	s.status = Win
}
