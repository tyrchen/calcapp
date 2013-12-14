package calc

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
	StopValue  Value
	StopCol    Value
	CurrentCol Value
	Last       Point
	Bp         Bpoint
}

func (self *Env) init() *Env {
	self.StopCol = STOP_COL
	self.StopValue = STOP_VALUE
	return self
}
