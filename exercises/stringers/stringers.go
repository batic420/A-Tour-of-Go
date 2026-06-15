package main

import (
	"fmt"
	"strings"
)

// Create the custom type for the IP address
type IPAddr [4]byte

// Implement a Stringer that will be used by the `fmt` package
// to print the output in `main()` in the desired format
func (ip IPAddr) String() string {
	// Use the `strings` package's `Builder` type to build the custom string
	var b strings.Builder

	// Iterate over the IP (the 4 bytes)
	for _, v := range ip {
		// Add each value to `b` using a pointer (so the result gets saved outside the for loop)
		// in a specific format ending each value insert with a `.`
		fmt.Fprintf(&b, "%d.", v)
	}

	// Remove the trailing dot as it doesn't belong to an IP
	return strings.Trim(b.String(), ".")
}

func main() {
	// Create a host map with values of type `IPAddr` (which can use the Stringer)
	hosts := map[string]IPAddr{
		"loopback":   {127, 0, 0, 1},
		"googleDNS":  {8, 8, 8, 8},
		"googleDNS2": {8, 8, 4, 4},
	}

	// Print each value to the console (the `fmt` package will look for the Stringer,
	// check if the type matches and then uses it to print out the result in the desired format)
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
