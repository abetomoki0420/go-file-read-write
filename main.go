package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dict struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func main() {
	raw, err := ioutil.ReadFile("./dict.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dict Dict

	json.Unmarshal(raw, &dict)

	fmt.Println(dict)

	dict.A = dict.A + 1

	b, _ := json.Marshal(dict)

	ioutil.WriteFile("./dict.json", b, os.ModePerm)
}
