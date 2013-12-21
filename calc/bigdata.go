package calc

import ()

type chanBigData struct {
	zg   Value
	gz   Value
	gzmm Value
	gf   Value
	gfmm Value
	gf1  Value
}

// load bp for all the groupData
func (self *BigData) LoadBp() {
	var i uint
	for i = 0; i < ZG_NUM; i++ {
		self.Data[i].LoadBp(i)
	}
}

func (self *BigData) LoadSelfBp() {
	var i uint
	for i = 0; i < ZG_NUM; i++ {
		self.Data[i].LoadSelfBp(i)
	}
}

func (self *BigData) LoadMp() {
	var i uint
	for i = 0; i < ZG_NUM; i++ {
		self.Data[i].LoadMp(i)
	}
}

func (self *BigData) Clear() {
	*self = *new(BigData)
}

func (self *BigData) calcDgValues(pos Value) {
	for i := 0; i < ZG_NUM; i++ {
		bp := self.Data[i].Bp[pos]
		self.Dg[pos].V += self.Data[i].Zg[pos].V
		self.Gz[pos].V += signFollow(self.Data[i].Gz[pos].V, bp)
		self.Gzmm[pos].V += signFollow(self.Data[i].Gzmm[pos].V, bp)
		self.Gf[pos].V += signFollow(self.Data[i].Gf[pos].V, bp)
		self.Gfmm[pos].V += signFollow(self.Data[i].Gfmm[pos].V, bp)
		self.Gf1[pos].V += signFollow(self.Data[i].Gf1[pos].V, bp)
	}
}

func (self *BigData) Init() {
	var i uint
	for i = 0; i < ZG_NUM; i++ {
		self.Data[i].Init(i)
	}
	self.calcDgValues(0)
}

func (self *BigData) Run(inst Bpoint, pos Value) {
	var j Value
	next_pos := pos + 1
	chn := make(chan chanBigData, ZG_NUM)

	self.Inst[pos] = inst

	worker := func(i Value, chn chan chanBigData) {
		self.Data[i].Run(inst, pos)
		bp := self.Data[i].Bp[next_pos]
		chn <- chanBigData{
			self.Data[i].Zg[next_pos].V,
			signFollow(self.Data[i].Gz[next_pos].V, bp),
			signFollow(self.Data[i].Gzmm[next_pos].V, bp),
			signFollow(self.Data[i].Gf[next_pos].V, bp),
			signFollow(self.Data[i].Gfmm[next_pos].V, bp),
			signFollow(self.Data[i].Gf1[next_pos].V, bp),
		}

	}
	for j = 0; j < ZG_NUM; j++ {
		go worker(j, chn)
	}

	received := 0
	for val := range chn {
		received++

		self.Dg[next_pos].V += val.zg
		self.Gz[next_pos].V += val.gz
		self.Gzmm[next_pos].V += val.gzmm
		self.Gf[next_pos].V += val.gf
		self.Gfmm[next_pos].V += val.gfmm
		self.Gf1[next_pos].V += val.gf1
		if received == ZG_NUM {
			close(chn)
		}
	}

	self.withZ(inst, pos)
}

func (self *BigData) withZ(inst Bpoint, pos Value) {
	withZ(&self.Dg[pos], inst)
	withZ(&self.Gz[pos], inst)
	withZ(&self.Gzmm[pos], inst)
	withZ(&self.Gf[pos], inst)
	withZ(&self.Gfmm[pos], inst)
	withZ(&self.Gf1[pos], inst)
}
