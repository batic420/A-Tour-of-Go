package main

import "golang.org/x/tour/reader"

// Create empty struct (doesn't need to hold any state)
type MyReader struct{}

// Create `Read()` method for the struct which replaces every byte
// inside the input value with `65` (which is the ASCII value for `A`)
func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	// Return the length of `bytes` and `nil` as error
	return len(bytes), nil
}

// Validte the custom struct and the method
// using a `Validate()` method by go.dev
func main() {
	reader.Validate(MyReader{})
}
