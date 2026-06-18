package game

import (
	"errors"
	"strings"
)

type Words struct {
	hidden  string
	guessed string
}

func (w *Words) Hidden() string { return w.hidden }

func (w *Words) Guessed() string { return w.guessed }

func (w *Words) IsEmpty() bool {
	return w.hidden == "" && w.guessed == ""
}

func (w *Words) UnmarshalText(text []byte) error {
	input := strings.ToLower(string(text))
	words := strings.Split(input, " ")

	if len(words) != 2 {
		return errors.New("invalid words format")
	}

	w.hidden = words[0]
	w.guessed = words[1]
	return nil
}
