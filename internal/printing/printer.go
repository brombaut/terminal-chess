package printing

import (
	"fmt"

	"github.com/brombaut/ascii-chess/internal/game"
	"github.com/brombaut/ascii-chess/internal/pieces"
	"github.com/gookit/color"
)

const (
	WHITE_KING   = "\u2654"
	WHITE_QUEEN  = "\u2655"
	WHITE_ROOK   = "\u2656"
	WHITE_BISHOP = "\u2657"
	WHITE_KNIGHT = "\u2658"
	WHITE_PAWN   = "\u2659"
	BLACK_KING   = "\u265A"
	BLACK_QUEEN  = "\u265B"
	BLACK_ROOK   = "\u265C"
	BLACK_BISHOP = "\u265D"
	BLACK_KNIGHT = "\u265E"
	BLACK_PAWN   = "\u265F"
	SPACE        = " "
)

func Print(g *game.Game) {
	printBoard(g.Board)
}

func printBoard(b *game.Board) {
	lightSquare := color.New(color.FgBlack, color.BgWhite)
	darkSquare := color.New(color.FgBlack, color.BgCyan)
	selectedSquare := color.New(color.FgBlack, color.BgLightRed)
	rowStr := "(%d) "
	squareStr := " %s  "
	for i, row := range b.Pieces {
		fmt.Printf(rowStr, 8-i)
		for j, piece := range row {
			if piece.IsSelected() {
				selectedSquare.Printf(squareStr, getBoardCode(piece))
			} else if (i+j)%2 == 0 {
				lightSquare.Printf(squareStr, getBoardCode(piece))
			} else {
				darkSquare.Printf(squareStr, getBoardCode(piece))
			}
		}
		fmt.Println()
	}
	cols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	colStr := "  %s "
	fmt.Print("    ")
	for _, col := range cols {
		fmt.Printf(colStr, col)
	}
	fmt.Println()
}

func getBoardCode(piece pieces.PlayingPiece) string {
	switch piece.Literal() {
	case pieces.KING:
		if piece.Color() == pieces.WHITE {
			return WHITE_KING
		}
		return BLACK_KING
	case pieces.QUEEN:
		if piece.Color() == pieces.WHITE {
			return WHITE_QUEEN
		}
		return BLACK_QUEEN
	case pieces.ROOK:
		if piece.Color() == pieces.WHITE {
			return WHITE_ROOK
		}
		return BLACK_ROOK
	case pieces.BISHOP:
		if piece.Color() == pieces.WHITE {
			return WHITE_BISHOP
		}
		return BLACK_BISHOP
	case pieces.KNIGHT:
		if piece.Color() == pieces.WHITE {
			return WHITE_KNIGHT
		}
		return BLACK_KNIGHT
	case pieces.PAWN:
		if piece.Color() == pieces.WHITE {
			return WHITE_PAWN
		}
		return BLACK_PAWN
	default:
		return SPACE
	}
}
