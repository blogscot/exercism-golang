package listops

// IntList is an array of int
type IntList []int

type binFunc func(x, y int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldl iterates the given function over a list starting from the left
func (list IntList) Foldl(f binFunc, init int) (acc int) {
	acc = init
	for _, v := range list {
		acc = f(acc, v)
	}
	return
}

// Foldr iterates the given function over a list starting from the right
func (list IntList) Foldr(f binFunc, init int) (acc int) {
	acc = init
	for _, v := range list.Reverse() {
		acc = f(v, acc)
	}
	return
}

// Filter retains list elements that meet the given predicate
func (list IntList) Filter(pred predFunc) (out IntList) {
	if len(list) == 0 {
		return list
	}
	for _, v := range list {
		if pred(v) {
			out = append(out, v)
		}
	}
	return
}

// Length calculates the length of the given list
func (list IntList) Length() int {
	return list.Foldl(func(acc, _ int) int {
		return acc + 1
	}, 0)
}

// Map iterates the given function over a list
func (list IntList) Map(f unaryFunc) (out IntList) {
	if len(list) == 0 {
		return list
	}
	for _, v := range list {
		out = append(out, f(v))
	}
	return

}

// Reverse reverses a list
func (list IntList) Reverse() IntList {
	length := len(list)
	if length <= 1 {
		return list
	}
	return append(list[1:].Reverse(), list[0])
}

// Append appends one list to another
func (list IntList) Append(other IntList) IntList {
	return append(list, other...)
}

// Concat concats the given arguments to a list
func (list IntList) Concat(args []IntList) (out IntList) {
	out = list
	for _, arg := range args {
		out = out.Append(arg)
	}
	return
}
