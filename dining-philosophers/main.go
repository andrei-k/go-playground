package main

import (
	"fmt"
	"sync"
)

// 5 philosophers sitting at a round table
// 1 chopstick is placed between each adjacent pair
// Want to eat rice from their plate, but needs 2 chopsticks
// Only one philosopher can hold a chopstick at a time
// Not enough chopsticks for everyone to eat at once

// Each chopstick is a mutex
// Each philosopher is associated with a goroutine and two chopsticks

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

// TODO: In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// TODO: The host allows no more than 2 philosophers to eat concurrently.

func (p Philosopher) eat() {
	// Each philosopher should eat 3 times
	for i := 0; i < 3; i++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("Starting to eat %d\n", p.number)
		fmt.Printf("Finishing eating %d\n", p.number)
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// Initialize chopsticks and philosophers
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		// Use modulo operator because there's a cyclical relationship of chopsticks (the fifth philosopher gets chopstick zero for one of them)

		// TODO: The philosophers pick up the chopsticks in any order, not lowest-numbered first
		// Each philosopher is numbered, 1 through 5
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	// Each philosopher starts eating
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat()
	}

	wg.Wait()
}
