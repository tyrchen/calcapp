package calc

import (
	"calcapp/utils"
)

var (
	Values GroupData
)

func LoadBp(index uint) {
	values := utils.LoadBpFile(index)
	for row, value := range values {
		data := utils.ValueToBp(value)
		Values.Data[row].LoadBp(data)
	}
}
