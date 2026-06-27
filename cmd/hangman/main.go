package main

import (
	"errors"
	"fmt"
	"hangman/internal/app/mode"
	"hangman/internal/domain/ctg"
	"hangman/internal/domain/dict"
	"hangman/internal/domain/diff"
	"hangman/internal/domain/session"
	"hangman/internal/domain/words"
	"hangman/internal/infra/repo"
	"hangman/internal/ui"
	"math/rand/v2"
	"os"

	"github.com/alexflint/go-arg"
)

// Args holds the command-line arguments configuration for the application.
type Args struct {
	Words      words.Words     `arg:"positional" placeholder:"<word>" help:"word pair for non interactive mode"`
	Category   ctg.Category    `arg:"-C,--category" help:"category for the word"`
	Difficulty diff.Difficulty `arg:"-D,--difficulty" help:"difficulty: easy, medium or hard"`
}

func (a Args) validate() error {
	if !a.Words.IsEmpty() &&
		(a.Category != ctg.RandomCategory || a.Difficulty != diff.RandomDifficulty) {
		return errors.New("cannot use flags (--category, --difficulty) in non interactive mode")
	}
	return nil
}

func main() {
	var args Args
	arg.MustParse(&args)

	if err := args.validate(); err != nil {
		fmt.Printf("usage error: %v\n", err)
		os.Exit(1)
	}

	if err := run(args); err != nil {
		fmt.Printf("runtime error: %v\n", err)
		os.Exit(1)
	}
}

func run(args Args) error {
	states := ui.Load()
	consoleUI := ui.New(os.Stdin, os.Stdout, states)

	var gm mode.Mode

	// non interactive mode
	if !args.Words.IsEmpty() {
		sess, err := session.NewSession(
			args.Words.GetHidden(),
			len(args.Words.GetGuessed()),
		)
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}

		gm = mode.NewNonInteractive(sess, args.Words.GetGuessed(), consoleUI)
		return gm.Run()
	}

	// interactive mode
	categories := repo.Load()
	dictionaryRepository := repo.NewJsonDictRepository(categories)

	cryptoRand := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))
	wordSelector := dict.NewWordSelector(dictionaryRepository, cryptoRand)
	selectedCategory := wordSelector.RandomCategory()

	if args.Category != ctg.RandomCategory {
		if category, ok := dictionaryRepository.Category(args.Category); ok {
			selectedCategory = category
		}
	}

	selectedWord := wordSelector.RandomWordFromCategory(selectedCategory)
	sess, err := session.NewSession(
		selectedWord.Word,
		args.Difficulty.GetMaxAttempts(),
		session.WithHint(selectedWord.Hint),
	)

	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	gm = mode.NewInteractive(sess, consoleUI)
	return gm.Run()
}
