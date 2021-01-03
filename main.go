package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/brombaut/ascii-chess/internal/game"
	"github.com/brombaut/ascii-chess/internal/pieces"
	"github.com/brombaut/ascii-chess/internal/printing"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome to Terminal Chess")
	fmt.Println()
	game := game.NewGame()
	printing.Print(game)
	fmt.Print("Piece to move: ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "quit" {
			fmt.Println("Goodbye")
			break
		}
		fmt.Print("\033[H\033[2J")
		chars := []rune(input)
		colLetter := chars[0]
		rowNumber := chars[1]
		posInput, err := pieces.NewBoardPosition(rowNumber, colLetter)
		if err != nil {
			fmt.Println(err.Error())
		}

		game.HandlePositionInput(*posInput)

		fmt.Println()
		printing.Print(game)
		fmt.Print("Piece to move: ")
	}
}
