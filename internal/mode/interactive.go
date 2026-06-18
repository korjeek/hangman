package mode

import (
	"fmt"
	"hangman/internal/game"
	"hangman/internal/ui"
)

type Interactive struct {
	session *game.Session
}

func NewInteractive(session *game.Session) *Interactive {
	return &Interactive{session: session}
}

func (i *Interactive) Run() error {
	csl := ui.New()

	for {
		status := game.GetStatus(i.session)

		if status != game.InProgress {
			fmt.Println("Игра окончена!")
			break
		}

		letter := csl.AskRune("enter the suggested letter")
		res := game.Guess(i.session, letter)
		csl.PrintGuessResult(res)
	}
}
