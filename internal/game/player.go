package game

type PlayerColor string

type Player struct {
	Color PlayerColor
}

const (
	WHITE = "WHITE"
	BLACK = "BLACK"
	NONE  = "NONE"
)
