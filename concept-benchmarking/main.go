package main

import (
	"fmt"
	"math"
)

// Determine if a number is a prime number by checking whether it is divisible by a number between two and its square root
func primeNumbers(max int) (primes []int) {
	for i := 2; i < max; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}
	return
}

func main() {
	fmt.Println("primeNumbers(10000):", primeNumbers(10000))
}
