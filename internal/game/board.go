package game

import (
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
