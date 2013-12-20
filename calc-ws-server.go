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

type CalcData struct {
	Method string
	Col    Value
	Values [2][4]Point
}

var (
	values *GroupData
)

func calcHandler(ws *websocket.Conn) {
	var err error

	// send initial data
	ret := getValues(0)
	b, _ := json.Marshal(ret)
	wsSend(ws, string(b))

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

func runCommand(ws *websocket.Conn, reply string) {
	var v interface{}
	json.Unmarshal([]byte(reply), &v)
	m := v.(map[string]interface{})
	switch m["method"].(string) {
	case "calc":
		inst, pos := Bpoint(m["inst"].(float64)), Value(m["pos"].(float64))

		ret := calc(inst, pos)
		b, _ := json.Marshal(ret)
		wsSend(ws, string(b))
	case "close":
		clear()
		wsSend(ws, `{"method":"close","value":"true"}`)
	}
}

func calc(inst Bpoint, col Value) CalcData {
	values.Run(inst, col)
	return getValues(col)
}

func getValues(col Value) (ret CalcData) {
	ret.Method = "calc"
	ret.Col = col
	ret.Values[0] = [4]Point{values.Zg[col], values.Gz[col], values.Gf[col], values.Gf1[col]}
	ret.Values[1] = [4]Point{values.Zg[col+1], values.Gz[col+1], values.Gf[col+1], values.Gf1[col+1]}

	return ret
}

func initValues(index uint) {
	values = new(GroupData)
	values.LoadBp(index)
	values.Init()
}

func clear() {
	index := values.Index
	values.Clear()
	initValues(index)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <index>", os.Args[0])
		os.Exit(1)
	}

	val, _ := strconv.Atoi(os.Args[1])
	index := uint(val)

	initValues(index)

	http.Handle("/", websocket.Handler(calcHandler))

	if err := http.ListenAndServe(":8210", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
