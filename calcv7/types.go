package calcv7

import (
	. "calcapp/utils"
)

const (
	COLS = 56 + 1
	ROWS = 129 //9
	//GROUP_SIZE = 9 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 // G20
	// GROUP_SIZE = 9 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3// G10
	GROUP_SIZE = 20
	//GROUP_SIZE = 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 // G9
	LAYER      = 11
	STOP_COL   = 0
	STOP_VALUE = 2047

	CONCURRENCY = 1

	CHUNK_SIZE = 3
)

type Env struct {
	CurrentCol Value
	Last       Point // 前一列同行
	S1         Point // 前一列上一行，用于计算符号
	S2         Point // 同列上一行，用于计算符号
	Stop       bool
}

type BaseData struct {
	//Inst [COLS]Bpoint
	Bp   [COLS]Bpoint
	Zbp  [COLS]Point
	Data [ROWS][COLS]Point
	Ag   [COLS]Point // A果
	G1   [COLS]Point // 果1
	G137 [COLS]Point
}

type BaseGroup struct {
	Data [GROUP_SIZE]BaseData
	Zg2  [COLS]Point
	G137 [COLS]Point
	Gz   [COLS]Point
	Gzr  [6][COLS]Point
	Gf   [COLS]Point
	Gfr  [6][COLS]Point
}