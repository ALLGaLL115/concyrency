package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	Announce("First", 3*time.Second, &wg)
	Announce("Second", 100*time.Millisecond, &wg)

	wg.Wait()

}

func Announce(message string, delay time.Duration, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		time.Sleep(delay)
		fmt.Println(message)
	}() // Note the parentheses - must call the function.
}

func Convertation_json_with_structures() {
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

func Work_with_interface_and_type_assertion() {
	var i interface{}
	i = "dafsd"
	i = 2342
	i = 3.1

	r := i.(float64)
	fmt.Print("the circle of area is ", float64(int(math.Pi*r*r*100))/100)
}
