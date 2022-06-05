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

func (p *Philosopher) Eat() {
	// Each philosopher should eat 3 times
	for i := 0; i < 3; i++ {
		fmt.Printf("Philosopher %d is starting to eat\n", p.number)
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		time.Sleep(500 * time.Millisecond)
		p.eatCount++
		fmt.Printf("Philosopher %d finished eating %d/3\n", p.number, p.eatCount)
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// Initialize chopsticks
	var chopsticks = make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		// Assign type *Chopstick
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

	for i := 0; i < 5; i++ {
		wg.Add(1)
		// Each philosopher starts eating
		go philosophers[i].Eat()
	}

	wg.Wait()
}
