package repo_test

import (
	"hangman/internal/infra/repo"
	"testing"

	"hangman/internal/domain/ctg"
	"hangman/internal/domain/dict"

	"github.com/stretchr/testify/assert"
)

func TestJsonDictRepository_Category(t *testing.T) {
	t.Parallel()

	animalsCategory := dict.Category{
		Name:  ctg.Animals,
		Words: []dict.Word{{Word: "cat", Hint: "dog"}},
	}

	categories := []dict.Category{animalsCategory}
	repository := repo.NewJsonDictRepository(categories)

	type TestCase struct {
		name           string
		searchCategory ctg.Category
		expectedFound  bool
		expectedData   dict.Category
	}

	tests := []TestCase{
		{
			name:           "category exists",
			searchCategory: ctg.Animals,
			expectedFound:  true,
			expectedData:   animalsCategory,
		},
		{
			name:           "category does not exist",
			searchCategory: ctg.Nature,
			expectedFound:  false,
			expectedData:   dict.Category{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()

			actualData, actualFound := repository.Category(test.searchCategory)

			assert.Equal(tt, test.expectedFound, actualFound)
			assert.Equal(tt, test.expectedData, actualData)
		})
	}
}

func TestJsonDictRepository_ListCategories(t *testing.T) {
	t.Parallel()

	cat1 := dict.Category{Name: ctg.Animals}
	cat2 := dict.Category{Name: ctg.Nature}

	repository := repo.NewJsonDictRepository([]dict.Category{cat1, cat2})

	actualList := repository.ListCategories()

	assert.Len(t, actualList, 2)
	assert.ElementsMatch(t, []dict.Category{cat1, cat2}, actualList)
}
