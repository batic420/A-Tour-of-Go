package main

import (
	"fmt"
	"math"
)

func buildFibonacciChecker() func(x int) []int {
	var numbers []int
	return func(x int) []int {
		n := float64(x)
		calc1 := 5*math.Pow(n, 2.0) + 4
		calc2 := 5*math.Pow(n, 2.0) - 4

		if isPerfectSquare(calc1) || isPerfectSquare(calc2) {
			numbers = append(numbers, x)
		}
		return numbers
	}
}

func isPerfectSquare(x float64) bool {
	r := math.Sqrt(x)
	return float64(int(r)) == r
}

func main() {
	f := buildFibonacciChecker()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
