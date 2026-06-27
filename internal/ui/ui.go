package ui

import (
	"bufio"
	"errors"
	"fmt"
	"hangman/internal/domain/session"
	"io"
	"os"
	"strings"
)

// UserInterface defines the abstract behavioral contract required for rendering
// game outputs and gathering input choices from the user.
//
//go:generate mockgen -source=ui.go -destination=../app/mode/mocks/mock_ui.go -package=mocks
type UserInterface interface {
	Clear()
	Println(a ...any)
	AskRune() rune
	PrintHangman(leftAttempts int)
	PrintGuessResult(result session.GuessResult, leftAtt int)
	PrintHint(hint string)
	PrintGameStatus(status session.Status)
}

// ConsoleUI implements the UserInterface behavioral contract using traditional
// standard terminal input/output communication streams.
type ConsoleUI struct {
	reader   *bufio.Reader
	output   io.Writer
	hmStates []string
}

func New(input io.Reader, output io.Writer, states []string) *ConsoleUI {
	return &ConsoleUI{
		reader:   bufio.NewReader(input),
		output:   output,
		hmStates: states,
	}
}

// Println writes formatted data to the underlying stream followed by a newline character.
func (c *ConsoleUI) Println(a ...any) {
	_, err := fmt.Fprintln(c.output, a...)
	if err != nil {
		panic(err)
	}
}

func (c *ConsoleUI) printf(format string, a ...any) {
	_, err := fmt.Fprintf(c.output, format, a...)
	if err != nil {
		panic(err)
	}
}

// Clear flushes the screen by executing standard ANSI terminal screen clearing escape sequence codes.
func (c *ConsoleUI) Clear() {
	c.Println("\033[H\033[2J")
}

// Ask displays a target message prompt string and reads characters from the input stream until a newline separator is hit.
func (c *ConsoleUI) Ask(prompt string) string {
	for {
		c.printf("%s > \n", prompt)
		input, err := c.reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				c.Println("\ninput stream closed. exiting game.")
				os.Exit(0)
			}
			continue
		}

		input = strings.TrimSpace(input)
		return input
	}
}

// AskRune continuously processes text input prompts until the user enters exactly one distinct standalone character string.
func (c *ConsoleUI) AskRune() rune {
	for {
		input := c.Ask("enter the suggested letter")
		runes := []rune(input)

		if len(runes) != 1 {
			c.Println("you should enter only one letter!")
			continue
		}

		return runes[0]
	}
}

// PrintGuessResult outputs localized contextual reaction text fragments matching explicit domain turn outcome values.
func (c *ConsoleUI) PrintGuessResult(result session.GuessResult, leftAtt int) {
	switch result {
	case session.AlreadyGuessed:
		c.Println("this letter is already guessed")
	case session.Incorrect:
		c.Println("this letter is not in the hidden word")
	case session.Correct:
		c.Println("the letter is guessed")
	}

	c.printf("attempts left > %d\n", leftAtt)
}

// PrintGameStatus displays contextual congratulations or commiseration texts matching explicit final lifecycle endpoints.
func (c *ConsoleUI) PrintGameStatus(status session.Status) {
	switch status {
	case session.Win:
		c.printf("you win, congratulations!")
	case session.Lose:
		c.printf("you lose, better luck next time!")
	case session.InProgress:
		c.printf("game still in progress...")
	}
}

// PrintHint draws explicit informative clue metadata strings onto the standard execution terminal window context.
func (c *ConsoleUI) PrintHint(hint string) {
	c.printf("hint > %s\n", hint)
}

// PrintHangman evaluates bounding safety constraints and draws the appropriate textual ascii frame illustration matching active attempts.
func (c *ConsoleUI) PrintHangman(leftAttempts int) {
	if leftAttempts < 0 {
		leftAttempts = 0
	}
	if leftAttempts >= len(c.hmStates) {
		leftAttempts = len(c.hmStates) - 1
	}

	c.Println(c.hmStates[leftAttempts])
}
