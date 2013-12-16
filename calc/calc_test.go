package calc

import (
	"testing"
)

func getEnv(bp Bpoint, currentCol Value, last Point) (env *Env) {
	env = new(Env)
	env.Bp = bp
	env.CurrentCol = currentCol
	env.Last = last
	return
}

func testCalcReverse(t *testing.T, env *Env, expected Point) {
	ret := calcReverse(env)
	if ret.T == expected.T && ret.V == expected.V {
		t.Logf("calcReverse(bp=%d, col=%d, last=%s) = %s",
			env.Bp, env.CurrentCol, env.Last.ToString(), expected.ToString())
	} else {
		t.Errorf("Error: calcReverse(bp=%d, col=%d, last=%s, expected=%s) = %s",
			env.Bp, env.CurrentCol, env.Last.ToString(), expected.ToString(), ret.ToString())
	}
}

func TestCalcReverse(t *testing.T) {
	// test increase with x*2+1
	env := getEnv(1, 1, Point{false, 15})
	testCalcReverse(t, env, Point{false, 31})

	env = getEnv(0, 1, Point{true, -15})
	testCalcReverse(t, env, Point{false, 1})
}
