package main

import (
	. "calcapp/calc"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	BASE_DATA_VALUE_ROWS = 3 + ROWS + 4
)

type CalcData struct {
	Method string
	Pos    Value
	Values [2][7]Point
}

type BaseDataValue struct {
	Method string
	Pos    Value
	Values [2][BASE_DATA_VALUE_ROWS]Point
}

type MultiplierData struct {
	Method string
	Values [2]int
}

type BpZgData struct {
	Method string
	Values [COLS]Bpoint
}

var (
	values *GroupData
)

func calcHandler(ws *websocket.Conn) {
	var err error

	// send initial data
	sendData(0, ws)
	sendBpZg(ws)
	sendMultiplier(ws)

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			clear()
			break
		}

		fmt.Println("Received back from client: " + reply)

		runCommand(ws, reply)
	}
}

func wsSend(ws *websocket.Conn, msg string) {
	if err := websocket.Message.Send(ws, msg); err != nil {
		fmt.Println("Can't send")
		clear()
	}
}

func sendData(pos Value, ws *websocket.Conn) {
	b, _ := json.Marshal(getValues(pos))
	wsSend(ws, string(b))
	// send one of the BaseData
	//b1, _ := json.Marshal(getBaseDataValue(pos))
	//wsSend(ws, string(b1))
}

func sendMultiplier(ws *websocket.Conn) {
	var data MultiplierData
	data.Method = "multiplier"
	data.Values[0] = int(values.Gzm)
	data.Values[1] = int(values.Gfm)
	b, _ := json.Marshal(data)
	wsSend(ws, string(b))
}

func sendBpZg(ws *websocket.Conn) {
	var data BpZgData
	data.Method = "bpzg"
	for i := 0; i < COLS; i++ {
		data.Values[i] = values.Bp[i]
	}
	b, _ := json.Marshal(data)
	wsSend(ws, string(b))
}

func runCommand(ws *websocket.Conn, reply string) {
	var v interface{}
	json.Unmarshal([]byte(reply), &v)
	m := v.(map[string]interface{})
	switch m["method"].(string) {
	case "calc":
		inst, pos := Bpoint(m["inst"].(float64)), Value(m["pos"].(float64))
		calc(inst, pos)
		sendData(pos, ws)

	case "close":
		clear()
		wsSend(ws, `{"method":"close","value":"true"}`)
	}
}

func calc(inst Bpoint, pos Value) {
	values.Run(inst, pos)

}

func getBaseDataValue(pos Value) (ret BaseDataValue) {
	var i Value
	data := values.Data[GROUP_SIZE-1]
	ret.Pos = pos
	ret.Method = "xg"
	for i = 0; i < 2; i++ {
		//ret.Values[i][0] = Point{false, Value(data.Inst[pos+i])}
		ret.Values[i][1] = Point{false, Value(data.Bp[pos+i])}
		ret.Values[i][2] = Point{false, Value(data.Nbp[pos+i])}

		for j := 0; j < ROWS; j++ {
			ret.Values[i][j+3] = data.Data[j][pos+i]
		}

		ret.Values[i][12] = data.Xg[pos+i]
		ret.Values[i][13] = data.Gz[pos+i]
		ret.Values[i][14] = data.Gf[pos+i]
		ret.Values[i][15] = data.Gf1[pos+i]
	}

	return ret
}

func getValues(pos Value) (ret CalcData) {
	var i Value
	ret.Method = "calc"
	ret.Pos = pos

	for i = 0; i < 2; i++ {
		inst := Point{false, Value(values.Inst[pos+i])}
		ret.Values[i] = [7]Point{
			inst, values.Zg[pos+i], values.Gz[pos+i], values.Gzmm[pos+i],
			values.Gf[pos+i], values.Gfmm[pos+i], values.Gf1[pos+i]}
	}

	return ret
}

func initValues(index uint) {
	values = new(GroupData)
	values.Init(index)
}

func clear() {
	index := values.Index
	values.Clear()
	initValues(index)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <port> <index>", os.Args[0])
		os.Exit(1)
	}

	sport, _ := strconv.Atoi(os.Args[1])
	val, _ := strconv.Atoi(os.Args[2])
	port := uint(sport)
	index := uint(val)

	initValues(index)

	http.Handle("/", websocket.Handler(calcHandler))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
