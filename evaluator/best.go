package evaluator

import (
	"fmt"
	"github.com/pkg/errors"
	"jane/board"
	"jane/paths"
)

type Best struct {
	pathSize *int
	sum      *uint
}

// CheckIsBest checks if the current pair of paths is the most optimal both in terms of ABC values and paths lengths
func (b *Best) CheckIsBest(cb *board.ChessBoard, values *Values, candidate0, candidate1 *paths.Path) error {
	sum := values.A + values.B + values.C

	// if the sum is less than the current best one, set it as best
	if b.sum == nil || sum < *b.sum {
		b.sum = &sum
		b.pathSize = nil
	}

	if sum == *b.sum {
		size := len(candidate0.Squares) + len(candidate1.Squares)
		if b.pathSize == nil || size < *b.pathSize {
			b.pathSize = &size
			fmt.Printf("\nA: %d, B: %d, C: %d\n", values.A, values.B, values.C)
			fmt.Printf("(1) %s (2) %s\n", candidate0.ToString(), candidate1.ToString())
			fmt.Printf("Sum: %d / Paths length: %d\n", sum, size)

			err := SaveAsPNG(cb, values, candidate0, candidate1)
			if err != nil {
				return errors.Wrap(err, "failed to save PNG")
			}
		}

		return nil
	}

	return nil
}
