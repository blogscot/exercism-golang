package secret

import "sort"

var secrets = map[int]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake converts a number into a list of secrets.
func Handshake(number uint) (out []string) {
	num := int(number)
	out = []string{}

	keys := []int{}
	for k := range secrets {
		keys = append(keys, int(k))
	}

	if num&16 > 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Ints(keys)
	}

	for _, key := range keys {
		if num&key > 0 {
			out = append(out, secrets[key])
		}
	}
	return
}
