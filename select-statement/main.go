package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan string)
var chan2 = make(chan string)

// The purpose of this channel is to abort the receiver loop from listening to channels.
// It’s good practice to use an empty struct as the type of channel. Because they are useless for anything other than signalling, it expresses intent in the code and it doesn't use much memory.
var abort = make(chan struct{})

// Pause for 1 second and send text to channel 1
func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "one"
}

// Pause for 2 seconds and send text to channel 2
func task2() {
	time.Sleep(2 * time.Second)
	chan2 <- "two"
}

func task3() {
	chan1 <- "111"
	chan2 <- "222"
	chan1 <- "333"
	close(abort)
	chan2 <- "444"
	chan1 <- "555"
}

func main() {
	// Start tasks as goroutines
	go task1()
	go task2()

	// The select statement allows you to wait for the first data from a set of channels, it's first-come, first-served and unblocks once it receives from any channel.
	select {
	case msg1 := <-chan1:
		fmt.Println("received", msg1)
	case msg2 := <-chan2:
		fmt.Println("received", msg2)
	}

	fmt.Println("-------------------")

	go task1()
	go task2()

	// In this case, the for loop makes sure information is received from both channels.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("received", msg1)
		case msg2 := <-chan2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("-------------------")

	go task1()
	go task2()
	go task3()

	// In this case, the for loop listens for a special abort channel. When it's closed, the select statement exits.
	for {
		select {
		case msg1 := <-chan1:
			fmt.Println("received", msg1)
		case msg2 := <-chan2:
			fmt.Println("received", msg2)
		case <-abort:
			return
		}
	}
}
