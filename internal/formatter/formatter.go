package formatter

import (
	"hangman/internal/game"
	"strings"
)

func FormatResult(
	hidden string,
	guessed map[rune]bool,
	status game.Status,
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
	if status == game.Win {
		res = "POS"
	}

	return result.String() + ";" + res
}
