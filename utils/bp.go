package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	BP_START       = 0x001111222233334444
	BP_END         = 0xeeeeddddccccbbbb
	BP_SLICE_START = 5
	BP_SLICE_END   = 61
	BP_GAP         = 87654321
	BP_TOTAL       = 37000
)

func ValueToString(val uint64) string {
	var arr []string
	for val > 0 {
		arr = append(arr, string(val%2))
		val /= 2
	}
	return strings.Join(arr, "")[BP_SLICE_START:BP_SLICE_END]
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

		filename := fmt.Sprintf("./basepoint%02d.dat", i+1)
		err := ioutil.WriteFile(filename, m.Bytes(), 0600)
		if err != nil {
			panic(err)
		}

		fmt.Printf("File %s saved\n", filename)
	}
}

func LoadBpFile(name string) (values []uint64) {
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
