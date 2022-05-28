package main

import (
	"fmt"
	"sync"
	"time"
)

// The code below has 2 goroutines. One goroutine is incrementing the value of x by 1. The other goroutine is printing the value of x. The printed value of x is not guaranteed to follow an expected sequence due to interleaving of the goroutines.

// A global variable that will be shared between 2 goroutines
var x = 1

// Enables us to wait for multiple goroutines to finish
var wg sync.WaitGroup

func goroutine1() {
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		x = x + 1
	}
	// Tell the WaitGroup that this goroutine is done
	defer wg.Done()
}

func goroutine2() {
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("x:", x)
	}
	defer wg.Done()
}

func main() {
	// Increment the WaitGroup counter for each goroutine
	wg.Add(1)
	go goroutine1()
	wg.Add(1)
	go goroutine2()
	wg.Wait()
}
