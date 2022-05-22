package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// Synchronized counter
	var wg sync.WaitGroup

	for i := 1; i <= 65535; i++ {
		wg.Add(1) // Increment counter for each new goroutine

		go func(j int) {
			defer wg.Done() // Decrement counter when goroutine completes

			// Scan all ports on localhost
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// Port is closed or filtered
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	// Wait for all goroutines to finish
	wg.Wait()
}
