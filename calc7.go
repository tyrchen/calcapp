package main

import (
	. "calcapp/calcv6"
	. "calcapp/utils"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

const (
	//BASE_DATA_VALUE_ROWS = 1 + CHUNK_SIZE + 1
	BASE_DATA_VALUE_ROWS = 2 + 1 + GROUP_SIZE
)

type CalcData struct {
	Method string
	Pos    Value
	Values [2][BASE_DATA_VALUE_ROWS]Point
}

var (
	values *BaseGroup
)

func calcHandler(ws *websocket.Conn) {
	var err error

	// send initial data
	sendData(0, ws)

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

func getValues(pos Value) (ret CalcData) {
	var i, j Value
	ret.Method = "calc"
	ret.Pos = pos

	data := values

	for i = 0; i < 2; i++ {
		ret.Values[i][1] = data.G137[pos+i]
		ret.Values[i][2] = data.Gz[pos+i]
		
		for j =0; j < GROUP_SIZE; j++ {
			ret.Values[i][j+2] = data.Data[j].G137[pos+i]
		}
	}

	return ret

}

func initValues() {
	values = new(BaseGroup)
	values.Init()
}

func clear() {
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <port>", os.Args[0])
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	sport, _ := strconv.Atoi(os.Args[1])
	port := uint(sport)

	initValues()

	http.Handle("/", websocket.Handler(calcHandler))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
