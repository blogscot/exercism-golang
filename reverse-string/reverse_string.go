package reverse

// String reverses a string
func String(text string) (out string) {
	for _, letter := range text {
		out = string(letter) + out
	}
	return
}
