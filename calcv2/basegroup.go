package calcv2

import (
	. "calcapp/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	BP_FILE = fmt.Sprintf("data/calcv2_%d.bp", GROUP_SIZE)
)

func (self *BaseGroup) Init() {
	// init the col 0
	self.LoadBp(BP_FILE)

	data := make([]Point, GROUP_SIZE)

	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].Init()
		data[i] = self.Data[i].G1[0]
	}

	ret := CalcReduce(data)

	for i := 0; i < len(ret); i++ {
		self.G1[i][0].V = ret[i].V
	}
}

func (self *BaseGroup) LoadBp(filename string) {
	bytes, _ := ioutil.ReadFile(filename)
	content := string(bytes)
	lines := strings.Split(content, "\n")

	for i := 0; i < GROUP_SIZE; i++ {
		line := strings.TrimSpace(lines[i])
		self.Data[i].LoadBp(strings.Repeat(line, (COLS-1)/len(line)))
	}
}

func (self *BaseGroup) Run(inst Bpoint, pos Value) {

	self.withZ(inst, pos)

	self.calc(pos + 1)

}

func (self *BaseGroup) withZ(inst Bpoint, pos Value) {
	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].withZ(inst, pos)
	}

	for i := 0; i < CHUNK_SIZE+1; i++ {
		WithZ(&self.G1[i][pos], inst)
	}
}

func (self *BaseGroup) calc(pos Value) {
	data := make([]Point, GROUP_SIZE)
	for i := 0; i < GROUP_SIZE; i++ {
		self.Data[i].calc(pos)
		data[i] = self.Data[i].G1[pos]
	}

	ret := CalcReduce(data)

	for i := 0; i < len(ret); i++ {
		self.G1[i][pos].V = ret[i].V
	}
}
