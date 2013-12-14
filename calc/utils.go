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
	var v = Copysign(1, env.Last.V)
	if env.Last.T {
		ret.V = v
	} else {
		ret.V = Copysign(Abs(env.Last.V)*2+1, env.Last.V)
		if ret.V > env.StopValue {
			ret.V = v
		}
	}
	return
}

/*
 * reverse sign if previous.t is true, not reverse otherwise
 */
func calcReverse(env *Env) (ret Point) {
	var sign = bsign(env.Bp)
	ret = calc137(env)
	var val = Abs(ret.V)
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

func getNextBp(bp, inst Bpoint) Bpoint {
	if bp == inst {
		return 1
	} else {
		return 0
	}
}
