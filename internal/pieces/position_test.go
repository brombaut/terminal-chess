package pieces

import "testing"

func TestValidPositions(t *testing.T) {
	squares := []struct {
		rowId    int
		colId    int
		boardRow rune
		boardCol rune
	}{
		{0, 0, '8', 'a'},
		{0, 1, '8', 'b'},
		{0, 2, '8', 'c'},
		{0, 3, '8', 'd'},
		{0, 4, '8', 'e'},
		{0, 5, '8', 'f'},
		{0, 6, '8', 'g'},
		{0, 7, '8', 'h'},
		{1, 0, '7', 'a'},
		{1, 1, '7', 'b'},
		{1, 2, '7', 'c'},
		{1, 3, '7', 'd'},
		{1, 4, '7', 'e'},
		{1, 5, '7', 'f'},
		{1, 6, '7', 'g'},
		{1, 7, '7', 'h'},
		{2, 0, '6', 'a'},
		{2, 1, '6', 'b'},
		{2, 2, '6', 'c'},
		{2, 3, '6', 'd'},
		{2, 4, '6', 'e'},
		{2, 5, '6', 'f'},
		{2, 6, '6', 'g'},
		{2, 7, '6', 'h'},
		{3, 0, '5', 'a'},
		{3, 1, '5', 'b'},
		{3, 2, '5', 'c'},
		{3, 3, '5', 'd'},
		{3, 4, '5', 'e'},
		{3, 5, '5', 'f'},
		{3, 6, '5', 'g'},
		{3, 7, '5', 'h'},
		{4, 0, '4', 'a'},
		{4, 1, '4', 'b'},
		{4, 2, '4', 'c'},
		{4, 3, '4', 'd'},
		{4, 4, '4', 'e'},
		{4, 5, '4', 'f'},
		{4, 6, '4', 'g'},
		{4, 7, '4', 'h'},
		{5, 0, '3', 'a'},
		{5, 1, '3', 'b'},
		{5, 2, '3', 'c'},
		{5, 3, '3', 'd'},
		{5, 4, '3', 'e'},
		{5, 5, '3', 'f'},
		{5, 6, '3', 'g'},
		{5, 7, '3', 'h'},
		{6, 0, '2', 'a'},
		{6, 1, '2', 'b'},
		{6, 2, '2', 'c'},
		{6, 3, '2', 'd'},
		{6, 4, '2', 'e'},
		{6, 5, '2', 'f'},
		{6, 6, '2', 'g'},
		{6, 7, '2', 'h'},
		{7, 0, '1', 'a'},
		{7, 1, '1', 'b'},
		{7, 2, '1', 'c'},
		{7, 3, '1', 'd'},
		{7, 4, '1', 'e'},
		{7, 5, '1', 'f'},
		{7, 6, '1', 'g'},
		{7, 7, '1', 'h'},
	}
	for _, square := range squares {
		var err error
		var fromId *Position
		fromId, err = NewIdxPosition(square.rowId, square.colId)
		if err != nil {
			t.Errorf("Err building from Idx: %s", err.Error())
		}
		if fromId.row != square.rowId {
			t.Errorf("Err building from Idx, want rowId=%d, got=%d", square.rowId, fromId.row)
		}
		if fromId.col != square.colId {
			t.Errorf("Err building from Idx, want colId=%d, got=%d", square.colId, fromId.col)
		}
		if fromId.boardRow != square.boardRow {
			t.Errorf("Err building from Idx, want boardRow=%c, got=%c", square.boardRow, fromId.boardRow)
		}
		if fromId.boardCol != square.boardCol {
			t.Errorf("Err building from Idx, want boardCol=%c, got=%c", square.boardCol, fromId.boardCol)
		}

		var fromBoard *Position
		fromBoard, err = NewBoardPosition(square.boardRow, square.boardCol)
		if err != nil {
			t.Errorf("Err building from Idx: %s", err.Error())
		}
		if fromBoard.row != square.rowId {
			t.Errorf("Err building from Idx, want rowId=%d, got=%d", square.rowId, fromBoard.row)
		}
		if fromBoard.col != square.colId {
			t.Errorf("Err building from Idx, want colId=%d, got=%d", square.colId, fromBoard.col)
		}
		if fromBoard.boardRow != square.boardRow {
			t.Errorf("Err building from Idx, want boardRow=%c, got=%c", square.boardRow, fromBoard.boardRow)
		}
		if fromBoard.boardCol != square.boardCol {
			t.Errorf("Err building from Idx, want boardCol=%c, got=%c", square.boardCol, fromBoard.boardCol)
		}
	}
}

func TestInvalidPositionsFromIdx(t *testing.T) {
	squares := []struct {
		rowId    int
		colId    int
		boardRow rune
		boardCol rune
	}{
		{-1, 0, '8', 'a'},
		{8, 0, '8', 'a'},
		{0, -1, '8', 'a'},
		{0, 8, '8', 'a'},
	}
	for _, square := range squares {
		var err error
		_, err = NewIdxPosition(square.rowId, square.colId)
		if err == nil {
			t.Errorf("Was expecting error when building position from ids, but got none: rowId=%d, colId=%d", square.rowId, square.colId)
		}
	}
}

func TestInvalidPositionsFromBoard(t *testing.T) {
	squares := []struct {
		rowId    int
		colId    int
		boardRow rune
		boardCol rune
	}{
		{0, 0, '9', 'a'},
		{0, 0, '0', 'a'},
		{0, 0, '8', '0'},
		{0, 0, '8', 'i'},
	}
	for _, square := range squares {
		var err error
		_, err = NewBoardPosition(square.boardRow, square.boardCol)
		if err == nil {
			t.Errorf("Was expecting error when building position from board, but got none: boardRow=%d, boardCol=%d", square.boardRow, square.boardCol)
		}
	}
}
