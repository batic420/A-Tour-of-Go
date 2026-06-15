package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Create an empty map to use in the for loop
	count := make(map[string]int)

	// Iterate over the slice of separate string that are the response of the `strings.Fields()` function
	for _, v := range strings.Fields(s) {
		// Because of Go's "Zero Value Rule" if an object in the map doesn't exist, `0` will be returned,
		// you don't need to check if that value already exists.
		//
		// Example:
		// If a word appears for the first time in the string, `count[v]` will be `0` (and not `nil` or something),
		// using `count[v] + 1` as assignment to that, the value will then be `1`.
		// If the same phrase appears another time, it now already exists in the map with a certain value, for example,
		// `1` - the assignment will now just add `1` to the existing positive value.
		//
		// Takeaway:
		// You don't need to check if a value exists, by default its just `0` - Go won't throw any error. Just assume
		// the value is there but its value is `0`, so you can directly start using it and editing it.
		count[v] = (count[v] + 1)
	}
	return count
}

func main() {
	// Use the `wc.Test()` function from go.dev to test your `WordCount()` function
	wc.Test(WordCount)
}
