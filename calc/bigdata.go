package calc

import ()

const (
	THREESOME_BIG_NUM       = 61
	THREESOME_DIFFERENTIATE = 1000
)

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

func (self *BigData) calcThreeSomeSum(pos, cursor Value) (ret Value) {
	for i := 0; i < THREESOME_NUM; i++ {
		v := self.TsData[cursor][i].Sum[pos].V
		if v > THREESOME_BIG_NUM {
			ret += THREESOME_DIFFERENTIATE
		} else if v < -THREESOME_BIG_NUM {
			ret += -THREESOME_DIFFERENTIATE
		} else {
			ret += v
		}

	}
	return
}

func (self *BigData) calcTsRet(pos Value) (ret Value) {
	for i := 0; i < THREESOME_TOTAL; i++ {
		ret += self.TsValue[i][pos].V
	}
	return
}

func (self *BigData) CalcDelta() (ret [5]Value) {
	ret[0] = calcDelta(self.Gz)
	ret[1] = calcDelta(self.Gzmm)
	ret[2] = calcDelta(self.Gf)
	ret[3] = calcDelta(self.Gfmm)
	ret[4] = calcDelta(self.Gf1)
	return ret
}

func (self *BigData) Init() {
	var i, j uint
	var k Value
	for i = 0; i < ZG_NUM; i++ {
		self.Data[i].Init(i)
	}
	self.calcDgValues(0)
	for i = 0; i < THREESOME_NUM; i++ {
		if i == 0 {
			self.TsData[0][i].Init(self.Gzmm[0])
			self.TsData[1][i].Init(self.Gfmm[0])
			self.TsData[2][i].Init(self.Gf1[0])
		} else {
			for j = 0; j < THREESOME_TOTAL; j++ {
				self.TsData[j][i].Init(self.TsData[j][i-1].Up[0])
			}
		}
	}
	for k = 0; k < THREESOME_TOTAL; k++ {
		self.TsValue[k][0].V = self.calcThreeSomeSum(0, k)
	}
	self.TsRet[0].V = self.calcTsRet(0)
}

func (self *BigData) Run(inst Bpoint, pos Value) {
	self.withZ(inst, pos)
	self.calc(pos + 1)
}

func (self *BigData) calc(pos Value) {
	var j, k Value
	chn := make(chan chanBigData, ZG_NUM)

	worker := func(i Value, chn chan chanBigData) {
		self.Data[i].calc(pos)
		bp := self.Data[i].Bp[pos]
		chn <- chanBigData{
			self.Data[i].Zg[pos].V,
			signFollow(self.Data[i].Gz[pos].V, bp),
			signFollow(self.Data[i].Gzmm[pos].V, bp),
			signFollow(self.Data[i].Gf[pos].V, bp),
			signFollow(self.Data[i].Gfmm[pos].V, bp),
			signFollow(self.Data[i].Gf1[pos].V, bp),
		}

	}
	for j = 0; j < ZG_NUM; j++ {
		go worker(j, chn)
	}

	received := 0
	for val := range chn {
		received++

		self.Dg[pos].V += val.zg
		self.Gz[pos].V += val.gz
		self.Gzmm[pos].V += val.gzmm
		self.Gf[pos].V += val.gf
		self.Gfmm[pos].V += val.gfmm
		self.Gf1[pos].V += val.gf1
		if received == ZG_NUM {
			close(chn)
		}
	}

	for j = 0; j < THREESOME_NUM; j++ {
		if j == 0 {
			self.TsData[0][j].Up[pos] = self.Gzmm[pos]
			self.TsData[1][j].Up[pos] = self.Gfmm[pos]
			self.TsData[2][j].Up[pos] = self.Gf1[pos]
		} else {
			for k = 0; k < THREESOME_TOTAL; k++ {
				self.TsData[k][j].Up[pos] = self.TsData[k][j-1].Sum[pos]
			}
		}
		for k = 0; k < THREESOME_TOTAL; k++ {
			self.TsData[k][j].calc(pos)
		}
	}

	for k = 0; k < THREESOME_TOTAL; k++ {
		self.TsValue[k][pos].V = self.calcThreeSomeSum(pos, k)
	}
	self.TsRet[pos].V = self.calcTsRet(pos)
}

func (self *BigData) withZ(inst Bpoint, pos Value) {
	var i, j Value
	self.Inst[pos] = inst
	for i = 0; i < ZG_NUM; i++ {
		self.Data[i].withZ(inst, pos)
	}
	withZ(&self.Dg[pos], inst)
	withZ(&self.Gz[pos], inst)
	withZ(&self.Gzmm[pos], inst)
	withZ(&self.Gf[pos], inst)
	withZ(&self.Gfmm[pos], inst)
	withZ(&self.Gf1[pos], inst)

	for i = 0; i < THREESOME_NUM; i++ {
		for j = 0; j < THREESOME_TOTAL; j++ {
			self.TsData[j][i].withZ(inst, pos)
		}
	}

	for j = 0; j < THREESOME_TOTAL; j++ {
		withZ(&self.TsValue[j][pos], inst)
	}
	withZ(&self.TsRet[pos], inst)
}
