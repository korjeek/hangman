package formatter

import (
	"hangman/internal/domain/session"
	"strings"
)

// FormatResult builds a string representation of the current game state.
// It masks unguessed letters with '*' and appends the win status ("POS" or "NEG").
func FormatResult(
	hidden string,
	guessed map[rune]bool,
	status session.Status,
) string {
	var result strings.Builder

	for _, char := range hidden {
		if _, ok := guessed[char]; ok {
			result.WriteRune(char)
		} else {
			result.WriteRune('*')
		}
	}

	res := "NEG"
	if status == session.Win {
		res = "POS"
	}

	return result.String() + ";" + res
}
