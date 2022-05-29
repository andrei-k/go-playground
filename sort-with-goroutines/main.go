package main

import (
	"bufio"
	"fmt"
	"os"
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

func sort(part []int) {
	fmt.Println("part:", part)
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

	fmt.Println("chunks:", chunks)

	if len(chunks) > 0 {
		fmt.Println("length:", len(chunks))

		// Send each partition to goroutines for sorting
		sort(chunks[0])

		// Then the main goroutine merges the 4 sorted subarrays into one sorted array
	}
}
