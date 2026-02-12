// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("defer cleanup")
	defer cleanup("defer cleanup 2")
	fmt.Printf("Working in main....")
	// defer recovery
	defer func() {
		if r := recover(); r != nil {
			 fmt.Printf("Recovered from panic: %v\n", r) } 
			}()
	// panic
	panic("Something went wrong!")
}
