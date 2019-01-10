package reverse

// String reverses a string
func String(text string) string {
	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[length-1-i], runes[i] = runes[i], runes[length-1-i]
	}
	return string(runes)
}
