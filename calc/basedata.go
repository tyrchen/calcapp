package calc

import (
	"fmt"
	//"reflect"
	"strings"
)

type BaseData struct {
	Inst [COLS]Bpoint
	Bp   [COLS]Bpoint
	Nbp  [COLS]Bpoint
	Data [ROWS][COLS]Point
	Xg   [COLS]Point
	Gz   [COLS]Point
	Gf   [COLS]Point
	Gf1  [COLS]Point
}

type BaseCalculator interface {
	Init()
	Run(inst Bpoint, pos Value)
}

func (self *BaseData) Init() {
	// init the col 0
	sign := bsign(self.Bp[0])
	for i := 0; i < ROWS; i++ {
		if i%2 == 0 {
			self.Data[i][0].V = 1 * sign
		} else {
			self.Data[i][0].V = -1 * sign
		}
	}
	self.Xg[0].V = 1 * sign
	self.Gz[0].V = 1 * sign
	self.Gf[0].V = -1 * sign
	self.Gf1[0].V = -1 * sign
}

func (self *BaseData) LoadBp(bp []uint8) {
	for i := 0; i < COLS; i++ {
		self.Bp[i] = Bpoint(bp[i])
	}
}

func (self *BaseData) Run(inst Bpoint, pos Value) {

	self.withZ(inst, pos)

	self.calc(pos + 1)

}

func (self *BaseData) withZ(inst Bpoint, pos Value) {
	self.Inst[pos] = inst
	self.Nbp[pos] = getNextBp(self.Bp[pos], inst)

	for i := 0; i < ROWS; i++ {
		withZ(&self.Data[i][pos], inst)
	}
	withZ(&self.Xg[pos], inst)
	withZ(&self.Gz[pos], inst)
	withZ(&self.Gf[pos], inst)
	withZ(&self.Gf1[pos], inst)
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
			env.Bp = self.Bp[pos]
		} else {
			env.Bp = pointToBp(self.Data[i-1][pos])
		}

		env.Last = self.Data[i][pos-1]
		tmp = calcReverse(&env)
		self.Data[i][pos] = tmp
		self.Xg[pos].V += tmp.V
	}

	self.calcGzf(pos, 1)
	self.calcGzf(pos, -1)
	self.calcGf1(pos)
}

func (self *BaseData) calcGzf(pos Value, same Value) {
	var env = Env{}
	var val *[COLS]Point
	if same == 1 {
		val = &self.Gz
	} else {
		val = &self.Gf
	}
	env.CurrentCol = pos
	env.Bp = pointToBp(self.Xg[pos])
	env.Last = val[pos-1]

	data := calcFollow(&env, same)
	val[pos] = data
}

func (self *BaseData) calcGf1(pos Value) {
	if self.Gf[pos].V > 0 {
		self.Gf1[pos].V = 1
	} else {
		self.Gf1[pos].V = -1
	}
}

func (self *BaseData) showValue(start int) (arr []string) {
	length := 12
	end := start + length
	template := "%-10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s\n"

	title := make([]Bpoint, 0)

	for i := start; i < end; i++ {
		title = append(title, Bpoint(i))
	}

	printBp := func(title string, val []Bpoint) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}
	printPoint := func(title string, val []Point) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}

	arr = append(arr, printBp("Col", title))
	arr = append(arr, fmt.Sprintln(strings.Repeat("-", 145)))
	arr = append(arr, printBp("Inst", self.Inst[start:end]))
	arr = append(arr, printBp("BP", self.Bp[start:end]))
	arr = append(arr, printBp("New BP", self.Nbp[start:end]))

	for i := 0; i < ROWS; i++ {
		arr = append(arr, printPoint(string(i), self.Data[i][start:end]))
	}

	arr = append(arr, printPoint("XG", self.Xg[start:end]))
	arr = append(arr, printPoint("GZ", self.Gz[start:end]))
	arr = append(arr, printPoint("GF", self.Gf[start:end]))
	arr = append(arr, printPoint("GF1", self.Gf1[start:end]))
	return arr
}

func (self BaseData) String() string {

	/*
		print := func(title string, val Duck) string {
			list := reflect.ValueOf(val)
			v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 := list.Index(0), list.Index(1),
				list.Index(3), list.Index(4), list.Index(5), list.Index(6),
				list.Index(7), list.Index(8), list.Index(9), list.Index(10),
				list.Index(11)
			ret := fmt.Sprintf(template, title, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11)
		}
	*/
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	arr = self.showValue(35)
	ret += strings.Join(arr, "")
	return ret

}
