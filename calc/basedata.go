package calc

import (
	"fmt"
	"strings"
)

type BaseData struct {
	Inst  [COLS]Bpoint
	Bp    [COLS]Bpoint
	Nbp   [COLS]Bpoint
	Data  [ROWS][COLS]Point
	Ret   [COLS]Point
	Pret  [COLS]Point
	Nret  [COLS]Point
	Nret1 [COLS]Point
}

type BaseCalculator interface {
	Init()
	Calc(inst Bpoint, pos Value)
}

func (self *BaseData) Init() {
	// init the col 0
	for i := 0; i < ROWS; i++ {
		if i%2 == 0 {
			self.Data[i][0].V = 1
		} else {
			self.Data[i][0].V = -1
		}

	}
}

func (self *BaseData) LoadBp(bp []uint8) {
	for i := 0; i < COLS; i++ {
		self.Bp[i] = Bpoint(bp[i])
	}
}

func (self *BaseData) WithZ(pos Value, inst Bpoint) {
	for i := 0; i < ROWS; i++ {
		withZ(&self.Data[i][pos], inst)
	}
}

func (self *BaseData) Calc(inst Bpoint, pos Value) {
	self.Inst[pos] = inst
	self.Nbp[pos] = getNextBp(self.Bp[pos], inst)

	// withZ for current pos
	self.WithZ(pos, inst)

	// calculate next pos

	var env = Env{}
	next_pos := pos + 1

	if next_pos >= COLS {
		return
	}

	for i := 0; i < ROWS; i++ {
		if i == 0 {
			env.Bp = self.Bp[next_pos]
		} else {
			env.Bp = pointToBp(self.Data[i-1][next_pos])
		}

		env.CurrentCol = next_pos
		env.Last = self.Data[i][pos]
		self.Data[i][next_pos] = calcReverse(&env)
	}
}

func (self *BaseData) String() string {
	printBp := func(title string, val []Bpoint) string {
		return fmt.Sprintf("%-10s %12s %12s %12s %12s %12s\n",
			title, val[0].String(), val[1].String(), val[2].String(),
			val[3].String(), val[4].String())
	}
	printPoint := func(title string, val []Point) string {
		return fmt.Sprintf("%-10s %12s %12s %12s %12s %12s\n",
			title, val[0].String(), val[1].String(), val[2].String(),
			val[3].String(), val[4].String())
	}

	arr := make([]string, 0)
	arr = append(arr, printBp("Inst", self.Inst[:]))
	arr = append(arr, printBp("BP", self.Bp[:]))

	for i := 0; i < ROWS; i++ {
		arr = append(arr, printPoint(string(i), self.Data[i][:]))
	}
	arr = append(arr, printBp("New BP", self.Nbp[:]))

	return strings.Join(arr, "")
}
