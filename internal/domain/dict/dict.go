package dict

import "hangman/internal/domain/ctg"

// Word represents a dictionary entry containing a target secret word and its associated clue.
type Word struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

// Category bundles a group of thematic words together under a specific domain category identifier.
type Category struct {
	Name  ctg.Category `json:"name"`
	Words []Word       `json:"words"`
}

// Repository defines the data store abstraction layer for retrieving game dictionaries and specific categories.
//
//go:generate mockgen -source=dict.go -destination=mocks/mock_dict.go -package=mocks
type Repository interface {
	ListCategories() []Category
	Category(ctg ctg.Category) (Category, bool)
}
