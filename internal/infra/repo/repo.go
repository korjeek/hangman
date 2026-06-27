package repo

import (
	"hangman/internal/domain/ctg"
	"hangman/internal/domain/dict"
)

// JsonDictRepository implements the dict.Repository interface, providing
// an in-memory data store backed by a map for fast category lookups.
type JsonDictRepository struct {
	categories map[ctg.Category]dict.Category
}

func NewJsonDictRepository(categories []dict.Category) *JsonDictRepository {
	repoMap := make(map[ctg.Category]dict.Category)
	for _, c := range categories {
		repoMap[c.Name] = c
	}

	return &JsonDictRepository{categories: repoMap}
}

// ListCategories returns a flat collection slice containing all available
// word Category structures currently managed by the repository.
func (r *JsonDictRepository) ListCategories() []dict.Category {
	list := make([]dict.Category, 0, len(r.categories))
	for _, c := range r.categories {
		list = append(list, c)
	}
	return list
}

// Category searches for a single dictionary category by its explicit identifier.
// It returns the Category data and true if found; otherwise, an empty Category and false.
func (r *JsonDictRepository) Category(category ctg.Category) (dict.Category, bool) {
	ctg, ok := r.categories[category]
	if !ok {
		return dict.Category{}, false
	}
	return ctg, true
}
