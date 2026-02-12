// flow-control/loop-range/begin/main.go
package main

import (
"fmt"

)

func main() {
	// initialize a slice of ints
	nums := []int{1, 2, 3, 4, 5}

	// use for-range to iterate over the slice and print each value
	for i, num := range nums {
		fmt.Println(i, num) 
	}

	// declare a map of strings to ints
	m:= map[string]int{"one": 1, "two": 2, "three": 3}

	// use for-range to iterate over the map and print each key/value pair
	for key,value := range m{
		fmt.Println(key, value)
	}
}
