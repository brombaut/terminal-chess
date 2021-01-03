package game

import (
	"github.com/brombaut/ascii-chess/internal/pieces"
)

type GameState string

const (
	SELECT_PIECE = "SELECT_PIECE"
	SELECT_MOVE  = "SELECT_MOVE"
)

type Game struct {
	Board *Board
	Moves []pieces.Move
	White Player
	Black Player
	Turn  PlayerColor
	State GameState
}

func (g Game) HandlePositionInput(input pieces.Position) {
	// For now, assume that input is a valid position like 'e3'

}

func (g Game) selectPiece(pos pieces.Position) {

	g.State = SELECT_MOVE
}

func (g Game) selectMove() {
	g.State = SELECT_PIECE
}

func (g Game) nextTurn() {
	if g.Turn == WHITE {
		g.Turn = BLACK
	} else {
		g.Turn = WHITE
	}
}

func NewGame() *Game {
	var g Game
	g.Board = NewBoard()
	g.Moves = []pieces.Move{}
	g.White = Player{Color: WHITE}
	g.Black = Player{Color: BLACK}
	g.Turn = WHITE
	g.State = SELECT_PIECE
	return &g
}
