package main

import (
	"fmt"

	"jcondeco-ciphers.com/affine"
)

func main() {

	original := "secret sentence"

	// Affine example
	fmt.Println("Before encryption: " + original)
	fmt.Println("After encryption: " + affine.Encrypt(original, 11, 17, false)) // Set true if you want spaces as well
}
