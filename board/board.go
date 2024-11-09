package board

import (
	"errors"
	"fmt"
)

type ChessBoard struct {
	Size    int
	Squares map[Coordinates]Square
}

func NewChessBoard(board [][]string) (*ChessBoard, error) {
	size := len(board)
	if size != len(board[0]) {
		return nil, errors.New("invalid board size")
	}

	b := &ChessBoard{
		Size:    size,
		Squares: make(map[Coordinates]Square),
	}

	for x, row := range board {
		for y, cell := range row {
			b.Squares[Coordinates{rune(y + 'a'), 1 + int(x)}] = Square(cell)
		}
	}

	return b, nil
}

func (cb *ChessBoard) GetSize() int {
	return cb.Size
}

func (cb *ChessBoard) IsInbound(c Coordinates) bool {
	s := c.Rank >= 1 && c.Rank <= cb.Size && c.File >= 'a' && c.File <= 'a'+rune(cb.Size-1)
	return s
}

func (cb *ChessBoard) GetSquare(coordinate Coordinates) Square {
	square, _ := cb.Squares[coordinate]
	return square
}

func (cb *ChessBoard) displayRankLine(rank int) {
	for file := 'a'; file < 'a'+rune(cb.Size); file++ {
		coordinate := Coordinates{File: file, Rank: rank}
		if square, ok := cb.Squares[coordinate]; ok {
			fmt.Printf("%s ", square)
		} else {
			fmt.Print(". ")
		}
	}
	fmt.Println()
}

func (cb *ChessBoard) Display() {
	for rank := cb.Size; rank >= 1; rank-- {
		fmt.Printf("%d ", rank)
		cb.displayRankLine(rank)
	}
	fmt.Print("  ")
	for file := 'a'; file < 'a'+rune(cb.Size); file++ {
		fmt.Printf("%c ", file)
	}
	fmt.Println()
}
