package main

// Import the `pic` package from golang.org
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create an array using `dy` as length
	a := make([][]uint8, dy)

	// Iterate over the array and create second array for each iteration creating a two-dimensional array
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Iterate over the two-dimensional array using a nested for loop and create a calculation for each entry
	// creating a unique blue-scaled picture using the values calculated on each iteration
	for y := range a {
		for x := range a[y] {
			// Use this calculation or try something else to generate other images, for example:
			// - (x ^ y) * (x ^ y)
			// - x*x + y*y
			// - y * 10000 / (x + 1)
			// - (x+y)/2 * (x^y)
			// - ((x - y) * -1) * (x + y) / 2
			// - 2*y ^ 2*3*x ^ 3
			// ...
			a[y][x] = uint8((x + y) / 2 * (x ^ y) << 2)
		}
	}

	// Return the array to display the image
	return a
}

func main() {
	// Use the imported package to display the picture
	pic.Show(Pic)
}
