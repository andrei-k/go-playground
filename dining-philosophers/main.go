package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doStuff(ch1, ch2 chan int) {
	<-ch1    // Wait for the first channel
	ch2 <- 1 // Write to the second channel
	wg.Done()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// for i := 0; i < 10; i++ {
	wg.Add(2)
	// This will result in a deadlock because the channels will be waiting for each other in the goroutine
	go doStuff(ch1, ch2)
	go doStuff(ch2, ch1)
	msg1 := <-ch1
	msg2 := <-ch2
	fmt.Println("m1:", msg1, "m2:", msg2)
	// }
	wg.Wait()
}
