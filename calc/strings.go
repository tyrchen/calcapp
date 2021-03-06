package calc

import (
	"fmt"
	"strconv"
	"strings"
)

func (self Bpoint) String() string {
	return strconv.Itoa(int(self))
}

func (self Point) String() string {
	if self.T {
		return fmt.Sprintf("z%d", self.V)
	} else {
		return fmt.Sprintf("%d", self.V)
	}
}

// for base data
func (self *BaseData) showValue(start int) (arr []string) {
	length := 12
	end := start + length
	template := "%-10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s\n"

	title := make([]Bpoint, 0)

	for i := start; i < end; i++ {
		title = append(title, Bpoint(i))
	}

	printBp := func(title string, val []Bpoint) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}
	printPoint := func(title string, val []Point) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}

	arr = append(arr, printBp("Col", title))
	arr = append(arr, fmt.Sprintln(strings.Repeat("-", 145)))
	//arr = append(arr, printBp("Inst", self.Inst[start:end]))
	arr = append(arr, printBp("BP", self.Bp[start:end]))
	arr = append(arr, printBp("New BP", self.Nbp[start:end]))

	for i := 0; i < ROWS; i++ {
		arr = append(arr, printPoint(strconv.Itoa(i+1), self.Data[i][start:end]))
	}

	arr = append(arr, printPoint("XG", self.Xg[start:end]))
	arr = append(arr, printPoint("GZ", self.Gz[start:end]))
	arr = append(arr, printPoint("GF", self.Gf[start:end]))
	arr = append(arr, printPoint("GF1", self.Gf1[start:end]))
	return arr
}

func (self BaseData) String() string {

	/*
		print := func(title string, val Duck) string {
			list := reflect.ValueOf(val)
			v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11 := list.Index(0), list.Index(1),
				list.Index(3), list.Index(4), list.Index(5), list.Index(6),
				list.Index(7), list.Index(8), list.Index(9), list.Index(10),
				list.Index(11)
			ret := fmt.Sprintf(template, title, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11)
		}
	*/
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	arr = self.showValue(35)
	ret += strings.Join(arr, "")
	return ret

}

// for group data
func (self *GroupData) showValue(start int) (arr []string) {
	length := 12
	end := start + length
	template := "%-10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s\n"

	title := make([]Bpoint, 0)

	for i := start; i < end; i++ {
		title = append(title, Bpoint(i))
	}

	printBp := func(title string, val []Bpoint) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}
	printPoint := func(title string, val []Point) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}

	arr = append(arr, printBp("Col", title))
	arr = append(arr, fmt.Sprintln(strings.Repeat("-", 145)))

	arr = append(arr, printBp("Bp", self.Bp[start:end]))
	arr = append(arr, printBp("Inst", self.Inst[start:end]))

	arr = append(arr, printPoint("ZG", self.Zg[start:end]))

	arr = append(arr, printPoint("GZ", self.Gz[start:end]))
	arr = append(arr, printPoint("GZMM", self.Gzmm[start:end]))

	arr = append(arr, printPoint("GF", self.Gf[start:end]))
	arr = append(arr, printPoint("GFMM", self.Gfmm[start:end]))

	arr = append(arr, printPoint("GF1", self.Gf1[start:end]))

	return arr
}

func (self *GroupData) String() string {
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	arr = self.showValue(35)
	ret += strings.Join(arr, "")
	return ret
}

// for big data
func (self *BigData) showValue(start int) (arr []string) {
	var j Value
	length := 12
	end := start + length
	template := "%-10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s\n"

	title := make([]Bpoint, 0)

	for i := start; i < end; i++ {
		title = append(title, Bpoint(i))
	}

	printBp := func(title string, val []Bpoint) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}
	printPoint := func(title string, val []Point) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}

	arr = append(arr, printBp("Col", title))
	arr = append(arr, fmt.Sprintln(strings.Repeat("-", 145)))

	arr = append(arr, printBp("Inst", self.Inst[start:end]))

	arr = append(arr, printPoint("DG", self.Dg[start:end]))

	arr = append(arr, printPoint("GZ", self.Gz[start:end]))
	arr = append(arr, printPoint("GZMM", self.Gzmm[start:end]))

	arr = append(arr, printPoint("GF", self.Gf[start:end]))
	arr = append(arr, printPoint("GFMM", self.Gfmm[start:end]))

	arr = append(arr, printPoint("GF1", self.Gf1[start:end]))

	arr = append(arr, printPoint("TS-zmm", self.TsValue[0][start:end]))
	arr = append(arr, printPoint("TS-fmm", self.TsValue[1][start:end]))
	arr = append(arr, printPoint("TS-f1", self.TsValue[2][start:end]))

	// only show zmm 9
	for j = 0; j < THREESOME_NUM_SHOW; j++ {
		arr = append(arr, printPoint("Up", self.TsData[0][j].Up[start:end]))
		arr = append(arr, printPoint("V1", self.TsData[0][j].V1[start:end]))
		arr = append(arr, printPoint("V2", self.TsData[0][j].V2[start:end]))
		arr = append(arr, printPoint("V3", self.TsData[0][j].V3[start:end]))
		arr = append(arr, printPoint("Sum", self.TsData[0][j].Sum[start:end]))
	}

	return arr
}

func (self *BigData) String() string {
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	arr = self.showValue(35)
	ret += strings.Join(arr, "")
	return ret
}

// for threesome
func (self *ThreeSome) showValue(start int) (arr []string) {
	length := 12
	end := start + length
	template := "%-10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s %10s\n"

	title := make([]Bpoint, 0)

	for i := start; i < end; i++ {
		title = append(title, Bpoint(i))
	}

	printBp := func(title string, val []Bpoint) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}
	printPoint := func(title string, val []Point) string {
		return fmt.Sprintf(template, title, val[0], val[1], val[2], val[3],
			val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])
	}

	arr = append(arr, printBp("Col", title))
	arr = append(arr, printPoint("Up", self.Up[start:end]))

	arr = append(arr, fmt.Sprintln(strings.Repeat("-", 145)))

	arr = append(arr, printPoint("V1", self.V1[start:end]))
	arr = append(arr, printPoint("V2", self.V2[start:end]))

	arr = append(arr, printPoint("V3", self.V3[start:end]))
	arr = append(arr, printPoint("Sum", self.Sum[start:end]))

	return arr
}

func (self *ThreeSome) String() string {
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	return ret
}
