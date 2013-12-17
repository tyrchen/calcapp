package calc

import (
	"math"
)

func Copysign(value, sign Value) Value {
	return Value(math.Copysign(float64(value), float64(sign)))
}

func Abs(v Value) Value {
	return Value(math.Abs(float64(v)))
}

/*
 * increase with 1,3,7 until stop value
 */

func calc137(env *Env) (ret Point) {
	ret.T = false
	last := env.Last
	if last.T {
		ret.V = 1
	} else {
		ret.V = Abs(last.V)*2 + 1
		if env.CurrentCol >= STOP_COL &&
			(ret.V > STOP_VALUE ||
				last.V == 0 ||
				last.T == true) {
			ret.V = 0
		}

		if ret.V > STOP_VALUE {
			ret.V = 1
		}
	}
	return
}

func calc137WithSign(env *Env) (ret Point) {
	ret.T = false
	v := Copysign(1, env.Last.V)
	last := env.Last
	if last.T {
		ret.V = v
	} else {
		ret.V = Copysign(Abs(last.V)*2+1, last.V)

		if env.CurrentCol >= STOP_COL &&
			(ret.V > STOP_VALUE ||
				last.V == 0 ||
				last.T == true) {
			ret.V = 0
		}

		if ret.V > STOP_VALUE {
			ret.V = v
		}
	}
	return
}

/*
 * reverse sign if previous.t is true, not reverse otherwise
 */
func calcReverse(env *Env) (ret Point) {
	sign := bsign(env.Bp)
	ret = calc137(env)
	val := Abs(ret.V)
	if env.Last.T {
		ret.V = val * sign * -1
	} else {
		ret.V = val * sign
	}
	return
}

/*
 * sign of bp: 0: -1, 1: 1
 */
func bsign(bp Bpoint) Value {
	if bp > 0 {
		return 1
	} else {
		return -1
	}
}

func pointToBp(p Point) Bpoint {
	if p.T {
		return 1
	}
	return 0
}

func getNextBp(bp, inst Bpoint) Bpoint {
	return bp ^ inst
}

func withZ(p *Point, inst Bpoint) {
	if (p.V > 0 && inst == 1) || (p.V < 0 && inst == 0) {
		p.T = true
	} else {
		p.T = false
	}
}
