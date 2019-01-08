package listops

// IntList is an array of int
type IntList []int

type binFunc func(x, y int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldl iterates the given function over a list starting from the left
func (list IntList) Foldl(f binFunc, acc int) int {
	for _, v := range list {
		acc = f(acc, v)
	}
	return acc
}

// Foldr iterates the given function over a list starting from the right
func (list IntList) Foldr(f binFunc, acc int) int {
	for _, v := range list.Reverse() {
		acc = f(v, acc)
	}
	return acc
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
	return len(list)
}

// Map iterates the given function over a list
func (list IntList) Map(f unaryFunc) IntList {
	for i := range list {
		list[i] = f(list[i])
	}
	return list

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
func (list IntList) Concat(args []IntList) IntList {
	for _, arg := range args {
		list = list.Append(arg)
	}
	return list
}
