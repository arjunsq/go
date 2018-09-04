package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type values interface{}

type trans struct {
	Result bool                `json:"result"`
	Tran   []map[string]values `json:"transactions"`
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

}

func (logWriter) Write(bs []byte) (int, error) {
	jsondata = string(bs)
	var m trans
	dec := json.NewDecoder(strings.NewReader(jsondata))
	for {

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, v := range m.Tran {

			productsJSON, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				fmt.Println("Error marshalling to JSON:", err)
			}
			fmt.Println(string(productsJSON))
			fmt.Printf("%T:%v", v["br_rg_date_id"], v["br_rg_date_id"])
			filename := v["br_rg_date_id"].(string)
			err = ioutil.WriteFile(filename, productsJSON, 0644)
			if err != nil {
				fmt.Println("Error writing JSON to file:", err)
			}
			fmt.Println("CONVERTED JSON IS WRITTEN TO JSON")
		}

	}

	return len(bs), nil
}
