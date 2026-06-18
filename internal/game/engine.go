package game

import "strings"

type GuessResult int

const (
	Correct GuessResult = iota
	Incorrect
	AlreadyGuessed
)

type Status int

const (
	Win Status = iota
	Lose
	InProgress
)

func Guess(session *Session, letter rune) GuessResult {
	if !strings.ContainsRune(session.hidden, letter) {
		session.leftAttempts--
		return Incorrect
	}

	if _, ok := session.guessedLetters[letter]; ok {
		return AlreadyGuessed
	}

	session.guessedLetters[letter] = true

	return Correct
}

func GetStatus(session *Session) Status {
	if session.leftAttempts <= 0 {
		return Lose
	}

	for _, char := range session.hidden {
		if !session.guessedLetters[char] {
			return InProgress
		}
	}

	return Win
}
