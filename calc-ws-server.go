package main

import (
	. "calcapp/calc"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"strconv"
	//"strings"
)

type CalcData struct {
	Method string
	Col    Value
	Values [2][3]Point
}

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
	Values.Run(inst, col)
	fmt.Println(Values)
	return getValues(col)
}

func getValues(col Value) (ret CalcData) {
	ret.Method = "calc"
	ret.Col = col
	ret.Values[0] = [3]Point{Values.Zg[col], Values.Gz[col], Values.Gf1[col]}
	ret.Values[1] = [3]Point{Values.Zg[col+1], Values.Gz[col+1], Values.Gf1[col+1]}

	return ret
}

func initValues() {
	Values.LoadBp(1)
	Values.Init()
}

func clear() {
	Values.Clear()
	initValues()
}

func main() {
	initValues()

	http.Handle("/", websocket.Handler(calcHandler))

	if err := http.ListenAndServe(":8210", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
