// Package parser lit et vérifie les arguments du programme.
package parser

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(args []string) ([]int, error) {
	mots := []string{}
	for _, arg := range args {
		mots = append(mots, strings.Fields(arg)...)
	}

	nombres := []int{}
	for _, mot := range mots {
		n, err := strconv.Atoi(mot)
		if err != nil {
			return nil, errors.New("argument non entier")
		}
		nombres = append(nombres, n)
	}
	if HasDuplicates(nombres) {
		return nil, errors.New("doublon")
	}

	return nombres, nil
}
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
