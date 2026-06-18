package main

import (
	"errors"
	"fmt"
	"hangman/internal/game"
	"hangman/internal/mode"
	"hangman/internal/repo"
	"os"

	"github.com/alexflint/go-arg"
)

type Args struct {
	Words      game.Words      `arg:"positional" placeholder:"<word>" help:"Words pair for testing mode"`
	Category   game.Category   `arg:"-C,--category" help:"Category for the word"`
	Difficulty game.Difficulty `arg:"-D,--difficulty" help:"Difficulty: EASY, MEDIUM or HARD"`
}

func (a Args) Validate() error {
	if !a.Words.IsEmpty() &&
		(a.Category != game.RandomCategory || a.Difficulty != game.RandomDifficulty) {
		return errors.New("cannot use flags (--category, --difficulty) in non interactive mode")
	}

	return nil
}

func main() {
	var args Args
	arg.MustParse(&args)

	if err := args.Validate(); err != nil {
		fmt.Printf("Usage error: %v\n", err)
		os.Exit(1)
	}

	if err := run(args); err != nil {
		fmt.Printf("Runtime error: %v\n", err)
		os.Exit(1)
	}
}

func run(args Args) error {
	var gm mode.Mode

	if !args.Words.IsEmpty() {
		session, err := game.NewSession(args.Words.Hidden(), len(args.Words.Guessed()))
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}

		gm = mode.NewNonInteractive(session, args.Words.Guessed())
	} else {
		r, err := repo.New()
		if err != nil {
			return fmt.Errorf("failed to init repo: %w", err)
		}

		var ctg repo.Category
		if args.Category != game.RandomCategory {
			ctg, err = r.Category(args.Category)
			if err != nil {
				return fmt.Errorf("failed to get category: %w", err)
			}
		} else {
			ctg = r.RandomCategory()
		}

		hidden := r.RandomWord(ctg).Word
		session, err := game.NewSession(hidden, args.Difficulty.MaxAttempts())
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}

		gm = mode.NewInteractive(session)
	}

	return gm.Run()
}
