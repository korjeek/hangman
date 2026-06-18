package config

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
)

type Category string

const (
	Animals   Category = "animals"
	Fruits    Category = "fruits"
	Countries Category = "countries"
)

type GameConfig struct {
	difficulty Difficulty
	category   Category
	maxTries   int
}

type Option func(*GameConfig)

func WithDifficulty(d Difficulty) Option {
	return func(c *GameConfig) {
		c.difficulty = d
		switch d {
		case Easy:
			c.maxTries = 8
		case Medium:
			c.maxTries = 6
		case Hard:
			c.maxTries = 4
		}
	}
}

func WithCategory(c Category) Option {
	return func(cfg *GameConfig) {
		cfg.category = c
	}
}
