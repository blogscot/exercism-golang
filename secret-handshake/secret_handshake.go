package secret

var secretNumbers = []int{1, 2, 4, 8}
var secretGestures = []string{"wink", "double blink", "close your eyes", "jump"}

const reverse = 16

// Handshake converts a number into a list of secrets.
func Handshake(number uint) (out []string) {
	num := int(number)
	out = []string{}

	var test = func(key int) {
		if num&secretNumbers[key] > 0 {
			out = append(out, secretGestures[key])
		}
	}

	if num&reverse > 0 {
		for key := len(secretGestures) - 1; key >= 0; key-- {
			test(key)
		}
	} else {
		for key := 0; key < len(secretGestures); key++ {
			test(key)
		}
	}
	return
}
