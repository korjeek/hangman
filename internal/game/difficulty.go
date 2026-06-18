package game

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Difficulty int

const (
	RandomDifficulty Difficulty = iota
	Easy
	Normal
	Hard
)

func (d *Difficulty) MaxAttempts() int {
	val := *d
	if val == RandomDifficulty {
		val = Difficulty(rand.IntN(3) + 1)
	}

	switch val {
	case Easy:
		return 10
	case Normal:
		return 7
	case Hard:
		return 4
	default:
		return 7
	}
}

func (d *Difficulty) UnmarshalText(text []byte) error {
	input := strings.ToLower(string(text))

	switch input {
	case "easy":
		*d = Easy
	case "normal":
		*d = Normal
	case "hard":
		*d = Hard
	default:
		return fmt.Errorf("unknown difficulty %q: must be easy, normal, or hard", input)
	}
	return nil
}
