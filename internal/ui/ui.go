package ui

import (
	"bufio"
	"fmt"
	"hangman/internal/game"
	"os"
	"strings"
)

type ConsoleUI struct {
	reader *bufio.Reader
}

func New() *ConsoleUI {
	return &ConsoleUI{
		reader: bufio.NewReader(os.Stdin),
	}
}

// Clear очищает экран консоли
func (c *ConsoleUI) Clear() {
	fmt.Print("\033[H\033[2J")
}

func (c *ConsoleUI) Ask(prompt string) string {
	for {
		fmt.Printf("%s > ", prompt)
		input, err := c.reader.ReadString('\n')
		if err != nil {
			continue
		}

		input = strings.TrimSpace(input)
		return input
	}
}

func (c *ConsoleUI) AskRune(prompt string) rune {
	for {
		input := c.Ask(prompt)
		runes := []rune(input)

		if len(runes) == 1 {
			return runes[0]
		}

		fmt.Println("Ошибка: введите ровно один символ.")
	}
}

func (c *ConsoleUI) PrintGuessResult(result game.GuessResult) string {
	switch result {
	case game.AlreadyGuessed:
		fmt.Println("This letter is already guessed")
	case game.Incorrect:
		fmt.Println("This letter is not in the hidden word")
	case game.Correct:
		fmt.Println("The letter is guessed")
	}
}
