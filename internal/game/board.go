package game

import (
	"github.com/brombaut/ascii-chess/internal/pieces"
)

type Board struct {
	Pieces [][]pieces.Piece
}

func NewBoard() *Board {
	var b Board
	var row8 = []pieces.Piece{
		pieces.NewRook(BLACK),
		pieces.NewKnight(BLACK),
		pieces.NewBishop(BLACK),
		pieces.NewQueen(BLACK),
		pieces.NewKing(BLACK),
		pieces.NewBishop(BLACK),
		pieces.NewKnight(BLACK),
		pieces.NewRook(BLACK),
	}
	var row7 = []pieces.Piece{
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
		pieces.NewPawn(BLACK),
	}
	var row6 = []pieces.Piece{
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
	}
	var row5 = []pieces.Piece{
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
	}
	var row4 = []pieces.Piece{
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
	}
	var row3 = []pieces.Piece{
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
		pieces.NewNoPiece(NONE),
	}
	var row2 = []pieces.Piece{
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
		pieces.NewPawn(WHITE),
	}
	var row1 = []pieces.Piece{
		pieces.NewRook(WHITE),
		pieces.NewKnight(WHITE),
		pieces.NewBishop(WHITE),
		pieces.NewQueen(WHITE),
		pieces.NewKing(WHITE),
		pieces.NewBishop(WHITE),
		pieces.NewKnight(WHITE),
		pieces.NewRook(WHITE),
	}
	b.Pieces = [][]pieces.Piece{
		row8,
		row7,
		row6,
		row5,
		row4,
		row3,
		row2,
		row1,
	}
	return &b
}
