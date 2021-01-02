package game

import (
	"github.com/brombaut/ascii-chess/internal/pieces"
)

type Game struct {
	Board *Board
	Moves []pieces.Move
	White Player
	Black Player
	Turn  PlayerColor
}

func NewGame() *Game {
	var g Game
	g.Board = NewBoard()
	g.Moves = []pieces.Move{}
	g.White = Player{Color: WHITE}
	g.Black = Player{Color: BLACK}
	g.Turn = WHITE
	return &g
}
