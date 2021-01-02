package main

import (
	"github.com/brombaut/ascii-chess/internal/game"
	"github.com/brombaut/ascii-chess/internal/printing"
)

func main() {
	printing.Print(game.NewGame())
}
