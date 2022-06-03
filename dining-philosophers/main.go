package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

func (p *Philosopher) Eat() {
	// Each philosopher should eat 3 times
	for i := 0; i < 3; i++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("Starting to eat:  %d\n", p.number)
		fmt.Printf("Finishing eating: %d\n", p.number)
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// Initialize chopsticks and philosophers
	var chopsticks = make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		// Assign type *Chopstick
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]Philosopher, 5)

	// Randomly generate 1 or 0
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(2)
	fmt.Println("i", i)

	for i := 0; i < 5; i++ {
		// TODO: The philosophers pick up the chopsticks in any order, not lowest-numbered first
		stick1 := chopsticks[i]
		// Use modulo operator because there's a cyclical relationship of chopsticks (the fifth philosopher gets chopstick zero for one of them)
		stick2 := chopsticks[(i+1)%5]

		philosophers[i] = Philosopher{
			number:         i + 1, // Philosophers are numbered 1 through 5
			leftChopstick:  stick1,
			rightChopstick: stick2,
		}
	}

	// Each philosopher starts eating
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].Eat()
	}

	wg.Wait()
}
