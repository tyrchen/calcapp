package calc

import (
	"fmt"
)

const (
	COLS       = 55 + 1
	ROWS       = 9
	GROUP_SIZE = 5760
	STOP_COL   = 36
	STOP_VALUE = 1048575
)

type Bpoint uint8
type Value int

type Point struct {
	T bool
	V Value
}

type Env struct {
	CurrentCol Value
	Last       Point
	Base       Point
	Bp         Bpoint
}

func (self *Point) ToString() string {
	if self.T {
		return fmt.Sprintf("z%d", self.V)
	} else {
		return fmt.Sprintf("%d", self.V)
	}
}
