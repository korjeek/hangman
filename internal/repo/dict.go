package repo

import (
	_ "embed"
	"hangman/internal/game"
)

//go:embed words.json
var words []byte

type Word struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

type Category struct {
	Name  game.Category `json:"name"`
	Words []Word        `json:"words"`
}

type Dictionary struct {
	Categories []Category `json:"categories"`
}
