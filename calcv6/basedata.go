package calcv6

import (
	. "calcapp/utils"
	//"strconv"
	"strings"
	// "fmt"
)

var (
	BP = strings.Repeat("11111110", 7)
)

func (self *BaseData) Init() {
	// init the col 0
	//self.LoadBp(BP)
	sign := Bsign(self.Bp[0])

	for i := 0; i < ROWS; i++ {
		self.Data[i][0].V = 1 * sign
	}

	self.Ag[0].V = ROWS * sign
	self.G1[0].V = 1 * sign
	self.G137[0].V = self.G1[0].V
}

/*
func (self *BaseData) LoadBp(bp string) {
	for i := 0; i < COLS - 1; i++ {
		tmp, _ := strconv.Atoi(string(bp[i]))
		self.Bp[i] = Bpoint(tmp)
		self.Zbp[i].V = Value(tmp)
	}
}
*/

func (self *BaseData) LoadBp(bp []uint8) {
	for i := 0; i < COLS-1; i++ {
		self.Bp[i] = Bpoint(bp[i])
		self.Zbp[i].V = Value(bp[i])
	}
	self.Bp[COLS-1] = 0
}

func (self *BaseData) Run(inst Bpoint, pos Value) {

	self.withZ(inst, pos)

	self.calc(pos + 1)

}

func (self *BaseData) withZ(inst Bpoint, pos Value) {
	WithZbp(&self.Zbp[pos], self.Bp[pos], inst)

	for i := 0; i < ROWS; i++ {
		WithZ(&self.Data[i][pos], inst)
	}
	WithZ(&self.Ag[pos], inst)
	WithZ(&self.G1[pos], inst)
	WithZ(&self.G137[pos], inst)
}

func (self *BaseData) calc(pos Value) {
	var env = Env{}
	var tmp Point
	if pos >= COLS {
		return
	}

	// calculate values and xg
	env.CurrentCol = pos
	for i := 0; i < ROWS; i++ {
		if i == 0 {
			env.S1 = self.Zbp[pos-1]
			env.S2 = self.Zbp[pos]
		} else {
			env.S1 = self.Data[i-1][pos-1]
			env.S2 = self.Data[i-1][pos]
		}

		env.Last = self.Data[i][pos-1]
		tmp = calcWithSign(&env)
		self.Data[i][pos] = tmp
		self.Ag[pos].V += tmp.V
	}

	self.calcG1(pos)

	self.calcG137(pos)
}

func (self *BaseData) calcG1(pos Value) {
	if self.Ag[pos].V >= 0 {
		self.G1[pos].V = 1
	} else {
		self.G1[pos].V = -1
	}
}

func (self *BaseData) calcG137(pos Value) {
	last := self.G137[pos-1]
	up := self.G1[pos]
	sign := Sign(up)

	var ret Value

	if last.T {
		ret = 1 * sign
	} else {
		ret = Abs(last.V)*2 + 1
		if ret > 63 {
			ret = 1
		}
		ret *= sign
	}

	self.G137[pos].V = ret

	return
}