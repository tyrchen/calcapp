package calc

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
