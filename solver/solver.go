// Package solver calcule les instructions qui trient la pile a.
package solver

import "push-swap/stack"

// Stacks regroupe les deux piles et la liste des instructions déjà faites.
type Stacks struct {
	A   []int
	B   []int
	Ops []string
}

// Do applique l'instruction op sur les piles et l'ajoute à la liste Ops.
func Do(p Stacks, op string) Stacks {
	p.A, p.B, _ = stack.Apply(op, p.A, p.B)
	p.Ops = append(p.Ops, op)
	return p
}

// Solve renvoie les instructions qui trient nums.
// Choisit la méthode selon le nombre d'éléments.
func Solve(nums []int) []string {
	p := Stacks{A: nums, B: []int{}}
	switch len(p.A) {
	case 2:
		p = sortTwo(p)
	case 3:
		p = sortThree(p)
	case 4, 5:
		p = sortSmall(p)
	default:
		p = radixSort(p)
	}
	return p.Ops
}

// sortTwo trie une pile a d'exactement 2 nombres.
func sortTwo(p Stacks) Stacks {
	if p.A[0] > p.A[1] {
		p = Do(p, "sa")
	}
	return p
}

// sortThree trie une pile a d'exactement 3 nombres, en 2 instructions maximum.
// Elle envoie d'abord le plus grand en bas, puis remet les deux du haut dans l'ordre.
func sortThree(p Stacks) Stacks {
	if p.A[0] > p.A[1] && p.A[0] > p.A[2] {
		p = Do(p, "ra")
	} else if p.A[1] > p.A[0] && p.A[1] > p.A[2] {
		p = Do(p, "rra")
	}
	if p.A[0] > p.A[1] {
		p = Do(p, "sa")
	}
	return p
}

// sortSmall trie une pile a de 4 ou 5 nombres.
// Elle met de côté les plus petits dans b, trie les 3 restants, puis les ramène.
func sortSmall(p Stacks) Stacks {
	for len(p.A) > 3 {
		p = pushMin(p)
	}
	p = sortThree(p)
	for len(p.B) > 0 {
		p = Do(p, "pa")
	}
	return p
}

// minIndex renvoie la position du plus petit nombre de s.
func minIndex(s []int) int {
	pos := 0
	for i := 1; i < len(s); i++ {
		if s[i] < s[pos] {
			pos = i
		}
	}
	return pos
}

// pushMin amène le plus petit nombre de a au sommet par le chemin le plus court,
// puis l'envoie dans b.
func pushMin(p Stacks) Stacks {
	length := len(p.A)
	pos := minIndex(p.A)
	if pos <= length/2 {
		for k := 0; k < pos; k++ {
			p = Do(p, "ra")
		}
	} else {
		for k := 0; k < length-pos; k++ {
			p = Do(p, "rra")
		}
	}
	p = Do(p, "pb")
	return p
}

// radixSort trie la pile a bit par bit (tri par base), en passant par b.
// Utilisé pour plus de 5 nombres.
func radixSort(p Stacks) Stacks {
	// On travaille sur les rangs : même ordre, mais valeurs de 0 à n-1.
	p.A = ranks(p.A)
	n := len(p.A)
	passes := bitCount(n - 1)
	for bit := 0; bit < passes; bit++ {
		// Bit à 0 : le nombre part dans b. Bit à 1 : il reste dans a.
		for k := 0; k < n; k++ {
			if bitAt(p.A[0], bit) == 0 {
				p = Do(p, "pb")
			} else {
				p = Do(p, "ra")
			}
		}
		for len(p.B) > 0 {
			p = Do(p, "pa")
		}
	}
	return p
}

// ranks remplace chaque nombre par sa position dans l'ordre trié
// (0 pour le plus petit). Le rang d'un nombre = combien de nombres sont plus petits.
func ranks(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}

// bitAt renvoie le bit numéro k de n (0 ou 1).
func bitAt(n, k int) int {
	return (n >> k) & 1
}

// bitCount renvoie le nombre de bits nécessaires pour écrire n.
func bitCount(n int) int {
	count := 0
	for n>>count > 0 {
		count++
	}
	return count
}
