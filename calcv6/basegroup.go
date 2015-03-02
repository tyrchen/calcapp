package calcv6

import (
	. "calcapp/utils"
	// "fmt"
	//"io/ioutil"
	//"strings"
	//"strconv"
)

var (
	//BP_FILE = fmt.Sprintf("data/calcv2_%d.bp", GROUP_SIZE)
	BP_FILE = "data/calcv4.bp"
)

func (self *BaseGroup) Init() {
	// init the col 0
	self.LoadBp(BP_FILE)

	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].Init()
		self.G137[0].V += self.Data[i].G137[0].V
	}
}

func (self *BaseGroup) LoadBp(filename string) {
	values := LoadBp2File(filename)

	for row, value := range values {
		if row >= GROUP_SIZE {
			break
		}
		self.Data[row].LoadBp(value[:])
	}
}

func (self *BaseGroup) Run(inst Bpoint, pos Value) {

	self.withZ(inst, pos)

	self.calc(pos + 1)

}

func (self *BaseGroup) withZ(inst Bpoint, pos Value) {
	sz := GROUP_SIZE / CONCURRENCY

	ch := make(chan bool, CONCURRENCY)

	worker := func(index int) {
		for i := index * sz; i < (index+1)*sz; i++ {
			self.Data[i].withZ(inst, pos)
		}
		ch <- true
	}

	for g := 0; g < CONCURRENCY; g++ {
		go worker(g)
	}

	for g := 0; g < CONCURRENCY; g++ {
		<-ch
	}

	WithZ(&self.G137[pos], inst)
	WithZ(&self.Gz[pos], inst)
	WithZ(&self.Gf[pos], inst)

	for i := 0; i < 6; i++ {
		WithZ(&self.Gzr[i][pos], inst)
	}

	for i := 0; i < 6; i++ {
		WithZ(&self.Gfr[i][pos], inst)
	}

	WithZ(&self.Zg2[pos], inst)

}

func (self *BaseGroup) calc(pos Value) {
	sz := GROUP_SIZE / CONCURRENCY
	ch := make(chan bool, CONCURRENCY)

	values := make([]Value, sz)

	worker := func(index int) {
		for i := index * sz; i < (index+1)*sz; i++ {
			self.Data[i].calc(pos)
			values[index] += self.Data[i].G137[pos].V
		}
		ch <- true
	}

	for g := 0; g < CONCURRENCY; g++ {
		go worker(g)
	}

	for g := 0; g < CONCURRENCY; g++ {
		<-ch
		self.G137[pos].V += values[g]
	}

	self.Gz[pos].V = calcGz(self.G137[pos].V)
	self.Gf[pos].V = -calcGz(self.G137[pos].V)

	for i := 0; i < 6; i++ {
		self.Gzr[i][pos].V = calcRev(self.Gzr[i][pos-1], self.Gz[pos], i)
		self.Gfr[i][pos].V = calcRev(self.Gfr[i][pos-1], self.Gf[pos], i)
		self.Zg2[pos].V += (self.Gzr[i][pos].V + self.Gfr[i][pos].V)
	}

}

func calcRev(last, up Point, rev int) (ret Value) {
	sign := Sign(up)

	if last.T {
		ret = 1 * sign
	} else {
		ret = Abs(last.V)*2 + 1
		if ret > 63 {
			ret = 1
		}
		ret *= sign
	}

	tmp := Abs(ret)
	if rev == 1 { // 五翻一
		if tmp == 63 {
			ret = -ret
		}
	} else if rev == 2 { // 四翻二
		if tmp == 31 {
			ret = -ret
		}
	} else if rev == 3 { // 三翻三
		if tmp == 15 || tmp == 31 || tmp == 63 {
			ret = -ret
		}
	} else if rev == 4 { // 二翻二
		if tmp == 7 || tmp == 15 {
			ret = -ret
		}
	} else if rev == 5 { // 一翻一
		if tmp == 3 || tmp == 15 || tmp == 63 {
			ret = -ret
		}
	}
	return
}
