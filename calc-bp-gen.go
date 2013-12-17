package main

import (
	"calcapp/calc"
	"calcapp/utils"
)

func generateBp() {
	utils.GenerateBpFiles(30)
}

func main() {
	generateBp()
	calc.LoadBp(30)
}
