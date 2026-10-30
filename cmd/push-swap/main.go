// Push-swap lit des entiers en arguments et affiche les instructions qui les trient.
package main

import (
	"fmt"
	"os"
	"push-swap/parser"
	"push-swap/solver"
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
	if stack.IsSorted(a) {
		return
	}
	ops := solver.Solve(a)
	for _, op := range ops {
		fmt.Println(op)
	}
}
