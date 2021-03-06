package calc

const (
	COLS       = 55 + 1
	ROWS       = 9
	GROUP_SIZE = 37000
	STOP_COL   = 36
	STOP_VALUE = 1048575

	// for concurrency
	CHUNKS     = 100
	CHUNK_SIZE = GROUP_SIZE / CHUNKS

	// for multiplier
	MUL_COND    = 600000
	MUL_STOP    = 63
	ZG_NUM      = 60
	ZG_NUM_SHOW = 0

	// threesome total num
	THREESOME_NUM      = 9
	THREESOME_NUM_SHOW = 0
	THREESOME_TOTAL    = 3
)

type Bpoint uint8
type Value int
type Duck interface{}

type Point struct {
	T bool
	V Value
}

type Env struct {
	CurrentCol Value
	Last       Point
	Bp         Bpoint
	Stop       bool
}

type BaseData struct {
	//Inst [COLS]Bpoint
	Bp   [COLS]Bpoint
	Nbp  [COLS]Bpoint
	Data [ROWS][COLS]Point
	Xg   [COLS]Point
	Gz   [COLS]Point
	Gf   [COLS]Point
	Gf1  [COLS]Point
}

type GroupData struct {
	Index uint
	Gzm   Value
	Gfm   Value
	Bp    [COLS]Bpoint
	Inst  [COLS]Bpoint
	Data  [GROUP_SIZE]BaseData
	Zg    [COLS]Point
	Gz    [COLS]Point
	Gzmm  [COLS]Point
	Gf    [COLS]Point
	Gfmm  [COLS]Point
	Gf1   [COLS]Point
}

type ThreeSome struct {
	Up  [COLS]Point
	V1  [COLS]Point
	V2  [COLS]Point
	V3  [COLS]Point
	Sum [COLS]Point
}

type BigData struct {
	Inst    [COLS]Bpoint
	Dg      [COLS]Point
	Gz      [COLS]Point
	Gzmm    [COLS]Point
	Gf      [COLS]Point
	Gfmm    [COLS]Point
	Gf1     [COLS]Point
	TsValue [THREESOME_TOTAL][COLS]Point
	TsRet   [COLS]Point
	TsData  [THREESOME_TOTAL][THREESOME_NUM]ThreeSome
	Data    [ZG_NUM]GroupData
}
