package mode

import (
	"hangman/internal/domain/session"
	"hangman/internal/ui"
)

// Interactive handles the logic for playing the game in the terminal
// with real-time user input and visual feedback.
type Interactive struct {
	session *session.Session
	ui      ui.UserInterface
}

func NewInteractive(session *session.Session, ui ui.UserInterface) *Interactive {
	return &Interactive{
		session: session,
		ui:      ui,
	}
}

// Run executes the main interactive game loop, processing user turns
// until the game session is finished, and displays the final game status.
func (i *Interactive) Run() error {
	for i.session.GetStatus() == session.InProgress {
		i.ui.Clear()

		letter := i.ui.AskRune()
		res := i.session.Guess(letter)

		i.ui.PrintHangman(i.session.GetLeftAttempts())
		i.ui.PrintGuessResult(res, i.session.GetLeftAttempts())

		if i.session.GetLeftAttempts() == 3 && i.session.GetHint() != "" {
			i.ui.PrintHint(i.session.GetHint())
		}
	}

	st := i.session.GetStatus()
	i.ui.PrintGameStatus(st)
	return nil
}
