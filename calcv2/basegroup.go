package calcv2

import (
	. "calcapp/utils"
	// "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

var (
	//BP_FILE = fmt.Sprintf("data/calcv2_%d.bp", GROUP_SIZE)
	BP_FILE = "data/calcv2.bp"
)

func (self *BaseGroup) Init() {
	// init the col 0
	self.LoadBpNew(BP_FILE)

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

type Bpf [8]Bpoint
type BpfData [LAYER][3]Bpf

func getBp(bpData *BpfData, index int) (bp Bpf) {
	arr_xor := func(b1 Bpf, b2 Bpf) (ret Bpf) {
		for i := 0; i < len(b1); i++ {
			ret[i] = b1[i] ^ b2[i]
		}
		return
	}

	l := 1
	for i:= 0; i < LAYER; i++ {
		var p Bpf
		if (index + 1) > 3 * l {
			p = bpData[i][2]
		} else if (index + 1) > 2 * l {
			p = bpData[i][1]
		} else {
			p = bpData[i][0]
		}
		bp = arr_xor(bp, p)

		l *= 3
	}
	return
}

func (self *BaseGroup) LoadBpNew(filename string) {
	var bpData BpfData
	bytes, _ := ioutil.ReadFile(filename)
	content := string(bytes)
	lines := strings.Split(content, "\n")

	for i := 0; i < LAYER; i++ {
		for j := 0; j < 3; j++ {
			line := strings.TrimSpace(lines[i*3+j])
			for k := 0; k < 8; k++ {
				tmp, _ := strconv.Atoi(string(line[k]))
				bpData[i][j][k] = Bpoint(tmp)
			}

		}
	}

	for i := 0; i < GROUP_SIZE; i++ {
		bp := getBp(&bpData, i)
		self.Data[i].LoadBpArray(bp)
	}
	return
}

func (self *BaseGroup) LoadBp(filename string) {
	bytes, _ := ioutil.ReadFile(filename)
	content := string(bytes)
	lines := strings.Split(content, "\n")

	sz := GROUP_SIZE / 3

	ch := make(chan bool, 3)

	worker := func(index int) {
		for i := index * sz; i < (index + 1) * sz; i++ {
			line := strings.TrimSpace(lines[index])
			self.Data[i].LoadBp(strings.Repeat(line, (COLS-1)/len(line)))
		}
		ch <- true
	}

	for g := 0; g < 3; g++ {
		go worker(g)	
	}

	for g := 0; g < 3; g++ {
		<- ch
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

	for i := 0; i < CHUNK_SIZE+1; i++ {
		WithZ(&self.G1[i][pos], inst)
	}

	for g := 0; g < CONCURRENCY; g++ {
		<- ch
	}
}

func (self *BaseGroup) calc(pos Value) {
	data := make([]Point, GROUP_SIZE)
	sz := GROUP_SIZE / CONCURRENCY
	ch := make(chan bool, CONCURRENCY)

	worker := func(index int) {
		for i := index * sz; i < (index + 1) * sz; i++ {
			self.Data[i].calc(pos)
			data[i] = self.Data[i].G1[pos]
		}
		ch <- true
	}

	for g := 0; g < CONCURRENCY; g++ {
		go worker(g)
	}

	for g := 0; g < CONCURRENCY; g++ {
		<- ch
	}

	ret := CalcReduce(data)

	for i := 0; i < len(ret); i++ {
		self.G1[i][pos].V = ret[i].V
	}
}
