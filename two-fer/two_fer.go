package twofer

import "fmt"

// ShareWith shares one for `name` and me
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
