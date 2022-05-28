package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port is closed or filtered
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// Create a buffered channel with a capacity of 100
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	fmt.Println("Start scanning...")

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

	fmt.Println("Scan finished.")
}
