package mode

import (
	"hangman/internal/app/formatter"
	"hangman/internal/domain/session"
	"hangman/internal/ui"
)

// NonInteractive handles the logic for executing a game session automatically
// based on a predefined sequence of letter guesses without human interaction.
type NonInteractive struct {
	session *session.Session
	result  string
	ui      ui.UserInterface
}

func NewNonInteractive(session *session.Session, result string, ui ui.UserInterface) *NonInteractive {
	return &NonInteractive{
		session: session,
		result:  result,
		ui:      ui,
	}
}

// Run processes all the pre-configured guesses sequentially, formats the final
// state of the game session, and writes it to the user interface output.
func (ni *NonInteractive) Run() error {
	for _, letter := range ni.result {
		ni.session.Guess(letter)
	}

	formatted := formatter.FormatResult(
		ni.session.GetHidden(),
		ni.session.GetGuessed(),
		ni.session.GetStatus(),
	)

	ni.ui.Println(formatted)
	return nil
}
