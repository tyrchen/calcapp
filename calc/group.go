package calc

import (
	"calcapp/utils"
	"fmt"
	"strings"
)

type GroupData struct {
	Inst [COLS]Bpoint
	Data [GROUP_SIZE]BaseData
	Zg   [COLS]Point
	Gz   [COLS]Point
	Gf1  [COLS]Point
}

type chanData struct {
	zg  Value
	gz  Value
	gf1 Value
}

func (self *GroupData) LoadBp(index uint) {
	values := utils.LoadBpFile(index)
	for row, value := range values {
		data := utils.ValueToBp(value)
		self.Data[row].LoadBp(data)
	}
}

func (self *GroupData) Clear() {
	*self = GroupData{}
}

func (self *GroupData) Init() {

	// init the col 0
	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].Init()
	}
	self.Zg[0].V = 1
	self.Gz[0].V = 1
	self.Gf1[0].V = -1
}

func (self *GroupData) Run(inst Bpoint, pos Value) {
	self.Inst[pos] = inst
	self.withZ(inst, pos)
	self.calc(pos + 1)
}

func (self *GroupData) withZ(inst Bpoint, pos Value) {
	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].withZ(inst, pos)
	}
	withZ(&self.Zg[pos], inst)
	withZ(&self.Gz[pos], inst)
	withZ(&self.Gf1[pos], inst)
}

func (self *GroupData) calc(pos Value) {
	if pos >= COLS {
		return
	}
	chn := make(chan chanData, CHUNKS)

	worker := func(start, end int, chn chan chanData) {
		var data chanData
		for i := start; i < end; i++ {
			self.Data[i].calc(pos)
			data.zg += self.Data[i].Xg[pos].V
			data.gz += self.Data[i].Gz[pos].V
			data.gf1 += self.Data[i].Gf1[pos].V
		}
		chn <- data
	}

	for i := 0; i < GROUP_SIZE; i += CHUNK_SIZE {
		go worker(i, i+CHUNK_SIZE, chn)
	}

	received := 0
	for val := range chn {
		received++

		self.Zg[pos].V += val.zg
		self.Gz[pos].V += val.gz
		self.Gf1[pos].V += val.gf1
		if received == CHUNKS {
			close(chn)
		}
	}

	/*
		for i := 0; i < GROUP_SIZE; i++ {
			self.Data[i].calc(pos)
			self.Zg[pos].V += self.Data[i].Xg[pos].V
			self.Gz[pos].V += self.Data[i].Gz[pos].V
			self.Gf1[pos].V += self.Data[i].Gf1[pos].V
		}
	*/

}

func (self *GroupData) showValue(start int) (arr []string) {
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

	arr = append(arr, printPoint("ZG", self.Zg[start:end]))
	arr = append(arr, printPoint("GZ", self.Gz[start:end]))
	arr = append(arr, printPoint("GF1", self.Gf1[start:end]))

	return arr
}

func (self GroupData) String() string {
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	arr = self.showValue(35)
	ret += strings.Join(arr, "")
	return ret
}
