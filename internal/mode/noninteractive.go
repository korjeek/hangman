package mode

import (
	"fmt"
	"hangman/internal/formatter"
	"hangman/internal/game"
)

type NonInteractive struct {
	session *game.Session
	result  string
}

func NewNonInteractive(session *game.Session, result string) *NonInteractive {
	return &NonInteractive{
		session: session,
		result:  result,
	}
}

func (ni *NonInteractive) Run() error {
	for _, letter := range ni.result {
		game.Guess(ni.session, letter)
	}

	output := formatter.FormatResult(
		ni.session.Hidden(),
		ni.session.Guessed(),
		game.GetStatus(ni.session),
	)

	fmt.Println(output)
	return nil
}
