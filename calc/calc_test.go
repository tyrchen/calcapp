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
	/*
		data.Clear()
		fmt.Println(data.Data[GROUP_SIZE-1])
		fmt.Println(data.String())

		data.LoadBp(1)
		data.Init()
		for i, v := range utils.StringToBp(INST_FOR_TEST) {
			data.Run(Bpoint(v), Value(i))
		}
		fmt.Println(data.Data[GROUP_SIZE-1])
		fmt.Println(data.String())
	*/
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
