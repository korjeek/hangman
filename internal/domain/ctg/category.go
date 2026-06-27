package ctg

import (
	"fmt"
	"strings"
)

// Category represents a distinct classification group for words in the game dictionary.
type Category int

const (
	RandomCategory Category = iota
	Animals
	Countries
	Nature
	Food
)

func (c *Category) UnmarshalText(text []byte) error {
	input := strings.ToLower(string(text))

	switch input {
	case "animals":
		*c = Animals
	case "countries":
		*c = Countries
	case "nature":
		*c = Nature
	case "food":
		*c = Food
	default:
		return fmt.Errorf("unknown category %q: must be animals, countries, nature, or food", input)
	}
	return nil
}
