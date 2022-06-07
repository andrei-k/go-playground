package main

import (
	"fmt"
	"sync"
	"time"
)

// Each chopstick is a mutex
type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number         int
	eatCount       int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p *Philosopher) Eat(hostCh chan bool) {
	// Each philosopher should eat 3 times
	for i := 0; i < 3; i++ {
		// Block channel while a philosopher is eating
		<-hostCh

		fmt.Printf("Philosopher %d is eating\n", p.number)
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		time.Sleep(time.Second) // Add a little delay
		p.eatCount++
		fmt.Printf("Philosopher %d finished eating %d/3\n", p.number, p.eatCount)
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		// Send a value and free up the channel
		hostCh <- true
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// Initialize chopsticks
	var chopsticks = make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		// Assign type *Chopstick
		// new() function creates a variable and returns a pointer to that variable
		chopsticks[i] = new(Chopstick)
	}

	// Initialize philosophers
	philosophers := make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		leftStick := chopsticks[i]
		rightStick := chopsticks[(i+1)%5]

		// Philosophers are numbered 1 through 5
		philosophers[i] = Philosopher{
			number:         i + 1,
			eatCount:       0,
			leftChopstick:  leftStick,
			rightChopstick: rightStick,
		}
	}

	// Host allows 2 philosophers to eat concurrently
	hostCh := make(chan bool, 2)
	hostCh <- true
	hostCh <- true

	// Each philosopher starts eating
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].Eat(hostCh)
	}

	wg.Wait()
}
