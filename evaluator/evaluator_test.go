package evaluator

import (
	"jane/board"
	"jane/paths"
	"testing"
)

func TestGetValues(t *testing.T) {
	e := NewEvaluator(&board.ChessBoard{})
	values := e.GetValues()
	if len(values) == 0 {
		t.Errorf("No values generated in evaluator")
	}
}

func TestAddToSolutionsAndAssess(t *testing.T) {
	tests := []struct {
		name    string
		path    *paths.Path
		val     Values
		wantLen int
		wantErr bool
	}{
		{
			name: "ValidPath1",
			path: &paths.Path{Squares: []board.Coordinates{
				{File: 'a', Rank: 1},
				{File: 'b', Rank: 2},
				{File: 'c', Rank: 3},
			}},
			val:     Values{1, 1, 1},
			wantLen: 1,
		},
		{
			name: "ValidPath2",
			path: &paths.Path{Squares: []board.Coordinates{
				{File: 'd', Rank: 4},
				{File: 'e', Rank: 5},
				{File: 'f', Rank: 6},
			}},
			val:     Values{2, 2, 2},
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		e := NewEvaluator(&board.ChessBoard{})
		e.addToSolutionsAndAssess(&tt.val, tt.path)
		if gotLen := len(e.pathMap[tt.val]); gotLen != tt.wantLen {
			t.Errorf("TestAddToSolutionsAndAssess(%s): expected pathMap length %v, got %v", tt.name, tt.wantLen, gotLen)
		}
	}
}

func TestCalculateScore(t *testing.T) {
	e := NewEvaluator(&board.ChessBoard{})
	tests := []struct {
		name          string
		currentScore  uint
		values        Values
		prevCoord     board.Coordinates
		currCoord     board.Coordinates
		expectedScore uint
	}{
		{
			name:          "SquareA",
			currentScore:  2,
			values:        Values{2, 2, 2},
			prevCoord:     board.Coordinates{File: 'a', Rank: 1},
			currCoord:     board.Coordinates{File: 'b', Rank: 2},
			expectedScore: 2,
		},
		{
			name:          "SquareB",
			currentScore:  3,
			values:        Values{3, 3, 3},
			prevCoord:     board.Coordinates{File: 'c', Rank: 3},
			currCoord:     board.Coordinates{File: 'd', Rank: 4},
			expectedScore: 3,
		},
	}

	for _, tt := range tests {
		score := e.calculateScore(tt.currentScore, &tt.values, tt.prevCoord, tt.currCoord)
		if score != tt.expectedScore {
			t.Errorf("TestCalculateScore(%s): expected %v, got %v", tt.name, tt.expectedScore, score)
		}
	}
}

func TestPruneValues(t *testing.T) {
	e := NewEvaluator(&board.ChessBoard{})
	e.pruneValues(4)
	vals := e.GetValues()
	for _, v := range vals {
		if v.A+v.B+v.C > 4 {
			t.Errorf("Values not pruned correctly")
		}
	}
}
