package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")

		// Capture user input until the return character is pressed
		userInput, _ := reader.ReadString('\n')

		// Replace return input on Windows
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		// Replace return input on Mac and Linux
		userInput = strings.Replace(userInput, "\n", "", -1)
		// (-1 indicates to replace everywhere the match it occurs)

		if userInput == "quit" {
			break
		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}
}
