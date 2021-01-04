package game

import (
	"fmt"

	"github.com/brombaut/ascii-chess/internal/pieces"
)

type Board struct {
	Pieces [8][8]pieces.PlayingPiece
}

func NewBoard() *Board {
	var b Board
	b.Pieces = pieces.StartingPlayingPieces()
	return &b
}

func (b Board) PieceAtPosition(p *pieces.Position) (*pieces.PlayingPiece, error) {
	if p.Row() < 0 || p.Row() > 7 || p.Col() < 0 || p.Col() > 7 {
		return nil, fmt.Errorf("Position index out of bounds")
	}
	return &b.Pieces[p.Row()][p.Col()], nil
}
