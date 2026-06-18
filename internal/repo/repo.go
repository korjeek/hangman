package repo

import (
	"encoding/json"
	"errors"
	"hangman/internal/game"
	"math/rand/v2"
)

type Repository interface {
	RandomWord(category Category) Word
	RandomCategory() Category
	ListCategories() []Category
	Category(ctg game.Category) (Category, error)
}

type repository struct {
	data   *Dictionary
	ctgIdx map[game.Category]Category
}

func New() (Repository, error) {
	var d Dictionary
	if err := json.Unmarshal(words, &d); err != nil {
		return nil, err
	}

	ctgIdx := make(map[game.Category]Category, len(d.Categories))
	for _, ctg := range d.Categories {
		ctgIdx[ctg.Name] = ctg
	}

	return &repository{
		data:   &d,
		ctgIdx: ctgIdx,
	}, nil
}

func (r *repository) RandomWord(category Category) Word {
	w := category.Words
	return w[rand.IntN(len(w))]
}

func (r *repository) RandomCategory() Category {
	cat := r.data.Categories
	return cat[rand.IntN(len(cat))]
}

func (r *repository) ListCategories() []Category {
	return r.data.Categories
}

func (r *repository) Category(ctg game.Category) (Category, error) {
	cat, ok := r.ctgIdx[ctg]

	if !ok {
		return Category{}, errors.New("category not found")
	}

	return cat, nil
}
