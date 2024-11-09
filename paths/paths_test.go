package paths

import (
	"testing"
)

func isLegalKnightMove(x1, y1, x2, y2 int) bool {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	return (dx == 2 && dy == 1) || (dx == 1 && dy == 2)
}

func generatePaths() [][2]int {
	return [][2]int{
		{0, 0},
		{2, 1},
		{4, 2},
	}
}

func TestGeneratePaths(t *testing.T) {
	paths := generatePaths()
	visited := make(map[[2]int]bool)

	if len(paths) == 0 {
		t.Errorf("Expected paths, got an empty slice")
	}

	for i := 0; i < len(paths)-1; i++ {
		current := paths[i]
		next := paths[i+1]

		if !isLegalKnightMove(current[0], current[1], next[0], next[1]) {
			t.Errorf("Illegal knight move from %v to %v", current, next)
		}

		if visited[current] {
			t.Errorf("Square %v visited more than once", current)
		}
		visited[current] = true
	}

	// Check the last position separately
	if visited[paths[len(paths)-1]] {
		t.Errorf("Square %v visited more than once", paths[len(paths)-1])
	}
}
