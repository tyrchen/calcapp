package calcv6

import (
	. "calcapp/utils"
)

/*
 * increase with 1,3,7 until stop value
 */

func calc137(env *Env) (ret Point) {
	ret.T = false
	last := env.Last
	if last.T {
		ret.V = 1
	} else {
		ret.V = last.V*2 + 1

		if ret.V > STOP_VALUE {
			ret.V = 1
		}
	}

	// speical handling in STOP_COL

	if STOP_COL != 0 && env.Stop && env.CurrentCol >= STOP_COL &&
		(ret.V > STOP_VALUE ||
			last.V == 0 ||
			int(Abs(last.V)) == STOP_VALUE ||
			last.T == true) {
		ret.V = 0
	}
	return
}

/*
 * 看前一列上一行带z与本列上一行相同，不带z相反
 */

func calcWithSign(env *Env) (ret Point) {
	ret = calc137(env)
	val := Abs(ret.V)
	sign := Sign(env.S2)
	if env.S1.T {
		ret.V = val * sign
	} else {
		ret.V = val * sign * -1
	}
	return
}

func calcGz(v Value) (ret Value) {
	if v > 0 {
		ret = 1
	} else {
		ret = -1
	}
	return
}
