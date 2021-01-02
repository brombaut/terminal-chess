package pieces

type PieceLiteral string
type PieceColor string

const (
	WHITE = "WHITE"
	BLACK = "BLACK"
	NONE  = "NONE"
)

const (
	BISHOP   = "BISHOP"
	KING     = "KING"
	KNIGHT   = "KNIGHT"
	NO_PIECE = "NO_PIECE"
	PAWN     = "PAWN"
	QUEEN    = "QUEEN"
	ROOK     = "ROOK"
)

type Piece interface {
	PossibleMoves() []Move
	Color() PieceColor
	Literal() PieceLiteral
}

type King struct {
	color   PieceColor
	literal PieceLiteral
}

func NewKing(color PieceColor) King {
	var k King
	k.color = color
	k.literal = KING
	return k
}
func (k King) PossibleMoves() []Move {
	return []Move{}
}
func (k King) Literal() PieceLiteral {
	return k.literal
}
func (k King) Color() PieceColor {
	return k.color
}

type Queen struct {
	color   PieceColor
	literal PieceLiteral
}

func NewQueen(color PieceColor) Queen {
	var q Queen
	q.color = color
	q.literal = QUEEN
	return q
}
func (q Queen) PossibleMoves() []Move {
	return []Move{}
}
func (q Queen) Literal() PieceLiteral {
	return q.literal
}
func (q Queen) Color() PieceColor {
	return q.color
}

type Rook struct {
	color   PieceColor
	literal PieceLiteral
}

func NewRook(color PieceColor) Rook {
	var r Rook
	r.color = color
	r.literal = ROOK
	return r
}
func (r Rook) PossibleMoves() []Move {
	return []Move{}
}
func (r Rook) Literal() PieceLiteral {
	return r.literal
}
func (r Rook) Color() PieceColor {
	return r.color
}

type Bishop struct {
	color   PieceColor
	literal PieceLiteral
}

func NewBishop(color PieceColor) Bishop {
	var b Bishop
	b.color = color
	b.literal = BISHOP
	return b
}
func (b Bishop) PossibleMoves() []Move {
	return []Move{}
}
func (b Bishop) Literal() PieceLiteral {
	return b.literal
}
func (b Bishop) Color() PieceColor {
	return b.color
}

type Knight struct {
	color   PieceColor
	literal PieceLiteral
}

func NewKnight(color PieceColor) Knight {
	var k Knight
	k.color = color
	k.literal = KNIGHT
	return k
}
func (k Knight) PossibleMoves() []Move {
	return []Move{}
}
func (k Knight) Literal() PieceLiteral {
	return k.literal
}
func (k Knight) Color() PieceColor {
	return k.color
}

type Pawn struct {
	color   PieceColor
	literal PieceLiteral
}

func NewPawn(color PieceColor) Pawn {
	var p Pawn
	p.color = color
	p.literal = PAWN
	return p
}
func (p Pawn) PossibleMoves() []Move {
	return []Move{}
}
func (p Pawn) Literal() PieceLiteral {
	return p.literal
}
func (p Pawn) Color() PieceColor {
	return p.color
}

type NoPiece struct {
	color   PieceColor
	literal PieceLiteral
}

func NewNoPiece(color PieceColor) NoPiece {
	var np NoPiece
	np.color = color
	np.literal = NO_PIECE
	return np
}
func (np NoPiece) PossibleMoves() []Move {
	return []Move{}
}
func (np NoPiece) Literal() PieceLiteral {
	return np.literal
}
func (np NoPiece) Color() PieceColor {
	return np.color
}
