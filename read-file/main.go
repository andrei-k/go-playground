package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

// Define the struct for a name
type Name struct {
	fname string
	lname string
}

// Define the slice of structs holding names
var names []Name

func main() {
	// Prompt user for the name of the text file
	reader = bufio.NewReader(os.Stdin)
	fileName := readString("Enter the name of the text file")

	// Read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// Read each line of the text file and create a struct with the first and last names found in the file
	for _, line := range strings.Split(string(data), "\n") {
		name := Name{
			fname: strings.Split(line, " ")[0],
			lname: strings.Split(line, " ")[1],
		}

		// Add each struct to a slice
		names = append(names, name)
	}

	// Iterate through the slice of structs and print the first and last names found in each struct
	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}

func readString(s string) string {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput != "" {
			return userInput
		}
	}
}
