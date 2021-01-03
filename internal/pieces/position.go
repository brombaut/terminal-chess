package pieces

import "fmt"

type Position struct {
	row      int
	col      int
	boardRow rune
	boardCol rune
}

func (p Position) Row() int {
	return p.row
}
func (p Position) Col() int {
	return p.col
}

func NewIdxPosition(row int, col int) (*Position, error) {
	var p Position
	p.row = row
	p.col = col
	var br, bc rune
	var ok bool
	br, ok = rowIdxToBoardRow[p.row]
	if !ok {
		return nil, fmt.Errorf("Invalid row id")
	}
	bc, ok = colIdxToBoardCol[p.col]
	if !ok {
		return nil, fmt.Errorf("Invalid col id")
	}
	p.boardRow = br
	p.boardCol = bc
	return &p, nil
}

func NewBoardPosition(row rune, col rune) (*Position, error) {
	var p Position
	p.boardRow = row
	p.boardCol = col
	var rid, cid int
	var ok bool
	rid, ok = boardRowToRowIdx[row]
	if !ok {
		return nil, fmt.Errorf("Invalid board row")
	}
	cid, ok = boardColToColIdx[col]
	if !ok {
		return nil, fmt.Errorf("Invalid board col")
	}
	p.row = rid
	p.col = cid
	return &p, nil
}

var rowIdxToBoardRow = map[int]rune{
	0: '8',
	1: '7',
	2: '6',
	3: '5',
	4: '4',
	5: '3',
	6: '2',
	7: '1',
}

var colIdxToBoardCol = map[int]rune{
	0: 'a',
	1: 'b',
	2: 'c',
	3: 'd',
	4: 'e',
	5: 'f',
	6: 'g',
	7: 'h',
}

var boardRowToRowIdx = map[rune]int{
	'8': 0,
	'7': 1,
	'6': 2,
	'5': 3,
	'4': 4,
	'3': 5,
	'2': 6,
	'1': 7,
}

var boardColToColIdx = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
}
