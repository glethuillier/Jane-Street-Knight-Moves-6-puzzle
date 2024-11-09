package board

import (
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	cases := []struct {
		boardInput [][]string
		wantErr    bool
	}{
		{[][]string{{"a1", "b1"}, {"a2", "b2"}}, false},
		{[][]string{{"a1", "b1", "c1"}, {"a2", "b2"}}, true},
	}

	for _, tc := range cases {
		_, err := NewChessBoard(tc.boardInput)
		if (err != nil) != tc.wantErr {
			t.Errorf("NewChessBoard(%v) returned error %q, want error? %v", tc.boardInput, err, tc.wantErr)
		}
	}
}

func TestChessBoard_GetSize(t *testing.T) {
	board := &ChessBoard{Size: 4}
	size := board.GetSize()
	if size != 4 {
		t.Errorf("GetSize() returned %d, want %d", size, 4)
	}
}

func TestChessBoard_IsInbound(t *testing.T) {
	board := &ChessBoard{Size: 4}
	cases := []struct {
		in   Coordinates
		want bool
	}{
		{Coordinates{'a', 1}, true},
		{Coordinates{'d', 4}, true},
		{Coordinates{'e', 5}, false},
		{Coordinates{'a', 0}, false},
		{Coordinates{'a', 5}, false},
		{Coordinates{'z', 1}, false},
	}

	for _, tc := range cases {
		got := board.IsInbound(tc.in)
		if got != tc.want {
			t.Errorf("IsInbound(%s) returned %v, want %v", tc.in.ToString(), got, tc.want)
		}
	}
}

func TestChessBoard_GetSquare(t *testing.T) {
	squares := make(map[Coordinates]Square)
	squares[Coordinates{'a', 1}] = A
	board := &ChessBoard{
		Size:    1,
		Squares: squares,
	}

	got := board.GetSquare(Coordinates{'a', 1})
	want := A

	if got != want {
		t.Errorf("GetSquare() returned %q, want %q", got, want)
	}
}
