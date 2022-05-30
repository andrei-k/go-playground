package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// TODO: explain purpose of this struct
type Channel struct {
	samples int
	result  float64
}

func calcPi(samples int) float64 {
	var inside int = 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < samples; i++ {
		// Generate a random x and y value between 0.0 and 1.0
		x, y := r.Float64(), r.Float64()
		// Use the equation of a unit circle to check if the generated random sample fell inside the circle
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	// Divide the count of samples inside the circle by total samples, which should equal to roughly π/4
	ratio := float64(inside) / float64(samples)

	//  Multiply by 4 to get our final estimate for π
	return ratio * 4
}

// Calculate Pi using concurrent goroutines
func calcPiConcurrent(samples int, ch chan Channel) {
	// TODO: do this calculation by calling the function above
	var inside int = 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < samples; i++ {
		x, y := r.Float64(), r.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}
	ratio := float64(inside) / float64(samples)
	ch <- Channel{samples, ratio * 4} // Send result to channel
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	threads := 8
	var samples float64
	ch := make(chan Channel, threads)

	for i := 1; i <= threads; i++ {
		// Start with 10 samples and go up to 100,000,000
		samples = math.Pow(10, float64(i))
		go calcPiConcurrent(int(samples), ch)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("----------- Running concurrently ----------- ")
	fmt.Println("---------------------------------------------")

	for i := 0; i < threads; i++ {
		output := <-ch // Receive result from channel
		// TODO: render output as a function

		fmt.Print("Samples: ", output.samples)
		// Adjust spaces for formatting
		fmt.Printf("%s", strings.Repeat(" ", 12-int(math.Log10(float64(output.samples)))))
		fmt.Println("Pi:", output.result)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("----------- Running sequentially ----------- ")
	fmt.Println("---------------------------------------------")

	for i := 1; i <= threads; i++ {
		// Start with 10 samples and go up to 100,000,000
		samples = math.Pow(10, float64(i))
		// TODO: render output as a function

		fmt.Print("Samples: ", int(samples))
		// Adjust spaces for formatting
		fmt.Printf("%s", strings.Repeat(" ", 12-i))
		fmt.Println("Pi:", calcPi(int(samples)))
	}
}
