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
}

type BaseData struct {
	Inst [COLS]Bpoint
	Bp   [COLS]Bpoint
	Nbp  [COLS]Bpoint
	Data [ROWS][COLS]Point
	Xg   [COLS]Point
	Gz   [COLS]Point
	Gf   [COLS]Point
	Gf1  [COLS]Point
}

type GroupData struct {
	Inst [COLS]Bpoint
	Data [GROUP_SIZE]BaseData
	Zg   [COLS]Point
	Gz   [COLS]Point
	Gf1  [COLS]Point
}
