package calcv5

import (
	. "calcapp/utils"
	//"fmt"
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
		for i := index * sz; i < (index + 1) * sz; i++ {
			self.Data[i].withZ(inst, pos)
		}
		ch <- true
	}

	for g := 0; g < CONCURRENCY; g++ {
		go worker(g)
	}
	
	WithZ(&self.G137[pos], inst)
	
	for g := 0; g < CONCURRENCY; g++ {
		<- ch
	}
}


func (self *BaseGroup) calc(pos Value) {
	sz := GROUP_SIZE / CONCURRENCY
	ch := make(chan bool, CONCURRENCY)

	worker := func(index int) {
		for i := index * sz; i < (index + 1) * sz; i++ {
			self.Data[i].calc(pos)
			self.G137[pos].V += self.Data[i].G137[pos].V
		}
		ch <- true
	}

	for g := 0; g < CONCURRENCY; g++ {
		go worker(g)
	}

	for g := 0; g < CONCURRENCY; g++ {
		<- ch
	}

}
