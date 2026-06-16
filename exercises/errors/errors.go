package main

import (
	"fmt"
	"math"
)

// Create a new custom error type
type ErrNegativeSqrt float64

// Create a value receiver for the custom type which will be called
// by the `fmt` package later on
func (e ErrNegativeSqrt) Error() string {
	// Convert `e` into a supported `float64`,
	// otherwise the program would be sent in an infinite loop
	e_fl64 := fmt.Sprint(float64(e))

	// Return the custom error message
	return fmt.Sprint("cannot Sqrt negative number: ", e_fl64, "\n")
}

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	var t float64

	// Update the existing function to return an error if `x` is negative
	// with `0` as value for the `float64` value
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}

	for {
		z, t = z-(z*z-x)/(2*z), z
		fmt.Println(z, t)
		if math.Abs(t-z) < 1e-8 {
			break
		}
	}

	// Return `z` and `nil` for no error
	return z, nil
}

func main() {
	// Test the function twice with a supported parameter versus unsupported
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
