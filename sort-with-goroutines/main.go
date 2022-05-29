package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() string {
	fmt.Println("Type a series of integers separated by spaces")
	fmt.Print("-> ")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func sortPartition(partition []int, result chan []int) {
	sort.Ints(partition)
	fmt.Println("Sorted partition:", partition)
	result <- partition
}

func main() {
	userInput := getInput()
	intString := strings.Split(userInput, " ")

	chunks := make(map[int][]int)
	var key int

	for i, val := range intString {
		integer, err := strconv.Atoi(val)
		// Inputs that aren't integers will be ignored
		if err == nil {
			// Create 4 partitions of integers
			key = i % 4
			chunks[key] = append(chunks[key], integer)
		}
	}

	// Create a buffered channel with a capacity of 4
	result := make(chan []int, 4)

	if len(chunks) > 0 {
		var final []int

		// Send each partition to goroutines for sorting
		for i := 0; i < len(chunks); i++ {
			go sortPartition(chunks[i], result)
		}

		// Receive the sorted partitions from the goroutines
		for i := 0; i < len(chunks); i++ {
			final = append(final, <-result...)
		}

		// Sort the final slice
		// Note that the point of this program is to practice goroutines rather than finding most efficient ways to merge slices into a sorted one
		sort.Ints(final)
		fmt.Println("final:", final)
	}
}
