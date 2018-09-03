package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type transactions interface{}

type trans struct {
	Result bool `json:"result"`
	Tran   []struct{ transactions }
}
type logWriter struct{}

var jsondata = ""

func main() {
	resp, err := http.Get("http://localhost:9000/transactions.json")
	if err != nil {
		fmt.Println("Err:", err)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	var transactionss trans
	if err := json.Unmarshal([]byte(jsondata), &transactionss); err != nil {
		panic(err)
	}
	fmt.Println(transactionss.Tran)
}

func (logWriter) Write(bs []byte) (int, error) {
	jsondata = string(bs)
	return len(bs), nil
}
