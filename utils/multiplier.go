package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	osutil "github.com/tyrchen/goutil/osutil"
	"io/ioutil"
)

const (
	MP_FILENAME = "/var/data/calcapp/mp/mp%02d.dat"
)

type Multiplier struct {
	Gzm int
	Gfm int
}

func saveMpFile(filename string, values Multiplier) {
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	enc.Encode(values)

	err := ioutil.WriteFile(filename, m.Bytes(), 0600)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File %s saved, values: %d, %d\n", filename, values.Gzm, values.Gfm)
}

func SaveMpFile(index uint, values Multiplier) {
	filename := getMpFileName(index)
	saveMpFile(filename, values)
}

func getMpFileName(index uint) string {
	return fmt.Sprintf(MP_FILENAME, index+1)
}

func LoadMpFile(index uint) (values Multiplier) {
	name := getMpFileName(index)

	if !osutil.FileExists(name) {
		values.Gfm = 1
		values.Gzm = 1
		fmt.Printf("Load multiplier not from file, values are %d, %d\n", values.Gzm, values.Gfm)
		return
	}

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

	fmt.Printf("Load multiplier from: %s, values are %d, %d\n", name, values.Gzm, values.Gfm)

	return
}
