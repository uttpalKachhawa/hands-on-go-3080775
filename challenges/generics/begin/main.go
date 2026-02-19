// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring
func print[T any](v T) { fmt.Println(v) }

// non-generic print functions

type numeric interface {
	constraints.Integer | constraints.Float
}

// Part 2 sum function refactoring

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	// call non-generic print functions
		print("Hello")
		print(42)
		print(true)

	// call generic printAny function for each value above

	// call sum function
	fmt.Println(sum(1, 2, 3))
}
