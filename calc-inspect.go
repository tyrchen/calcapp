package main

import (
	. "calcapp/calc"
	"fmt"
	"unsafe"
)

func inspect() {
	var bp Bpoint
	var value Value
	var point Point
	var data BaseData
	group := new(GroupData)
	fmt.Printf("sizeof Bpoint: %d\n", unsafe.Sizeof(bp))
	fmt.Printf("sizeof Value: %d\n", unsafe.Sizeof(value))
	fmt.Printf("sizeof Point: %d\n", unsafe.Sizeof(point))
	fmt.Printf("sizeof BaseData: %d\n", unsafe.Sizeof(data))
	fmt.Printf("sizeof GroupData: %d\n", unsafe.Sizeof(*group))
}

func main() {
	inspect()
}
