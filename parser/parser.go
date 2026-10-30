// Package parser lit et vérifie les arguments du programme.
package parser

import (
	"errors"
	"strconv"
	"strings"
)

// Parse transforme les arguments en liste d'entiers.
// Accepte "3 2 1" (un seul argument) comme 3 2 1 (plusieurs arguments).
// Renvoie une erreur si un mot n'est pas un entier ou s'il y a un doublon.
func Parse(args []string) ([]int, error) {
	words := []string{}
	for _, arg := range args {
		words = append(words, strings.Fields(arg)...)
	}

	numbers := []int{}
	for _, word := range words {
		n, err := strconv.Atoi(word)
		if err != nil {
			return nil, errors.New("argument non entier")
		}
		numbers = append(numbers, n)
	}
	if HasDuplicates(numbers) {
		return nil, errors.New("doublon")
	}

	return numbers, nil
}

// HasDuplicates renvoie true si un même nombre apparaît deux fois.
func HasDuplicates(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
