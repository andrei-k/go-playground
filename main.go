package main

import (
	"fmt"
	"myapp/doctor"
)

func main() {
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}
