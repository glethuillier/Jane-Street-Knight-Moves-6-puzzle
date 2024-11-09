package evaluator

import (
	"jane/board"
	"jane/paths"
	"sync"
)

type Evaluator struct {
	mu      sync.RWMutex
	pathMap map[Values][]paths.Path
	known   map[string]bool
	board   *board.ChessBoard
	values  []Values
	best    *Best
}

func NewEvaluator(cb *board.ChessBoard) *Evaluator {
	return &Evaluator{
		pathMap: make(map[Values][]paths.Path),
		known:   make(map[string]bool),
		board:   cb,
		values:  generateAllValues(),
		best:    &Best{},
	}
}

func (e *Evaluator) Evaluate(values *Values, path *paths.Path, targetValue uint) {
	if len(path.Squares) == 0 {
		return
	}

	initialSquare := e.board.GetSquare(path.Squares[0])
	score := getSquareValue(initialSquare, values)

	for i := 1; i < len(path.Squares); i++ {
		score = e.calculateScore(score, values, path.Squares[i-1], path.Squares[i])
		if score > targetValue {
			return
		}
	}

	if score == targetValue {
		e.addToSolutionsAndAssess(values, path)
	}
}

// haveDifferentTypes if the paths are different (i.e., they use different corners)
func haveDifferentTypes(paths []paths.Path) bool {
	firstPathType := paths[0].Type
	for _, path := range paths {
		if path.Type != firstPathType {
			return true
		}
	}

	return false
}

// getShortestCandidates finds the shortest paths among the candidates
func getShortestCandidates(ps []paths.Path) (*paths.Path, *paths.Path) {
	var candidate0, candidate1 *paths.Path
	for _, path := range ps {
		switch path.Type {
		case paths.PathA1ToF6:
			if candidate0 == nil || len(path.Squares) < len(candidate0.Squares) {
				candidate0 = &path
			}
		case paths.PathA6ToF1:
			if candidate1 == nil || len(path.Squares) < len(candidate1.Squares) {
				candidate1 = &path
			}
		}
	}
	return candidate0, candidate1
}

// addToSolutionsAndAssess adds the current path to the set of solutions
// and assess if it exists a pair of paths for the given values that is
// the most optimal one
func (e *Evaluator) addToSolutionsAndAssess(values *Values, path *paths.Path) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// add current solution to the set of solutions for this combination of values
	e.pathMap[*values] = append(e.pathMap[*values], *path)

	ps := e.pathMap[*values]

	// paths using the same corners cannot lead to a solution
	if !haveDifferentTypes(ps) {
		return
	}

	// get the shortest candidates
	candidate0, candidate1 := getShortestCandidates(ps)
	if candidate0 == nil || candidate1 == nil {
		return
	}

	if err := e.best.CheckIsBest(e.board, values, candidate0, candidate1); err != nil {
		e.pruneValues(values.A + values.B + values.C)
	}
}

// pruneValues removes the values that sum above a threshold to converge quickly
// to a global minimum sum during the exploration
func (e *Evaluator) pruneValues(maxSum uint) {
	j := 0
	for _, v := range e.values {
		if v.A+v.B+v.C <= maxSum {
			e.values[j] = v
			j++
		}
	}
	e.values = e.values[:j]
}

func getSquareValue(square board.Square, values *Values) uint {
	switch square {
	case board.A:
		return values.A
	case board.B:
		return values.B
	case board.C:
		return values.C
	default:
		return 0
	}
}

func (e *Evaluator) calculateScore(
	currentScore uint,
	values *Values,
	prevCoordinates, currCoordinates board.Coordinates,
) uint {
	previousSquare := e.board.GetSquare(prevCoordinates)
	currentSquare := e.board.GetSquare(currCoordinates)
	nextValue := getSquareValue(currentSquare, values)

	// "if your move is between two different integers, multiply your score by the value you are moving to"
	if previousSquare != currentSquare {
		return currentScore * nextValue
	}

	// "otherwise, increment your score by the value you are moving to"
	return currentScore + nextValue
}
