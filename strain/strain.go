package strain

// Ints is an array of int
type Ints []int

// Lists is a 2D array of int
type Lists []Ints

// Strings is an array of string
type Strings []string

// Keep keeps the list items that meet the given predicate
func (list Ints) Keep(pred func(int) bool) (out Ints) {
	for _, v := range list {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

// Keep keeps the list items that meet the given predicate
func (list Lists) Keep(pred func([]int) bool) (out Lists) {
	for _, v := range list {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

// Keep keeps the list items that meet the given predicate
func (list Strings) Keep(pred func(s string) bool) (out Strings) {
	for _, v := range list {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

// Discard discards the list items that meet the given predicate
func (list Ints) Discard(pred func(int) bool) (out Ints) {
	for _, v := range list {
		if !pred(v) {
			out = append(out, v)
		}
	}
	return
}
