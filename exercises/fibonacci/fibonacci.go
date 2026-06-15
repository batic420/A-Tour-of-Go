package main

import (
	"fmt"
	"math"
)

func buildFibonacciChecker() func(x int) []int {
	// Create an empty slice for the future fibonacci numbers
	var numbers []int

	// Create the closure function which returns the `numbers[]` slice
	return func(x int) []int {
		// Declare a variable `n` of type `float64` because `math.Pow()` only accepts floating point numbers
		n := float64(x)

		// Create two calculation for fibonacci checking (at least one of these later needs to create a perfect square).
		// If so, then the number `n` is a fibonacci number.
		calc1 := 5*math.Pow(n, 2.0) + 4
		calc2 := 5*math.Pow(n, 2.0) - 4

		// Use the `isPerfectSquare()` function to check if at least one of calculations created a perfect square
		if isPerfectSquare(calc1) || isPerfectSquare(calc2) {
			// If yes, add `x` to the `numbers[]` slice
			numbers = append(numbers, x)
		}
		return numbers
	}
}

// Create a helper function which calculates the square root and then checks if its a whole number (float64 type with `.0` ending)
func isPerfectSquare(x float64) bool {
	r := math.Sqrt(x)
	return float64(int(r)) == r
}

func main() {
	// Assign the function to a variable
	f := buildFibonacciChecker()
	// Iterate over a specific range of numbers and print the `numbers[]` slice containing all the fibonacci numbers inside the range
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
