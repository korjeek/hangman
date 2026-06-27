package ui

import (
	"embed"
	"path/filepath"
	"sort"
)

//go:embed images/*.txt
var hangmanFS embed.FS

const imagesDir = "images"

// Load reads all available hangman state text asset files from the embedded filesystem,
// sorts them alphabetically by filename, and returns them as a slice of raw strings.
// It panics if the asset directory cannot be accessed or a file reading error occurs.
func Load() []string {
	files, err := hangmanFS.ReadDir(imagesDir)
	if err != nil {
		panic(ErrReadDir{Err: err})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	hangmanStates := make([]string, 0, len(files))

	for _, file := range files {
		fullPath := filepath.Join(imagesDir, file.Name())

		data, err := hangmanFS.ReadFile(fullPath)
		if err != nil {
			panic(ErrReadFile{Err: err, fullPath: fullPath})
		}

		hangmanStates = append(hangmanStates, string(data))
	}

	return hangmanStates
}
