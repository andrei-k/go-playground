package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter in a sequence of integers separated by spaces")
	fmt.Print("-> ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// Strip off the carriage return
	input = strings.Replace(input, "\r\n", "", -1) // Windows
	input = strings.Replace(input, "\n", "", -1)   // Mac

	strings := strings.Split(input, " ")

	ints := make([]int, len(strings))
	for i, v := range strings {
		ints[i], _ = strconv.Atoi(v)
	}

	BubbleSort(ints)
	fmt.Println(ints)
}

// BubbleSort() takes a slice of integers as an argument and returns nothing It modifies the slice so that the elements are in sorted order.
func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1-i; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

// Swap() function takes a slice of integers and an index value i which indicates a position in the slice. 	The function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}
