package game

import (
	"fmt"

	"github.com/brombaut/ascii-chess/internal/pieces"
)

type GameState string

const (
	SELECT_PIECE = "SELECT_PIECE"
	SELECT_MOVE  = "SELECT_MOVE"
)

type Game struct {
	Board         *Board
	Moves         []pieces.Move
	White         Player
	Black         Player
	Turn          PlayerColor
	State         GameState
	SelectedPiece *pieces.PlayingPiece
}

func (g Game) HandlePositionInput(input *pieces.Position) error {
	// For now, assume that input is a valid position like 'e3'
	var err error
	if g.State == SELECT_PIECE {
		err = g.selectPiece(input)
	} else {
		err = g.selectMove(input)
	}
	return err
}

func (g Game) selectPiece(pos *pieces.Position) error {
	pap, err := g.Board.PieceAtPosition(pos)
	if err != nil {
		return err
	}
	if (*pap).Literal() == pieces.NO_PIECE {
		return fmt.Errorf("Can't select no piece")
	}
	err = g.setSelectedPiece(pap)
	if err != nil {
		return err
	}

	g.State = SELECT_MOVE
	return nil
}

func (g Game) setSelectedPiece(piece *pieces.PlayingPiece) error {
	if g.SelectedPiece != nil {
		(*g.SelectedPiece).SetIsSelected(false)
	}
	g.SelectedPiece = piece
	(*g.SelectedPiece).SetIsSelected(true)
	return nil
}

func (g Game) selectMove(pos *pieces.Position) error {
	g.State = SELECT_PIECE
	g.nextTurn()
	return nil
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
