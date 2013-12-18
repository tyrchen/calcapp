package calc

import (
	"fmt"
	"strconv"
)

const (
	COLS       = 55 + 1
	ROWS       = 9
	GROUP_SIZE = 37000
	STOP_COL   = 36
	STOP_VALUE = 1048575

	// for concurrency
	CHUNKS     = 100
	CHUNK_SIZE = GROUP_SIZE / CHUNKS
)

type Bpoint uint8
type Value int
type Duck interface{}

type Point struct {
	T bool
	V Value
}

type Env struct {
	CurrentCol Value
	Last       Point
	Bp         Bpoint
}

var (
	Values GroupData
)

func (self Bpoint) String() string {
	return strconv.Itoa(int(self))
}

func (self Point) String() string {
	if self.T {
		return fmt.Sprintf("z%d", self.V)
	} else {
		return fmt.Sprintf("%d", self.V)
	}
}
