package calc

import (
	"bufio"
	"fmt"
	"os"
)

var (
	Values GroupData
)

func init_bp(filename string) {
	var f, err = os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var row = 0
	for s.Scan() {
		var line = s.Text()
		for col := 0; col < COLS; col++ {
			Values.Data[row].Inst[col] = Bpoint(line[col])
		}
		fmt.Println(line)
	}

	if err := s.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
}

func Init(filename string) {
	init_bp(filename)
}
