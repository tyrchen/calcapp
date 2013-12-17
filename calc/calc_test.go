package calc

import (
	"calcapp/utils"
	"fmt"
	"testing"
)

const (
	BP_FOR_TEST = "10110111110011110101101011010101101111100001111010110001"
)

var (
	insts = [...]string{
		"11111111111111111111111111111111111111111111111111111111",
		"00000000000000000000000000000000000000000000000000000000",
		"10101010101010101010101010101010101010101010101010101010",
	}
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
		t.Logf("calcReverse(bp=%d, col=%d, last=%v) = %v",
			env.Bp, env.CurrentCol, env.Last.String(), expected.String())
	} else {
		t.Errorf("Error: calcReverse(bp=%d, col=%d, last=%s, expected=%s) = %s",
			env.Bp, env.CurrentCol, env.Last, expected.String(), ret.String())
	}
}

func testBaseDataCalc(t *testing.T, data BaseData, insts string) {
	data.LoadBp(utils.StringToBp(BP_FOR_TEST))
	data.Init()
	for i, v := range utils.StringToBp(insts) {
		data.Calc(Bpoint(v), Value(i))
	}

	fmt.Println(data.String())
	t.Errorf("Error")

}

func TestCalcReverse(t *testing.T) {
	// test increase with x*2+1
	env := getEnv(1, 1, Point{false, 15})
	testCalcReverse(t, env, Point{false, 31})

	env = getEnv(0, 1, Point{true, -15})
	testCalcReverse(t, env, Point{false, 1})
}

func TestBaseDataCalc(t *testing.T) {
	var data BaseData
	for i := 0; i < len(insts); i++ {
		testBaseDataCalc(t, data, insts[i])
	}
}
