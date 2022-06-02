package main

import (
	"fmt"
	"sync"
)

// Counter struct that also holds the mutex used to lock/unlock the counter variable so it can be used by one goroutine instance at a time
type Counter struct {
	mutex   sync.Mutex
	counter int
}

// This is a bad example where the counter will likely be incremented out of order due to non-deterministic interleaving
func (c *Counter) increment1(wg *sync.WaitGroup) {
	c.counter++
	fmt.Println("x:", c.counter)
	wg.Done()
}

// The counter will be incremented in order because of the mutex
func (c *Counter) increment2(wg *sync.WaitGroup) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter++
	fmt.Println("x:", c.counter)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	c := Counter{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go c.increment1(&wg)
	}
	wg.Wait()

	fmt.Println("-------------------")
	c = Counter{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go c.increment2(&wg)
	}
	wg.Wait()
}
