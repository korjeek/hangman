package ui

import "fmt"

// ErrReadDir represents a critical infrastructure error that occurs when the system
// fails to access or read the directory containing the embedded hangman state images.
type ErrReadDir struct {
	Err error
}

func (e ErrReadDir) Error() string {
	return fmt.Sprintf("critical ui error: can't read images directory: %v", e.Err)
}

// ErrReadFile represents a critical infrastructure error that occurs when the system
// fails to open or read a specific explicit hangman state text asset file.
type ErrReadFile struct {
	fullPath string
	Err      error
}

func (e ErrReadFile) Error() string {
	return fmt.Sprintf("critical ui error: can't read image file %s: %v", e.fullPath, e.Err)
}
