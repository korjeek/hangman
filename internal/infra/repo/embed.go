package repo

import (
	"embed"
	"encoding/json"
	"hangman/internal/domain/dict"
	"io/fs"
)

//go:embed words.json
var dictFS embed.FS

const dictName = "words.json"

// Load reads the embedded words.json dictionary data asset file, parses its structured
// JSON content configuration payload, and returns a collection of initialized Category instances.
// It panics if a physical file reading sequence failure happens or a parsing error occurs.
func Load() []dict.Category {
	bytes, err := fs.ReadFile(dictFS, dictName)
	if err != nil {
		panic(&ErrReadFile{FileName: dictName, Err: err})
	}

	var categories []dict.Category
	if err := json.Unmarshal(bytes, &categories); err != nil {
		panic(&ErrUnmarshalJSON{FileName: dictName, Err: err})
	}

	return categories
}
