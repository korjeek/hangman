package words

import (
	"strings"
)

// Words represents a parsed pair of text strings containing a hidden secret word
// and a sequence of guessed characters, typically used for non-interactive execution setup.
type Words struct {
	hidden  string
	guessed string
}

// GetHidden returns the secret hidden word extracted from the source data format.
func (w *Words) GetHidden() string { return w.hidden }

// GetGuessed returns the string sequence of pre-configured letter guesses.
func (w *Words) GetGuessed() string { return w.guessed }

// IsEmpty reports whether the Words instance holds blank values for both its internal fields.
func (w *Words) IsEmpty() bool {
	return w.hidden == "" && w.guessed == ""
}

func (w *Words) UnmarshalText(text []byte) error {
	input := strings.ToLower(string(text))
	words := strings.Split(input, " ")

	if len(words) != 2 {
		return ErrUnmarshal{Words: words}
	}

	w.hidden = words[0]
	w.guessed = words[1]
	return nil
}
