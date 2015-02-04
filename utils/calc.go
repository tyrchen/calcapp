package utils

import (
	"math"
	"strconv"
	"fmt"
)

type Bpoint uint8
type Value int16
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

func calc_fold(v1, v2, v3 Point) (r Point) {
	r.V = v1.V + v2.V + v3.V
	if r.V > 1 {
		r.V = 1
	} else if r.V < -1 {
		r.V = -1
	}
	return
}

func calc_reduce(part []Point) Point {
	l := len(part) / 3
	if l == 0 {
		panic("cannot be 0")
	}
	
	if l == 1 {
		return calc_fold(part[0], part[1], part[2])
	} else {
		return calc_fold(
			calc_reduce(part[:l]),
			calc_reduce(part[l:l*2]),
			calc_reduce(part[l*2:l*3]))
	}
}

func CalcReduce(data []Point) (ret [4]Point) {
	l := len(data)
	if l == 0 || l % 3 != 0 {
		panic("chunk size should be multiple of 3")
	}
	l = l / 3

	ret[1] = calc_reduce(data[:l])
	ret[2] = calc_reduce(data[l:l*2])
	ret[3] = calc_reduce(data[l*2:l*3])
	ret[0] = calc_fold(ret[1],ret[2], ret[3])
	return
}

