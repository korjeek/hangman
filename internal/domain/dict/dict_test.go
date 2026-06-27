package dict_test

import (
	"testing"

	"hangman/internal/domain/ctg"
	"hangman/internal/domain/dict"
	"hangman/internal/domain/dict/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWordSelector_RandomWordFromCategory_WithMockedRng(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockRng := mocks.NewMockIntPrng(ctrl)

	inputCategory := dict.Category{
		Name: ctg.Animals,
		Words: []dict.Word{
			{Word: "cat", Hint: "meow"},
			{Word: "dog", Hint: "bark"},
		},
	}

	mockRng.EXPECT().
		IntN(2).
		Return(1).
		Times(1)

	selector := dict.NewWordSelector(nil, mockRng)
	actualWord := selector.RandomWordFromCategory(inputCategory)

	assert.Equal(t, "dog", actualWord.Word)
}

func TestWordSelector_RandomCategory_WithBothMocks(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockRepository(ctrl)
	mockRng := mocks.NewMockIntPrng(ctrl)

	mockCats := []dict.Category{
		{Name: ctg.Animals},
		{Name: ctg.Nature},
	}

	mockRepo.EXPECT().
		ListCategories().
		Return(mockCats).
		Times(1)

	mockRng.EXPECT().
		IntN(2).
		Return(0).
		Times(1)

	selector := dict.NewWordSelector(mockRepo, mockRng)

	actualCategory := selector.RandomCategory()

	assert.Equal(t, ctg.Animals, actualCategory.Name)
}
