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
	BP_START       = 0x222233334444
	BP_END         = 0xf08eafddccccbbbb
	BP_SLICE_START = 2
	BP_SLICE_END   = 58
	BP_GAP         = 0xe1111189321
	BP_TOTAL       = 37000
	BP_FILENAME    = "/var/tmp/calcapp/bp/basepoint%02d.dat"
)

func ValueToBp(val uint64) []uint8 {
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

func ValueToString(val uint64) string {
	arr := make([]string, 0)
	for val > 0 {
		arr = append(arr, string(val%2))
		val /= 2
	}
	return strings.Join(arr[BP_SLICE_START:BP_SLICE_END], "")
}

func GenerateBpFiles(max uint) {
	values_list := make([][]uint64, 0)
	var i uint
	var val uint64

	i = 0
	values := make([]uint64, 0)
	for val = BP_START; val < BP_END; val += BP_GAP {

		values = append(values, val)
		if len(values) >= BP_TOTAL {
			values_list = append(values_list, values)
			fmt.Printf("i=%d, values[0]=%d\n", i, values[0])
			i++
			if i >= max {
				break
			}
			values = make([]uint64, 0)
		}
	}

	for i = 0; i < max; i++ {
		m := new(bytes.Buffer)
		enc := gob.NewEncoder(m)
		fmt.Printf("Length of values_list[%d] is %d", i, len(values_list[i]))
		enc.Encode(values_list[i])

		filename := fmt.Sprintf(BP_FILENAME, i+1)
		err := ioutil.WriteFile(filename, m.Bytes(), 0600)
		if err != nil {
			panic(err)
		}

		fmt.Printf("File %s saved\n", filename)
	}
}

func LoadBpFile(index uint) (values []uint64) {
	name := fmt.Sprintf(BP_FILENAME, index)
	n, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	m := bytes.NewBuffer(n)
	dec := gob.NewDecoder(m)
	values = make([]uint64, 0)
	err = dec.Decode(&values)
	if err != nil {
		panic(err)
	}
	return
}
