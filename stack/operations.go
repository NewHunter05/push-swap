// Package stack contient les piles et les mouvements possibles.
// Une pile est une slice d'entiers : la case 0 est le sommet.
package stack

// Swap échange les deux éléments du sommet.
// Ne fait rien s'il y a moins de 2 éléments.
func Swap(s []int) []int {
	if len(s) < 2 {
		return s
	}
	s[0], s[1] = s[1], s[0]
	return s
}

// Rotate fait passer le premier élément en dernier.
// Ne fait rien s'il y a moins de 2 éléments.
func Rotate(s []int) []int {
	if len(s) < 2 {
		return s
	}
	return append(s[1:], s[0])
}

// ReverseRotate fait passer le dernier élément en premier.
// Ne fait rien s'il y a moins de 2 éléments.
func ReverseRotate(s []int) []int {
	if len(s) < 2 {
		return s
	}
	bottom := s[len(s)-1]
	return append([]int{bottom}, s[:len(s)-1]...)
}

// Push retire le sommet de from et le pose au sommet de to.
// Renvoie les deux piles modifiées. Ne fait rien si from est vide.
func Push(from, to []int) ([]int, []int) {
	if len(from) == 0 {
		return from, to
	}
	top := from[0]
	from = from[1:]
	to = append([]int{top}, to...)
	return from, to
}
