package calc

import (
	"calcapp/utils"
)

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