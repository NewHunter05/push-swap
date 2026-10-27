package stack

func Swap(s []int) []int {
	if len(s) < 2 {
		return s
	}
	s[0], s[1] = s[1], s[0]
	return s
}

func Rotate(s []int) []int {
	if len(s) < 2 {
		return s
	}
	return append(s[1:], s[0])
}

func ReverseRotate(s []int) []int {
	if len(s) < 2 {
		return s
	}
	fond := s[len(s)-1]
	return append([]int{fond}, s[:len(s)-1]...)
}

func Push(from, to []int) ([]int, []int) {
	if len(from) == 0 {
		return from, to
	}
	top := from[0]
	from = from[1:]
	to = append([]int{top}, to...)
	return from, to
}
