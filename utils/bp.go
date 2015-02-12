package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	BP_START       = 0x11a7222233334444
	BP_END         = 0xf08eafddccccbbbb
	BP_SLICE_START = 2
	BP_SLICE_END   = 58
	BP_GAP         = 0xe11893219c86
	BP_TOTAL       = 81 * 81 //37000
	BP_ZG          = 60
	BP_COLS        = 55 + 1
	BP_FILENAME    = "/var/data/calcapp/bp/%s/basepoint%02d.dat"
	BP_ZG_FILENAME = "/var/data/calcapp/bpzg.txt"
)

type Bp [BP_COLS]uint8
type BpData [BP_TOTAL]Bp
type BpZgData [BP_ZG]Bp

func uint64ToBp(val uint64) []uint8 {
	arr := make([]uint8, 0)
	for val > 0 {
		arr = append(arr, uint8(val%2))
		val /= 2
	}
	return arr[BP_SLICE_START:BP_SLICE_END]
}

func StringToBp(val string) []uint8 {
	arr := make([]uint8, 0)
	for i := 0; i < len(val); i++ {
		v, _ := strconv.Atoi(string(val[i]))
		arr = append(arr, uint8(v))
	}
	return arr
}

func generateBpData(val uint64) (ret Bp) {
	data := uint64ToBp(val)

	for i := 0; i < BP_COLS; i++ {
		ret[i] = data[i]
	}
	return ret
}

func GenerateBpFiles(max uint) {
	var zero BpData
	var val uint64
	var i, j uint = 0, 0

	values := zero
	values_list := make([]BpData, 0)

	for val = BP_START; val < BP_END; val += BP_GAP {

		if j < BP_TOTAL {
			values[j] = generateBpData(val)
			j++
		} else {
			values_list = append(values_list, values)
			i++
			if i >= max {
				break
			}
			values = zero
			j = 0
		}
	}

	for i = 0; i < max; i++ {
		filename := GetFileName(i, true)
		saveBpFile(filename, values_list[i])
		filename1 := GetFileName(i+30, true)
		saveBpFile(filename1, values_list[i])
	}
}

func GenerateBp2File(filename string) {
	var value BpData
	var val uint64
	var i uint

	for val = BP_START; val < BP_END; val += BP_GAP {
		value[i] = generateBpData(val)
		i++

		if (i >= BP_TOTAL) {
			break
		}
	}
	saveBpFile(filename, value)
}

func LoadBp2File(filename string) (values BpData) {
	n, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	m := bytes.NewBuffer(n)
	dec := gob.NewDecoder(m)
	err = dec.Decode(&values)
	if err != nil {
		panic(err)
	}
	return
}

func saveBpFile(filename string, values BpData) {
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	enc.Encode(values)

	err := ioutil.WriteFile(filename, m.Bytes(), 0600)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File %s saved\n", filename)
}

func SaveBpFile(index uint, values BpData) {
	filename := GetFileName(index, false)
	saveBpFile(filename, values)
}

func GetFileName(index uint, origin bool) string {
	options := map[bool]string{
		false: "new",
		true:  "origin",
	}

	return fmt.Sprintf(BP_FILENAME, options[origin], index+1)
}

func LoadBpFile(index uint, origin bool) (values BpData) {
	name := GetFileName(index, origin)

	fmt.Println("Load basepoint from: ", name)

	n, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	m := bytes.NewBuffer(n)
	dec := gob.NewDecoder(m)
	err = dec.Decode(&values)
	if err != nil {
		panic(err)
	}
	return
}

func LoadZgBp() (values BpZgData) {
	bytes, _ := ioutil.ReadFile(BP_ZG_FILENAME)
	content := string(bytes)
	lines := strings.Split(content, "\n")

	for i := 0; i < BP_ZG; i++ {
		line := strings.TrimSpace(lines[i])
		for j := 0; j < BP_COLS; j++ {
			tmp, _ := strconv.Atoi(string(line[j]))
			values[i][j] = uint8(tmp)
		}
	}
	return values
}

func GetZgBp(index uint) (value Bp) {
	return LoadZgBp()[index]
}
