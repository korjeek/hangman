package words

import "fmt"

// ErrUnmarshal represents an error that occurs when a byte slice cannot be successfully
// parsed into a valid Words structure due to an incorrect number of elements.
type ErrUnmarshal struct {
	Words []string
}

func (e ErrUnmarshal) Error() string {
	return fmt.Sprintf("Invalid words format: %v. Should be: [<hidden>, <guessed>]", e.Words)
}
