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
	nombres, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	a := nombres
	b := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	var ok bool
	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne == "" {
			continue
		}
		a, b, ok = stack.Apply(ligne, a, b)
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
