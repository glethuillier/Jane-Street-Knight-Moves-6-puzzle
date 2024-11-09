package paths

import "jane/board"

type PathType string

const (
	PathA1ToF6 PathType = "a1_f6"
	PathA6ToF1 PathType = "a6_f1"
)

type Path struct {
	Type    PathType
	Squares []board.Coordinates
}

func (p Path) ToString() string {
	var result string
	for i, coordinates := range p.Squares {
		if i > 0 {
			result += ","
		}
		result += coordinates.ToString()
	}
	return result
}
