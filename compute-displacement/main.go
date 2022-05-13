package main

import (
	"fmt"
)

func main() {
	var a float64  // Acceleration
	var vo float64 // Initial velocity
	var so float64 // Initial displacement

	fmt.Print("Enter values for acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter values for initial velocity: ")
	fmt.Scan(&vo)
	fmt.Print("Enter values for initial displacement: ")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}
