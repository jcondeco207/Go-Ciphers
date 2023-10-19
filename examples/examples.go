package main

import (
	"fmt"

	"jcondeco-ciphers.com/affine"
	"jcondeco-ciphers.com/xor"
)

func main() {

	original := "secret sentence"

	// Affine example
	fmt.Println("Before encryption: " + original)
	fmt.Println("After encryption: " + affine.Encrypt(original, 11, 17, false)) // Set true if you want spaces as well

	// Xor example
	original = "Hello, XOR Cipher!"

	fmt.Println("Before encryption: " + original)
	xorCipher := xor.XorString(original, byte(0x7F))
	fmt.Println("After encryption: " + xorCipher)
	fmt.Println("Back to normal: " + xor.XorString(xorCipher, byte(0x7F)))
}
