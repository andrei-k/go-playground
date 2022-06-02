package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func setup() {
	fmt.Println("Setting up")
}

func doStuff(counter *int) {
	// Although this is included in every goroutine, it'll only run once
	once.Do(setup)
	*counter++
	fmt.Println("Doing stuff", *counter)
	wg.Done()
}

func main() {
	var counter int = 0
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go doStuff(&counter)
	}
	wg.Wait()
}
