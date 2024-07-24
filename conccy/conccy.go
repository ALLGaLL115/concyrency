package main

import (
	"fmt"
	"sync"
	"time"
)

func squares(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Print(num)
	}
	wg.Done()

}

func main() {
	DoHardWorkSync()
	DoHardWorkConcyrency()
}

func startWorker(tasks <-chan int, responses chan<- map[string]int, id int) {
	for task := range tasks {
		res := make(map[string]int)
		for i := 0; i < 1000000; i++ {
			res["num"] = 123456789876 * task
		}
		res["id"] = id
		responses <- res
	}

}

func DoHardWorkConcyrency() {
	start := time.Now()

	tasks := make(chan int, 99)
	respnoses := make(chan map[string]int, 99)

	for i := 0; i < 10; i++ {
		go startWorker(tasks, respnoses, i)
	}

	for i := 0; i < 100; i++ {
		tasks <- i
	}
	results := make([]map[string]int, 100)
	for i := 0; i < 100; i++ {
		result := <-respnoses
		results = append(results, result)
	}

	fmt.Printf("DoHardWorkConcyrency working time is %v \n", time.Since(start))
	// fmt.Print(results)

}

func DoHardWorkSync() {
	start := time.Now()
	results := make([]map[string]int, 100)

	for i := 0; i < 100; i++ {
		res := make(map[string]int)
		for k := 0; k < 1000000; k++ {
			res["num"] = 123456789876 * i
		}
		res["id"] = i
		results = append(results, res)
	}

	fmt.Printf("DoHardWorSync working time is %v \n", time.Since(start))
	// fmt.Print(results)

}

func with_mutex() {
	c := make(chan string)
	var my sync.Mutex
	var wg sync.WaitGroup
	m := make(map[string]string)
	m["name"] = "Hello"

	wg.Add(1)
	go func(mu *sync.Mutex, wg *sync.WaitGroup) {
		mu.Lock()
		m["name"] = <-c
		mu.Unlock()
		wg.Done()

	}(&my, &wg)

	c <- "HEhehe"
	wg.Wait()
	fmt.Println("Hello,", m["name"])
}
