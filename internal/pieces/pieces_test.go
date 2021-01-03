package pieces

import (
	"testing"
)

func TestKing(t *testing.T) {
	var p *King
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewKing(BLACK, pos)
	if p.Color() != BLACK {
		t.Errorf("Wrong color, want=BLACK, got=%s", p.Color())
	}
	if p.Literal() != KING {
		t.Errorf("Wrong literal, want=KING, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}

func TestQueen(t *testing.T) {
	var p *Queen
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewQueen(WHITE, pos)
	if p.Color() != WHITE {
		t.Errorf("Wrong color, want=WHITE, got=%s", p.Color())
	}
	if p.Literal() != QUEEN {
		t.Errorf("Wrong literal, want=QUEEN, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}

func TestBishop(t *testing.T) {
	var p *Bishop
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewBishop(WHITE, pos)
	if p.Color() != WHITE {
		t.Errorf("Wrong color, want=WHITE, got=%s", p.Color())
	}
	if p.Literal() != BISHOP {
		t.Errorf("Wrong literal, want=BISHOP, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}

func TestKnight(t *testing.T) {
	var p *Knight
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewKnight(WHITE, pos)
	if p.Color() != WHITE {
		t.Errorf("Wrong color, want=WHITE, got=%s", p.Color())
	}
	if p.Literal() != KNIGHT {
		t.Errorf("Wrong literal, want=KNIGHT, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}

func TestRook(t *testing.T) {
	var p *Rook
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewRook(WHITE, pos)
	if p.Color() != WHITE {
		t.Errorf("Wrong color, want=WHITE, got=%s", p.Color())
	}
	if p.Literal() != ROOK {
		t.Errorf("Wrong literal, want=ROOK, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}

func TestPawn(t *testing.T) {
	var p *Pawn
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewPawn(WHITE, pos)
	if p.Color() != WHITE {
		t.Errorf("Wrong color, want=WHITE, got=%s", p.Color())
	}
	if p.Literal() != PAWN {
		t.Errorf("Wrong literal, want=PAWN, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}

func TestNoPiece(t *testing.T) {
	var p *NoPiece
	var pos Position = Position{row: 0, col: 0, boardRow: '8', boardCol: 'a'}
	p = NewNoPiece(WHITE, pos)
	if p.Color() != WHITE {
		t.Errorf("Wrong color, want=WHITE, got=%s", p.Color())
	}
	if p.Literal() != NO_PIECE {
		t.Errorf("Wrong literal, want=NO_PIECE, got=%s", p.Literal())
	}
	if p.Position() != pos {
		t.Errorf("Wrong position, want=%v, got=%v", pos, p.Position())
	}
	if p.IsSelected() != false {
		t.Errorf("Wrong is selected, want=false, got=%v", p.IsSelected())
	}
	p.SetIsSelected(true)
	if p.IsSelected() != true {
		t.Errorf("Wrong is selected, want=true, got=%v", p.IsSelected())
	}
}
