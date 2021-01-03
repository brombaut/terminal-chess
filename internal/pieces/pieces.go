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

type PlayingPiece interface {
	PossibleMoves() []Move
	Color() PieceColor
	Literal() PieceLiteral
	Position() Position
}

type Piece struct {
	color    PieceColor
	literal  PieceLiteral
	position Position
}

func (p Piece) Literal() PieceLiteral {
	return p.literal
}
func (p Piece) Color() PieceColor {
	return p.color
}
func (p Piece) Position() Position {
	return p.position
}

type King struct {
	Piece
}

func NewKing(color PieceColor, startingPosition Position) King {
	var k King
	k.color = color
	k.literal = KING
	k.position = startingPosition
	return k
}
func (k King) PossibleMoves() []Move {
	return []Move{}
}

type Queen struct {
	Piece
}

func NewQueen(color PieceColor, startingPosition Position) Queen {
	var q Queen
	q.color = color
	q.literal = QUEEN
	q.position = startingPosition
	return q
}
func (q Queen) PossibleMoves() []Move {
	return []Move{}
}

type Rook struct {
	Piece
}

func NewRook(color PieceColor, startingPosition Position) Rook {
	var r Rook
	r.color = color
	r.literal = ROOK
	r.position = startingPosition
	return r
}
func (r Rook) PossibleMoves() []Move {
	return []Move{}
}

type Bishop struct {
	Piece
}

func NewBishop(color PieceColor, startingPosition Position) Bishop {
	var b Bishop
	b.color = color
	b.literal = BISHOP
	b.position = startingPosition
	return b
}
func (b Bishop) PossibleMoves() []Move {
	return []Move{}
}

type Knight struct {
	Piece
}

func NewKnight(color PieceColor, startingPosition Position) Knight {
	var k Knight
	k.color = color
	k.literal = KNIGHT
	k.position = startingPosition
	return k
}
func (k Knight) PossibleMoves() []Move {
	return []Move{}
}

type Pawn struct {
	Piece
}

func NewPawn(color PieceColor, startingPosition Position) Pawn {
	var p Pawn
	p.color = color
	p.literal = PAWN
	p.position = startingPosition
	return p
}
func (p Pawn) PossibleMoves() []Move {
	return []Move{}
}

type NoPiece struct {
	Piece
}

func NewNoPiece(color PieceColor, startingPosition Position) NoPiece {
	var np NoPiece
	np.color = color
	np.literal = NO_PIECE
	np.position = startingPosition
	return np
}
func (np NoPiece) PossibleMoves() []Move {
	return []Move{}
}
