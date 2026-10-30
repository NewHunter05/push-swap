package stack

// Apply applique l'instruction op (sa, pb, rrr...) sur les piles a et b.
// Renvoie les deux piles modifiées, et false si l'instruction est inconnue.
func Apply(op string, a, b []int) ([]int, []int, bool) {
	switch op {
	case "sa":
		a = Swap(a)
	case "sb":
		b = Swap(b)
	case "ss":
		a = Swap(a)
		b = Swap(b)
	case "ra":
		a = Rotate(a)
	case "rb":
		b = Rotate(b)
	case "rra":
		a = ReverseRotate(a)
	case "rrb":
		b = ReverseRotate(b)
	case "rr":
		a = Rotate(a)
		b = Rotate(b)
	case "rrr":
		a = ReverseRotate(a)
		b = ReverseRotate(b)
	case "pa":
		b, a = Push(b, a)
	case "pb":
		a, b = Push(a, b)

	default:
		return a, b, false
	}
	return a, b, true
}
