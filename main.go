package main

import (
	"github.com/tyrchen/calcapp/calc"
	"github.com/tyrchen/calcapp/network"
	"fmt"
	"unsafe"
)

func inspect() {
	var bp calc.Bpoint
	var value calc.Value
	var point calc.Point
	var data calc.BaseData
	var group calc.GroupData
	fmt.Printf("sizeof Bpoint: %d\n", unsafe.Sizeof(bp))
	fmt.Printf("sizeof Value: %d\n", unsafe.Sizeof(value))
	fmt.Printf("sizeof Point: %d\n", unsafe.Sizeof(point))
	fmt.Printf("sizeof BaseData: %d\n", unsafe.Sizeof(data))
	fmt.Printf("sizeof GroupData: %d\n", unsafe.Sizeof(group))
}

func getArpTable() {
	fmt.Println("Arp table: \n", network.ParseArpTable())
}

func main() {
	getArpTable()
	inspect()
	//calc.Init("basepoint00.dat")
}
