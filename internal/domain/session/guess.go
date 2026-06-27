package session

import (
	"strings"
)

// GuessResult represents the categorical evaluation outcome of a single letter guess attempt.
type GuessResult int

const (
	Correct GuessResult = iota
	Incorrect
	AlreadyGuessed
)

// Guess evaluates a single input rune against the hidden secret word, updates
// the set of processed characters, reduces remaining attempts on failure, and refreshes the game status.
func (s *Session) Guess(letter rune) GuessResult {
	if s.status != InProgress {
		return Incorrect
	}

	if s.guessedLetters[letter] {
		return AlreadyGuessed
	}

	s.guessedLetters[letter] = true

	result := Correct
	if !strings.ContainsRune(s.hidden, letter) {
		s.leftAttempts--
		result = Incorrect
	}

	s.updateStatus()
	return result
}
