package utils

import (
	"math"
	"strconv"
	"fmt"
)

type Bpoint uint8
type Value int
type Point struct {
	T bool
	V Value
}


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

func Copysign(value, sign Value) Value {
	return Value(math.Copysign(float64(value), float64(sign)))
}

func Abs(v Value) Value {
	return Value(math.Abs(float64(v)))
}

func SignFollow(value Value, bp Bpoint) (ret Value) {
	sign := Bsign(bp)
	ret = sign * value
	return ret
}

/*
 * sign of bp: 0: -1, 1: 1
 */
func Bsign(bp Bpoint) Value {
	if bp > 0 {
		return 1
	} else {
		return -1
	}
}

func Sign(val Point) Value {
	if val.V > 0 {
		return 1
	} else {
		return -1
	}
}

func Zsign(val Point) Value {
	if val.T {
		return 1
	} else {
		return -1
	}
}

func PointToBp(p Point) Bpoint {
	if p.V > 0 {
		return 1
	}
	return 0
}

func GetNextBp(bp, inst Bpoint) Bpoint {
	return bp ^ ^inst + 2
}

func WithZ(p *Point, inst Bpoint) {
	if (p.V > 0 && inst == 1) || (p.V < 0 && inst == 0) {
		p.T = true
	} else {
		p.T = false
	}
}

func WithZbp(p *Point, bp, inst Bpoint) {
	if (bp == inst) {
		p.V = 1
		p.T = true
	} else {
		p.V = 0
		p.T = false
	}
}

