package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Ano", "Hehe", 1294706395881547000}

	b, err := json.Marshal(m)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(b)

	var res Message

	err = json.Unmarshal(b, &res)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(res)

}
