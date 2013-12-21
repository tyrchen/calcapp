package calc

import (
	"calcapp/utils"
	osutil "github.com/tyrchen/goutil/osutil"
)

type chanData struct {
	zg  Value
	gz  Value
	gf  Value
	gf1 Value
}

// load bp for all the groups
func (self *GroupData) LoadBp(index uint) {
	var values utils.BpData

	self.Index = index

	if osutil.FileExists(utils.GetFileName(index, false)) {
		values = utils.LoadBpFile(index, false)
	} else {
		values = utils.LoadBpFile(index, true)
	}

	for row, value := range values {
		self.Data[row].LoadBp(value[:])
	}
}

func (self *GroupData) LoadSelfBp(index uint) {
	bp := utils.GetZgBp(index)
	for i := 0; i < COLS; i++ {
		self.Bp[i] = Bpoint(bp[i])
	}
}

func (self *GroupData) LoadMp(index uint) {
	values := utils.LoadMpFile(index)
	self.Gzm = Value(values.Gzm)
	self.Gfm = Value(values.Gfm)
}

func (self *GroupData) Clear() {
	*self = *new(GroupData)
}

func (self *GroupData) Init(index uint) {
	self.LoadBp(index)
	self.LoadSelfBp(index)
	self.LoadMp(index)
	// init the col 0
	for i := 0; i < GROUP_SIZE-1; i++ {
		self.Data[i].Init()
		self.Zg[0].V += self.Data[i].Xg[0].V
		self.Gz[0].V += self.Data[i].Gz[0].V
		self.Gf[0].V += self.Data[i].Gf[0].V
		self.Gf1[0].V += self.Data[i].Gf1[0].V
	}
	self.Gzmm[0].V = self.Gz[0].V * self.Gzm
	self.Gfmm[0].V = self.Gf[0].V * self.Gfm
}

func (self *GroupData) Run(inst Bpoint, pos Value) {
	new_inst := getNextBp(self.Bp[pos], inst)
	self.Inst[pos] = new_inst
	self.withZ(new_inst, pos)
	self.calc(pos + 1)

	if pos == COLS-2 {
		self.SaveNewBp()
		self.CalcMultiply()
	}
}

func (self *GroupData) withZ(inst Bpoint, pos Value) {
	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].withZ(inst, pos)
	}
	withZ(&self.Zg[pos], inst)
	withZ(&self.Gz[pos], inst)
	withZ(&self.Gzmm[pos], inst)
	withZ(&self.Gf[pos], inst)
	withZ(&self.Gfmm[pos], inst)
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
			data.gf += self.Data[i].Gf[pos].V
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
		self.Gf[pos].V += val.gf
		self.Gf1[pos].V += val.gf1
		if received == CHUNKS {
			close(chn)
		}
	}

	// remove last value, so that we actually calculate GROUP_SIZE - 1
	self.Zg[pos].V -= self.Data[GROUP_SIZE-1].Xg[pos].V
	self.Gz[pos].V -= self.Data[GROUP_SIZE-1].Gz[pos].V
	self.Gf[pos].V -= self.Data[GROUP_SIZE-1].Gf[pos].V
	self.Gf1[pos].V -= self.Data[GROUP_SIZE-1].Gf1[pos].V

	self.Gzmm[pos].V = self.Gzm * self.Gz[pos].V
	self.Gfmm[pos].V = self.Gfm * self.Gf[pos].V

	/*
		for i := 0; i < GROUP_SIZE; i++ {
			self.Data[i].calc(pos)
			self.Zg[pos].V += self.Data[i].Xg[pos].V
			self.Gz[pos].V += self.Data[i].Gz[pos].V
			self.Gf1[pos].V += self.Data[i].Gf1[pos].V
		}
	*/

}

func (self *GroupData) SaveNewBp() {
	var values utils.BpData

	for i := 0; i < GROUP_SIZE; i++ {
		for j := 0; j < COLS; j++ {
			values[i][j] = uint8(self.Data[i].Nbp[j])
		}
	}
	utils.SaveBpFile(self.Index, values)
}

func (self *GroupData) CalcMultiply() {
	// calc gfm
	self.Gzm = calc137forGzf(self.Gzm, calcMultiply(self.Gz))
	self.Gfm = calc137forGzf(self.Gfm, calcMultiply(self.Gf))
	utils.SaveMpFile(self.Index, utils.Multiplier{int(self.Gzm), int(self.Gfm)})
}
