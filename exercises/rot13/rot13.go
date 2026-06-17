package main

import (
	"io"
	"os"
	"strings"
)

// Implement a custom `rot13Reader` struct that has a single attribute of `r`
type rot13Reader struct {
	r io.Reader
}

// Create a pointer receiver for the `rot13Reader` struct
func (r *rot13Reader) Read(bytes []byte) (int, error) {
	// Use `Read()` to read all the bytes from the input variable
	// and assign `n` and `e` as variables.
	//
	// This works because `r` of type `rot13Reader` as an `r` attribute which is of type `io.Reader`,
	// it supports the `Read()` method to read a slice of bytes and return an integer
	// (returning the number of bytes that have been read) as well as an error to get details about
	// problems that occurred during the execution.
	n, e := r.r.Read(bytes)

	// Iterate over `n` (number of bytes) and use the `rot13()` method on each
	// iteration to change the value of each byte inside the slice.
	for i := range n {
		bytes[i] = rot13(bytes[i])
	}

	return n, e
}

// Checks whether `b` is in the first or second half of the alphabet
// and adds 13 or subtracts 13 from it.
func rot13(b byte) byte {
	if (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M') {
		return b + 13
	}
	if (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z') {
		return b - 13
	}
	return b
}

// Create `s` of type `strings.Reader` which can be used in the `rot13Reader` struct as attribute which
// then gets assigned to `r`, `io.Copy()` uses `os.Stdout` as writer and a pointer to `r`
// to read from it (and also directly modify it).
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
