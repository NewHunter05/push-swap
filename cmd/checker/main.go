// Checker lit des entiers en arguments et des instructions sur l'entrée standard.
// Il applique les instructions puis affiche OK si a est triée et b vide, sinon KO.
package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/parser"
	"push-swap/stack"
)

func main() {
	if len(os.Args[1:]) == 0 {
		return
	}
	numbers, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	a := numbers
	b := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	var ok bool
	for scanner.Scan() {
		lign := scanner.Text()
		if lign == "" {
			continue
		}
		a, b, ok = stack.Apply(lign, a, b)
		if !ok {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	if stack.IsSorted(a) && len(b) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
