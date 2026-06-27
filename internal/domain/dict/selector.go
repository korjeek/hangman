package dict

// IntPrng defines the behavior required for generating pseudo-random integers.
//
//go:generate mockgen -source=selector.go -destination=mocks/mock_selector.go -package=mocks
type IntPrng interface {
	IntN(n int) int
}

// WordSelector coordinates the business logic for randomly choosing game categories and target words.
type WordSelector struct {
	repo Repository
	rng  IntPrng
}

// NewWordSelector creates and returns a pointer to a newly initialized WordSelector.
func NewWordSelector(repo Repository, rng IntPrng) *WordSelector {
	return &WordSelector{
		repo: repo,
		rng:  rng,
	}
}

// RandomWordFromCategory picks a random Word structure from the provided Category.
// It returns an empty Word structure if the category contains no words.
func (s *WordSelector) RandomWordFromCategory(cat Category) Word {
	if len(cat.Words) == 0 {
		return Word{}
	}
	return cat.Words[s.rng.IntN(len(cat.Words))]
}

// RandomCategory retrieves all available categories from the storage repository and selects one at random.
// It returns an empty Category structure if no categories are found.
func (s *WordSelector) RandomCategory() Category {
	categories := s.repo.ListCategories()
	if len(categories) == 0 {
		return Category{}
	}
	return categories[s.rng.IntN(len(categories))]
}
