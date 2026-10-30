package stack

// IsSorted renvoie true si s est triée par ordre croissant.
func IsSorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
