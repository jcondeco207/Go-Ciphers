package xor

func XorString(text string, key byte) string {

	encrypted := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		encrypted[i] = text[i] ^ key
	}

	return string(encrypted)
}
