package calc

// 反1为基本点
// 1. 与基本点符号相同
// 2. 与基本点符号相反
// 3. 上一个基本点带z与上相反，不带z与上相同。
// 9组的结果

const (
	TS_STOP_VALUE = 511
)

func (self *ThreeSome) Init(up Point) {
	zero := Point{false, 0}
	self.Up[0] = up
	self.V1[0] = ts_calcFollow1(zero, self.Up[0])
	self.V2[0] = ts_calcFollow2(zero, self.Up[0])
	self.V3[0] = ts_calcFollow3(zero, self.Up[0], zero)
	self.Sum[0].V = self.V1[0].V + self.V2[0].V + self.V3[0].V
}

func (self *ThreeSome) calc(pos Value) {
	self.V1[pos] = ts_calcFollow1(self.V1[pos-1], self.Up[pos])
	self.V2[pos] = ts_calcFollow2(self.V2[pos-1], self.Up[pos])
	self.V3[pos] = ts_calcFollow3(self.V3[pos-1], self.Up[pos], self.Up[pos-1])
	self.Sum[pos].V = self.V1[pos].V + self.V2[pos].V + self.V3[pos].V
}

func (self *ThreeSome) withZ(inst Bpoint, pos Value) {
	withZ(&self.Up[pos], inst)
	withZ(&self.V1[pos], inst)
	withZ(&self.V2[pos], inst)
	withZ(&self.V3[pos], inst)
	withZ(&self.Sum[pos], inst)
}

func (self *ThreeSome) Run(inst Bpoint, pos Value) {
	self.withZ(inst, pos)
	self.calc(pos + 1)
}

func ts_calc137(last Point) (ret Point) {
	ret.T = false
	if !last.T {
		ret.V = 1
	} else {
		ret.V = Abs(last.V)*2 + 1

		if ret.V > TS_STOP_VALUE {
			ret.V = 1
		}
	}
	return ret
}

// same sign with basepoint
func ts_calcFollow1(last, up Point) (ret Point) {
	sign := sign(up)
	ret = ts_calc137(last)
	ret.V *= sign
	return ret
}

// same sign with basepoint
func ts_calcFollow2(last, up Point) (ret Point) {
	sign := sign(up)
	ret = ts_calc137(last)
	ret.V *= sign * -1
	return ret
}

// same sign with basepoint
func ts_calcFollow3(last, up, last_up Point) (ret Point) {
	sign := sign(up)
	sign1 := zsign(last_up)
	ret = ts_calc137(last)
	ret.V *= sign * (sign1 * -1)
	return ret
}
