package calc

import (
	"calcapp/utils"
	"fmt"
	"runtime"
	"testing"
)

const (
	BP_FOR_TEST   = "10001010100111101100000111010100100011111100011100111111"
	INST_FOR_TEST = "10011111110000101101101010011111100110010010011111000000"
	INST_ALL_ONE  = "11111111111111111111111111111111111111111111111111111111"
	INST_ALL_ZERO = "00000000000000000000000000000000000000000000000000000000"
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
			env.Bp, env.CurrentCol, env.Last, expected)
	} else {
		t.Errorf("Error: calcReverse(bp=%d, col=%d, last=%s, expected=%s) = %s",
			env.Bp, env.CurrentCol, env.Last, expected, ret)
	}
}

func testBaseDataCalc(t *testing.T, data BaseData, insts string) {
	data.LoadBp(utils.StringToBp(BP_FOR_TEST))
	data.Init()
	for i, v := range utils.StringToBp(insts) {
		data.Run(Bpoint(v), Value(i))
	}

	//fmt.Println(data)

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
	testBaseDataCalc(t, data, INST_FOR_TEST)
}

func TestGroupCalc(t *testing.T) {
	data := new(GroupData)
	runtime.GOMAXPROCS(8)
	data.Init(0)
	for i, v := range utils.StringToBp(INST_FOR_TEST) {
		data.Run(Bpoint(v), Value(i))
	}
	//fmt.Println(data.Data[GROUP_SIZE-1])
	fmt.Println(data.String())
}

func TestMp(t *testing.T) {
	data := new(GroupData)
	runtime.GOMAXPROCS(8)

	for g := 1; g < 2; g++ {
		data.Init(uint(g))
		for i, v := range utils.StringToBp(INST_FOR_TEST) {
			data.Run(Bpoint(v), Value(i))
		}
	}
}

func TestBigData(t *testing.T) {
	data := new(BigData)
	runtime.GOMAXPROCS(8)
	data.Init()
	col := 0
	for i, v := range utils.StringToBp(INST_FOR_TEST) {
		data.Run(Bpoint(v), Value(i))
		col++
		if col >= COLS-1 {
			break
		}
	}
	fmt.Println(data.String())
}

func testThreeSome(t *testing.T, up [COLS]Point, inst string) {
	data := new(ThreeSome)
	data.Up = up
	data.Init(up[0])
	col := 0
	for i, v := range utils.StringToBp(inst) {
		data.Run(Bpoint(v), Value(i))
		col++
		if col >= COLS-1 {
			break
		}
	}
	fmt.Println(data.String())
}

func TestThreeSome(t *testing.T) {
	var up [COLS]Point

	up = [COLS]Point{
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20}, {false, 10}, {true, -20}, {false, 10}, {true, -20},
		{false, 10}, {true, -20},
	}
	testThreeSome(t, up, INST_FOR_TEST)

	up = [COLS]Point{
		{true, 10}, {false, 20}, {true, 10}, {true, 20}, {false, 10}, {true, 20},
		{true, 10}, {false, 20}, {true, 10}, {true, 20}, {false, 10}, {true, 20},
		{true, 10}, {false, 20}, {true, 10}, {true, 20}, {false, 10}, {true, 20},
		{true, 10}, {false, 20}, {true, 10}, {true, 20}, {false, 10}, {true, 20},
		{true, 10}, {false, 20}, {true, 10}, {true, 20}, {false, 10}, {true, 20},
		{true, 10}, {true, -20}, {true, 10}, {true, -20}, {true, 10}, {true, -20},
		{true, 10}, {true, -20}, {true, 10}, {true, -20}, {true, 10}, {true, -20},
		{true, 10}, {true, -20}, {true, 10}, {true, -20}, {true, 10}, {true, -20},
		{true, 10}, {true, -20}, {true, 10}, {true, -20}, {true, 10}, {true, -20},
		{false, 10}, {true, -20},
	}
	testThreeSome(t, up, INST_ALL_ONE)
	testThreeSome(t, up, INST_ALL_ZERO)

}
