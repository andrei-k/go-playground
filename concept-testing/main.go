package main

import "fmt"

func Calc(x int) (result int) {
	result = x + 10
	return result
}

func main() {
	fmt.Println("Calc(10):", Calc(10))
}
