package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (out []string) {
	if len(s) < n {
		return nil
	}
	out = []string{}

	for index := 0; index <= len(s)-n; index++ {
		out = append(out, s[index:index+n])
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	if substrings := All(n, s); substrings != nil {
		return substrings[0]
	}
	return ""
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	if substrings := All(n, s); substrings != nil {
		return substrings[0], true
	}
	return "", false
}
