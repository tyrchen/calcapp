package calcv5

import (
	"fmt"
	"strconv"
	"strings"
	. "calcapp/utils"
)


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
	arr = append(arr, printPoint("ZBP", self.Zbp[start:end]))

	for i := 0; i < ROWS; i++ {
		arr = append(arr, printPoint(strconv.Itoa(i+1), self.Data[i][start:end]))
	}

	arr = append(arr, printPoint("AG", self.Ag[start:end]))
	arr = append(arr, printPoint("G1", self.G1[start:end]))
	return arr
}

func (self BaseData) String() string {
	arr := self.showValue(0)
	ret := strings.Join(arr, "") + "\n"

	arr = self.showValue(35)
	ret += strings.Join(arr, "")
	return ret

}