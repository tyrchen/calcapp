package calc

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
	Calc(inst Bpoint, pos Value)
}

func (self *BaseData) Calc(inst Bpoint, pos Value) {
	self.Inst[pos] = inst
	self.Nbp[pos] = getNextBp(self.Bp[pos], inst)
	var env = Env{}
	env.init()
	for i := 0; i < ROWS; i++ {
		env.Bp = self.Bp[pos]
		env.CurrentCol = pos
		env.Last = self.Data[i][pos-1]
		self.Data[i][pos] = calcReverse(&env)
	}
}
