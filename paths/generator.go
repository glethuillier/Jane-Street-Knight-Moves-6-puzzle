package paths

import (
	"crypto/rand"
	"errors"
	"jane/board"
	"math/big"
)

var (
	// "two corner-to-corner trips — one from a1 to f6, and the other from a6 to f1"
	pathCorners = map[PathType][2]board.Coordinates{
		PathA1ToF6: {{File: 'a', Rank: 1}, {File: 'f', Rank: 6}},
		PathA6ToF1: {{File: 'a', Rank: 6}, {File: 'f', Rank: 1}},
	}

	// legal knight moves
	knightMoves = []board.Coordinates{
		{File: 2, Rank: 1}, {File: 2, Rank: -1},
		{File: -2, Rank: 1}, {File: -2, Rank: -1},
		{File: 1, Rank: 2}, {File: 1, Rank: -2},
		{File: -1, Rank: 2}, {File: -1, Rank: -2},
	}
)

type PathGenerator struct {
	board         *board.ChessBoard
	pathEndpoints [2]board.Coordinates
}

func NewPathGenerator(board *board.ChessBoard, pathType PathType) (*PathGenerator, error) {
	endpoints, ok := pathCorners[pathType]
	if !ok {
		return nil, errors.New("invalid path type")
	}
	return &PathGenerator{
		board:         board,
		pathEndpoints: endpoints,
	}, nil
}

func (pg *PathGenerator) GeneratePaths(pathType PathType) (Path, error) {
	from, to := pg.pathEndpoints[0], pg.pathEndpoints[1]

	// initialize a map to track visited coordinates and a variable to store the path
	visited := make(map[board.Coordinates]bool)
	path := Path{Type: pathType, Squares: []board.Coordinates{from}}
	visited[from] = true

	// Use a loop with tracking the current position
	for current := from; current != to; {
		// get all possible moves from the current position that are not visited
		possibleMoves := pg.getPossibleMoves(&current, visited)
		if len(possibleMoves) == 0 {
			return Path{}, errors.New("no path found")
		}

		// move to a randomly chosen position from the possible moves
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(possibleMoves))))
		current = possibleMoves[randomIndex.Int64()]
		visited[current] = true
		path.Squares = append(path.Squares, current)
	}

	return path, nil
}

func (pg *PathGenerator) getPossibleMoves(
	current *board.Coordinates,
	visited map[board.Coordinates]bool,
) []board.Coordinates {
	var possibleMoves []board.Coordinates
	// iterate over each knight move
	for _, move := range knightMoves {
		// calculate the next possible move by adding current position and move
		next := board.Coordinates{File: current.File + move.File, Rank: current.Rank + move.Rank}

		// check if the next move is within the bounds of the board and not visited
		if pg.board.IsInbound(next) && !visited[next] {
			// if so, append the move to the list of possible moves
			possibleMoves = append(possibleMoves, next)
		}
	}

	return possibleMoves
}
