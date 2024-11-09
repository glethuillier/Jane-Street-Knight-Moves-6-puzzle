package board

import "fmt"

type Square string

const (
	A Square = "A"
	B        = "B"
	C        = "C"
)

type Coordinates struct {
	File rune // 'a' to 'f'
	Rank int  // 1 to 6
}

func (c Coordinates) ToString() string {
	return fmt.Sprintf("%c%d", c.File, c.Rank)
}
